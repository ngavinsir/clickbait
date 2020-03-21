package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/ngavinsir/clickbait/handlers"
	"github.com/ngavinsir/clickbait/model"
)

//go:generate sqlboiler --wipe psql

func main() {
	router := chi.NewRouter()

	db, err := setupDB()
	handleErr(err)
	log.Println("connected to db")

	env := handlers.CreateEnv(db)

	//inputDataset("./dataset/cnn.csv", db)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})
	router.Use(c.Handler)

	router.Post("/register", handlers.Register(db.DB))
	router.Post("/login", handlers.Login(db.DB))

	router.Group(func(router chi.Router) {
		router.Use(handlers.AuthMiddleware)
		
		router.Route("/{labelType}", func(router chi.Router) {
			router.Route("/article", func(router chi.Router) {
				router.Get("/random", env.RandomArticle)
			})
	
			router.Route("/label", func(router chi.Router) {
				router.Get("/", env.GetAllLabel)
				router.Post("/", env.AddLabel)
				router.Route("/{labelID}", func(router chi.Router) {
					router.Delete("/", env.DeleteLabel)
				})
			})
	
			router.Post("/labeling", env.Labeling)
		})
		
		router.Post("/article", env.AddArticle)
	})

	name, _ := os.Executable()
	port := ":4040"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = fmt.Sprintf(":%s", envPort)
	}

	log.Printf("Server started on %s, pid: %s", port, name)
	log.Fatal(http.ListenAndServe(port, router))
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func setupDB() (*model.DB, error) {
	conn := "dbname=clickbait host=localhost user=postgres password=postgres"
	if envConn := os.Getenv("DATABASE_URL"); envConn != "" {
		conn = envConn
	}

	return model.InitDB(conn)
}

func inputDataset(datasetPath string, db *model.DB) {
	csvfile, err := os.Open(datasetPath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	articleRepository := &model.ArticleDatastore{db}

	i := 0
	for row := range processCSV(csvfile) {
		i++
		log.Printf("%s\n", row[2])
		articleRepository.InsertArticle(context.Background(), row[2], row[0])
	}
}

func processCSV(rc io.Reader) (ch chan []string) {
	ch = make(chan []string, 10)
	go func() {
		r := csv.NewReader(rc)
		if _, err := r.Read(); err != nil { //read header
			log.Fatal(err)
		}
		defer close(ch)
		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)

			}
			ch <- rec
		}
	}()
	return
}

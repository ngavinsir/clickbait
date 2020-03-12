package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"github.com/ngavinsir/clickbait/handlers"
	"github.com/ngavinsir/clickbait/model"
	"github.com/volatiletech/sqlboiler/boil"
)

//go:generate sqlboiler --wipe psql

func main() {
	router := chi.NewRouter()

	db, err := setupDB()
	handleErr(err)

	//inputDataset("./dataset/cnn.csv", db)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})
	router.Use(c.Handler)

	router.Post("/register", handlers.Register(db))
	router.Post("/login", handlers.Login(db))

	router.Group(func(router chi.Router) {
		router.Use(handlers.AuthMiddleware)

		router.Route("/article", func(router chi.Router) {
			router.Post("/", handlers.AddArticle(db))
			router.Get("/random/{labelType}", handlers.RandomArticle(db))
		})

		router.Route("/label/{labelType}", func(router chi.Router) {
			router.Get("/", handlers.GetAllLabel(db))
			router.Post("/", handlers.AddLabel(db))
			router.Route("/{labelID}", func(router chi.Router) {
				router.Delete("/", handlers.DeleteLabel(db))
			})
		})

		router.Post("/labeling/{labelType}", handlers.Labeling(db))
	})

	name, _ := os.Executable()
	log.Printf("Server started on :4040, pid: %s", name)
	log.Fatal(http.ListenAndServe(":4040", router))
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func setupDB() (*sql.DB, error) {
	return sql.Open("postgres", `dbname=clickbait host=localhost user=postgres password=postgres`)
}

func inputDataset(datasetPath string, exec boil.ContextExecutor) {
	csvfile, err := os.Open(datasetPath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	i := 0
	for row := range processCSV(csvfile) {
		i++
		log.Printf("%s\n", row[2])
		model.InsertArticle(context.Background(), exec, row[2], row[0])
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

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/ngavinsir/clickbait/handlers"
)

//go:generate sqlboiler --wipe psql

func main() {
	router := chi.NewRouter()

	db, err := setupDB()
	handleErr(err)
	
	router.Post("/register", handlers.Register(db))
	router.Post("/login", handlers.Login(db))

	router.Group(func(router chi.Router) {
		router.Use(handlers.AuthMiddleware)

		router.Route("/headline", func(router chi.Router) {
			router.Post("/", handlers.AddHeadline(db))
			router.Get("/random", handlers.RandomHeadline(db))
		})

		router.Post("/label", handlers.AddLabel(db))
	})

	log.Println("Server started on :4040")
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
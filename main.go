package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	_ "github.com/lib/pq"
	"github.com/ngavinsir/clickbait/handlers"
)

//go:generate sqlboiler --wipe psql

func main() {
	router := chi.NewRouter()

	db, err := setupDB()
	handleErr(err)
	
	tokenAuth := jwtauth.New("HS256", []byte("clickbait^secret"), nil)

	router.Post("/register", handlers.Register(db))
	router.Post("/login", handlers.Login(db, tokenAuth))

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
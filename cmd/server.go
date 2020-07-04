package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/ngavinsir/clickbait/handlers"
	"github.com/ngavinsir/clickbait/model"
	"github.com/spf13/cobra"
)

var cmdServer = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		router := chi.NewRouter()

		db, err := model.InitDB()
		if err != nil {
			panic(err)
		}
		log.Println("connected to db")
		defer db.Close()

		env := handlers.CreateEnv(db)

		c := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		})
		router.Use(c.Handler)
		router.Use(middleware.RequestID)
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)

		router.Post("/register", env.Register)
		router.Post("/login", env.Login)

		router.Group(func(router chi.Router) {
			router.Use(env.AuthMiddleware)

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
				router.Get("/leaderboard/{limit}", env.GetLabelLeaderboard)
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
	},
}

package main

import (
	"net/http"

	"example.com/serve_html/pkg/config"
	"example.com/serve_html/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// Can also use pat to implement this.
	// Mux := pat.New()
	// Mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// Mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()

	// middleware allows us to process request before handle it.
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

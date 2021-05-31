package main

import (
	"log"
	"net/http"
	"time"

	"example.com/serve_html/pkg/config"
	"example.com/serve_html/pkg/handlers"
	"example.com/serve_html/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.AppConfig
var sessionManager *scs.SessionManager

func main() {
	// Changethis to true in production
	app.InProduction = false

	// Create a new session manager.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	// In production you want this to be true.
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplaceCache = tc
	app.UseCache = false

	render.NewTemplates(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	log.Println("Starting server on port:", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

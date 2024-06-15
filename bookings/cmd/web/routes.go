package main

import (
	"net/http"

	"github.com/devj1003/go-course/pkg/config"
	"github.com/devj1003/go-course/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(App *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.About))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Home)
	mux.Get("/about", handlers.About)

	return mux
}

package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/devj1003/go-course/pkg/config"
	"github.com/devj1003/go-course/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/about", http.HandlerFunc(handlers.About))

	return mux
}

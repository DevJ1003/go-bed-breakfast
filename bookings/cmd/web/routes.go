package main

import (
	"net/http"

	"github.com/devj1003/bookings/internal/config"
	"github.com/devj1003/bookings/internal/handlers"
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
	mux.Get("/generals-quarter", handlers.Generals)
	mux.Get("/majors-suite", handlers.Majors)

	mux.Get("/search-availability", handlers.Availability)
	mux.Post("/search-availability", handlers.PostAvailability)
	mux.Post("/search-availability-json", handlers.AvailabilityJSON)

	mux.Get("/reservation", handlers.Reservation)
	mux.Post("/reservation", handlers.PostReservation)
	// mux.Get("/reservation-summary", handlers.ReservationSummary)

	mux.Get("/contact", handlers.Contact)

	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}

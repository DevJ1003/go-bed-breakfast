package main

import (
	"net/http"

	"github.com/devj1003/bookings/internal/config"
	"github.com/devj1003/bookings/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(App *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarter", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)

	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Post("/reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/user/login", handlers.Repo.ShowLogin)
	mux.Post("/user/login", handlers.Repo.PostShowLogin)
	mux.Get("/user/logout", handlers.Repo.Logout)

	mux.Get("/contact", handlers.Repo.Contact)

	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.Route("/admin", func(mux chi.Router) {
		// mux.Use(Auth) is used for login, disabled till the completion of the project
		mux.Use(Auth)

		mux.Get("/dashboard", handlers.Repo.AdminDashboard)

		mux.Get("/reservations-new", handlers.Repo.AdminNewReservations)
		mux.Get("/reservations-all", handlers.Repo.AdminAllReservations)
		mux.Get("/reservations/{src}/{id}/show", handlers.Repo.AdminShowReservations)
		mux.Post("/reservations/{src}/{id}", handlers.Repo.AdminPostShowReservations)
		mux.Get("/process-reservation/{src}/{id}/do", handlers.Repo.AdminProcessReservations)
		mux.Get("/delete-reservation/{src}/{id}/do", handlers.Repo.AdminDeleteReservations)

		mux.Get("/reservations-calendar", handlers.Repo.AdminReservationsCalendar)
		mux.Post("/reservations-calendar", handlers.Repo.AdminPostReservationsCalendar)
	})

	return mux
}

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/devj1003/bookings/internal/config"
	"github.com/devj1003/bookings/internal/drivers"
	"github.com/devj1003/bookings/internal/forms"
	"github.com/devj1003/bookings/internal/helpers"
	"github.com/devj1003/bookings/internal/models"
	"github.com/devj1003/bookings/internal/render"
	"github.com/devj1003/bookings/internal/repository"
	"github.com/devj1003/bookings/internal/repository/dbrepo"
)

var App config.AppConfig
var session *scs.SessionManager
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *drivers.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewMysqlRepo(db.SQL, a),
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// HOME page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})

}

// ABOUT page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {

	// Repo.App.Session.Put(r.Context(), "error", "No Availability")
	fmt.Println("test", Repo.App.Session.Get(r.Context(), "error"))
	var errorStatus bool
	if Repo.App.Session.Get(r.Context(), "error") != nil {
		errorStatus = true
	}

	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{
		ShowError: errorStatus,
	})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, start)
	if err != nil {
		helpers.ServerError(w, err)
	}
	endDate, err := time.Parse(layout, end)
	if err != nil {
		helpers.ServerError(w, err)
	}

	rooms, err := Repo.DB.SearchAvailabilityForAllRooms(startDate, endDate)
	if err != nil {
		helpers.ServerError(w, err)
	}

	if len(rooms) == 0 {
		// no availability
		Repo.App.Session.Put(r.Context(), "error", "No Availability")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		// return
	}

	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.Template(w, r, "reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		// log.Println(err)
		helpers.ServerError(w, err)
		return
	}

	sd := r.Form.Get("start_date")
	ed := r.Form.Get("end_date")

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, sd)
	if err != nil {
		helpers.ServerError(w, err)
	}
	endDate, err := time.Parse(layout, ed)
	if err != nil {
		helpers.ServerError(w, err)
	}

	roomID, err := strconv.Atoi(r.Form.Get("room_id"))

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomID:    roomID,
	}

	form := forms.New(r.PostForm)

	// Rules for form validation
	// form.Has("first_name", r)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.Template(w, r, "reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})

	}

	// fmt.Println("repo:", Repo)
	// fmt.Println("m:", m)

	newReservationID, err := Repo.DB.InsertReservation(reservation)
	if err != nil {
		App.Session.Put(r.Context(), "error", "can't insert room restriction!")
		helpers.ServerError(w, err)
	}

	restriction := models.RoomRestriction{
		StartDate:     startDate,
		EndDate:       endDate,
		RoomID:        roomID,
		ReservationID: newReservationID,
		RestrictionID: 1,
	}

	err = Repo.DB.InsertRoomRestriction(restriction)
	if err != nil {
		App.Session.Put(r.Context(), "error", "can't insert room restriction!")
		helpers.ServerError(w, err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

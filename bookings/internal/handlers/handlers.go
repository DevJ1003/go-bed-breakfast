package handlers

import (
	"encoding/json"
	"errors"
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
	"github.com/go-chi/chi"
)

var App config.AppConfig
var Session *scs.SessionManager
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

	var errorStatus bool
	var msgType string
	var msgText string
	if Repo.App.Session.Get(r.Context(), "error_for_login") != nil {
		errorStatus = true
		msgType = "success"
		msgText = "Logged-in Successfully!"
	}

	render.Template(w, r, "home.page.tmpl", &models.TemplateData{
		ShowError:   errorStatus,
		MessageType: msgType,
		MessageText: msgText,
	})

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

	var errorStatus bool
	var msgType string
	var msgText string
	if Repo.App.Session.Get(r.Context(), "error") != nil {
		errorStatus = true
		msgType = "warning"
		msgText = "No Availability!"
	}

	render.Template(w, r, "search-availability.page.tmpl", &models.TemplateData{
		ShowError:   errorStatus,
		MessageType: msgType,
		MessageText: msgText,
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
		return
	}

	if len(rooms) == 0 {
		// no availability
		Repo.App.Session.Put(r.Context(), "error", "No Availability")
		http.Redirect(w, r, "/search-availability", http.StatusSeeOther)
		// return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	res := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}

	Repo.App.Session.Put(r.Context(), "reservation", res)

	// w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
	render.Template(w, r, "/choose-room.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

func (m *Repository) ChooseRoom(w http.ResponseWriter, r *http.Request) {

	roomID, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res, ok := Repo.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, err)
		return
	}

	res.RoomID = roomID
	Repo.App.Session.Put(r.Context(), "reservation", res)
	http.Redirect(w, r, "/reservation", http.StatusSeeOther)

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

	res, ok := Repo.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		helpers.ServerError(w, errors.New("cannot get reservation from session"))
		return
	}

	roomname, err := Repo.DB.GetRoomByID(res.RoomID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	res.Room.RoomName = roomname

	sd := res.StartDate.Format("2006-01-02")
	ed := res.EndDate.Format("2006-01-02")
	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed

	data := make(map[string]interface{})
	data["reservation"] = res

	render.Template(w, r, "reservation.page.tmpl", &models.TemplateData{
		Form:      forms.New(nil),
		Data:      data,
		StringMap: stringMap,
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
	if err != nil {
		helpers.ServerError(w, err)
	}

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

	// ===================Email Code===============================================================================
	// sends email to the customer
	htmlMessage := fmt.Sprintf(`
		<strong>Reservation Confirmation!</strong><br>
		Dear %s,<br>
		This is to confirm your reservation from %s to %s.
	`, reservation.FirstName, reservation.StartDate.Format("2006-01-02"), reservation.EndDate.Format("2006-01-02"))

	msg := models.MailData{
		To:      reservation.Email,
		From:    "me@here.com",
		Subject: "Reservation Confirmation",
		Content: htmlMessage,
	}

	Repo.App.MailChain <- msg
	//
	// sends email to the company
	htmlMessage = fmt.Sprintf(`
		<strong>Reservation Confirmation!</strong><br>
		Dear %s,<br>
		This is to confirm your reservation from %s to %s.
	`, reservation.FirstName, reservation.StartDate.Format("2006-01-02"), reservation.EndDate.Format("2006-01-02"))

	msg = models.MailData{
		To:      "me@here.com",
		From:    "me@here.com",
		Subject: "Reservation Confirmation",
		Content: htmlMessage,
	}

	Repo.App.MailChain <- msg
	// ============================================================================================================

	Repo.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)

}

// ReservationSummary displays the res summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := Repo.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		Repo.App.ErrorLog.Println("Can't get reservation from session")
		Repo.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	roomname, err := Repo.DB.GetRoomByID(reservation.RoomID)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	Repo.App.Session.Remove(r.Context(), "reservation")

	sd := reservation.StartDate.Format("2006-01-02")
	ed := reservation.EndDate.Format("2006-01-02")
	stringMap := make(map[string]string)
	stringMap["start_date"] = sd
	stringMap["end_date"] = ed
	stringMap["roomname"] = roomname

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.Template(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data:      data,
		StringMap: stringMap,
	})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) ShowLogin(w http.ResponseWriter, r *http.Request) {

	var errorStatus bool
	var msgType string
	var msgText string
	if Repo.App.Session.Get(r.Context(), "error_for_login") != nil {
		errorStatus = true
		msgType = "warning"
		msgText = "Invalid Credentials!"
	}

	render.Template(w, r, "login-page.tmpl", &models.TemplateData{
		Form:        forms.New(nil),
		ShowError:   errorStatus,
		MessageType: msgType,
		MessageText: msgText,
	})

}

// PostShowLogin handles logging the user in
func (m *Repository) PostShowLogin(w http.ResponseWriter, r *http.Request) {

	_ = Repo.App.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	form := forms.New(r.PostForm)
	form.Required("email", "password")
	form.IsEmail("email")

	if !form.Valid() {
		render.Template(w, r, "login.page.tmpl", &models.TemplateData{
			Form: form,
		})
		return
	}

	id, _, err := Repo.DB.Authenticate(email, password)
	if err != nil {
		log.Println(err)

		Repo.App.Session.Put(r.Context(), "error_for_login", "Invalid login credentials")
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}

	Repo.App.Session.Put(r.Context(), "user_id", id)
	Repo.App.Session.Put(r.Context(), "flash", "Logged in successfully!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Logout logs user out
func (m *Repository) Logout(w http.ResponseWriter, r *http.Request) {

	_ = Repo.App.Session.Destroy(r.Context())
	_ = Repo.App.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (m *Repository) AdminDashboard(w http.ResponseWriter, r *http.Request) {
	render.AdminTemplate(w, r, "admin-dashboard.page.tmpl", &models.TemplateData{})
}

func (m *Repository) AdminAllReservations(w http.ResponseWriter, r *http.Request) {
	render.AdminTemplate(w, r, "admin-all-reservations.page.tmpl", &models.TemplateData{})
}

func (m *Repository) AdminNewReservations(w http.ResponseWriter, r *http.Request) {
	render.AdminTemplate(w, r, "admin-new-reservations.page.tmpl", &models.TemplateData{})
}

func (m *Repository) AdminReservationsCalendar(w http.ResponseWriter, r *http.Request) {
	render.AdminTemplate(w, r, "admin-reservations-calendar.page.tmpl", &models.TemplateData{})
}

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/devj1003/bookings/internal/config"
	"github.com/devj1003/bookings/internal/forms"
	"github.com/devj1003/bookings/internal/helpers"
	"github.com/devj1003/bookings/internal/models"
	"github.com/devj1003/bookings/internal/render"
)

var App config.AppConfig
var session *scs.SessionManager

// HOME page handler
func Home(w http.ResponseWriter, r *http.Request) {

	// // var App *config.AppConfig

	// var Session *scs.SessionManager

	// // App.Session = session

	// remoteIP := r.RemoteAddr
	// Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})

}

// ABOUT page handler
func About(w http.ResponseWriter, r *http.Request) {

	// remoteIP := session.GetString(r.Context(), "remote_ip")
	// var Tdata models.TemplateData
	// Tdata.StringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{})
}

func Generals(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func Majors(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func Availability(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

func PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
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

func Reservation(w http.ResponseWriter, r *http.Request) {

	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

func PostReservation(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		// log.Println(err)
		helpers.ServerError(w, err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
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

		render.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})

	}

	// App.Session = session
	// session.Put(r.Context(), "reservation", reservation)
	// // http.Redirect wants ResponseWriter, Request
	// http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// NOT WORKING
// func ReservationSummary(w http.ResponseWriter, r *http.Request) {

// 	// reservation, ok := App.Session.Get(r.Context(), "reservation").(models.Reservation)
// 	// if !ok {
// 	// 	log.Println("cannot get item from session")
// 	// 	return
// 	// }

// 	// data := make(map[string]interface{})
// 	// data["reservation"] = reservation

// 	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{})
// 	// 	Data: data,
// 	// })
// }

func Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

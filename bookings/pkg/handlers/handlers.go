package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/devj1003/bookings/pkg/models"
	"github.com/devj1003/bookings/pkg/render"
)

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

	render.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{})
}

func Contact(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}

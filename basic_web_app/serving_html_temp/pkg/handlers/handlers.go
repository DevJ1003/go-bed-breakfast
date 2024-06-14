package handlers

import (
	"net/http"

	"github.com/devj1003/go-course/pkg/render"
)

// HOME page handler
func Home(w http.ResponseWriter, r *http.Request) {

	// // var App *config.AppConfig

	// var Session *scs.SessionManager

	// // App.Session = session

	// remoteIP := r.RemoteAddr
	// Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl")

}

// ABOUT page handler
func About(w http.ResponseWriter, r *http.Request) {

	// remoteIP := session.GetString(r.Context(), "remote_ip")
	// var Tdata models.TemplateData
	// Tdata.StringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl")
}

package handlers

import (
	"net/http"

	"github.com/devj1003/go-course/pkg/render"
)

// HOME page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")

	remoteIP := r.RemoteAddr
	app.Session.Put(r.Context(), "remote_ip", remoteIP)

}

// ABOUT page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")

	remoteIP := app.Session.GetString(r.Context(), "remote_ip")
	StringMap["remote_ip"] = remoteIP
}

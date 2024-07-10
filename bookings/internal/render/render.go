package render

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/devj1003/bookings/internal/config"
	"github.com/devj1003/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{
	"humanDate": HumanDate,
}

var App *config.AppConfig

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	App = a
}

// HumanDate returns the date in YYYY-MM-DD format
func HumanDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = App.Session.PopString(r.Context(), "flash")
	td.Error = App.Session.PopString(r.Context(), "error")
	td.Warning = App.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	if App.Session.Exists(r.Context(), "user_id") {
		td.IsAuthenticated = 1
	}
	return td
}

// RENDERING the html templates using html/template
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("../../templates/"+tmpl, "../../templates/base.layout.tmpl")

	td = AddDefaultData(td, r)

	err := parsedTemplate.Execute(w, td)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}

// RENDERING the html templates using html/template
func AdminTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("../../templates/"+tmpl, "../../templates/admin.layout.tmpl")

	td = AddDefaultData(td, r)

	err := parsedTemplate.Execute(w, td)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}

package render

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/devj1003/bookings/internal/config"
	"github.com/devj1003/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
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

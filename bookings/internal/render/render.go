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

// var functions = template.FuncMap{
// 	"humanDate":  HumanDate,
// 	"formatDate": FormatDate,
// 	"iterate":    Iterate,
// }

var App *config.AppConfig

// // NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	App = a
}

// // HumanDate returns the date in YYYY-MM-DD format
func HumanDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func FormatDate(t time.Time, f string) string {
	return t.Format(f)
}

// Iterate returns a slice of ints, starting at 1, going to count
func Iterate(count int) []int {
	var i int
	var items []int
	for i = 0; i < count; i++ {
		items = append(items, i)
	}

	return items
}

func Add(a, b int) int {
	return a + b
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
	parsedTemplate, _ := template.New(tmpl).Funcs(template.FuncMap{
		"humanDate":  HumanDate,
		"formatDate": FormatDate,
		"iterate":    Iterate,
		"add":        Add,
	}).ParseFiles("../../templates/"+tmpl, "../../templates/base.layout.tmpl")

	// template.ParseFiles("../../templates/"+tmpl, "../../templates/base.layout.tmpl")

	td = AddDefaultData(td, r)

	err := parsedTemplate.Execute(w, td)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}

// RENDERING the html templates using html/template
func AdminTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	parsedTemplate, _ := template.New(tmpl).Funcs(template.FuncMap{
		"humanDate":  HumanDate,
		"formatDate": FormatDate,
		"iterate":    Iterate,
		"add":        Add,
	}).ParseFiles("../../templates/"+tmpl, "../../templates/admin.layout.tmpl")

	// parsedTemplate, _ := template.ParseFiles("../../templates/"+tmpl, "../../templates/admin.layout.tmpl")

	td = AddDefaultData(td, r)

	err := parsedTemplate.Execute(w, td)

	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}

}

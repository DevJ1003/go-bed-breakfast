package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/devj1003/mygoapp/pkg/config"
	"github.com/devj1003/mygoapp/pkg/handlers"
	"github.com/devj1003/mygoapp/pkg/render"
)

// Port number function
const portNumber = ":8081"

var app config.AppConfig
var session *scs.SessionManager

// Main func
func main() {

	// var app config.AppConfig

	// change this to true when in production
	app.InProduction = false

	// Session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		// fmt.Println(err)
		log.Fatal("Cannot create template!")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.HomePage)
	// http.HandleFunc("/about", handlers.Repo.AboutPage)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

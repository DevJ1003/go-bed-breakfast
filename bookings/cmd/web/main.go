package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/devj1003/bookings/internal/config"
	"github.com/devj1003/bookings/internal/handlers"
)

const portNumber = ":8080"

var App config.AppConfig
var session *scs.SessionManager

func main() {

	App.UseCache = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = App.InProduction

	App.Session = session

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&App),
	}

	srv.ListenAndServe()
}

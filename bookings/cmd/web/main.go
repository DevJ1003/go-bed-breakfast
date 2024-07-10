package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/devj1003/bookings/internal/config"
	"github.com/devj1003/bookings/internal/drivers"
	"github.com/devj1003/bookings/internal/handlers"
	"github.com/devj1003/bookings/internal/helpers"
	"github.com/devj1003/bookings/internal/models"
	"github.com/devj1003/bookings/internal/render"
)

const portNumber = ":8080"

var App config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
	defer close(App.MailChain)
	fmt.Println("Starting mail listner...  ")
	listenForMail()

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&App),
	}

	srv.ListenAndServe()
}

func run() (*drivers.DB, error) {

	App.UseCache = false

	// what am i going to put in session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	mailChan := make(chan models.MailData)
	App.MailChain = mailChan

	// for improving error handling ===============================================
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	App.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	App.ErrorLog = errorLog
	// ==============================================================================

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = App.InProduction

	App.Session = session

	// connect to database ============================================================
	fmt.Println("Connection to database...")
	db, err := drivers.ConnectSQL("root:password@tcp(127.0.0.1:3306)/bookings?parseTime=true")
	if err != nil {
		log.Fatal("Cannot connect to the database, Dying...!")
	}
	fmt.Println("Connected to database!")

	// defer db.SQL.Close()
	// ==================================================================================

	// http.HandleFunc("/", handlers.Home)
	// http.HandleFunc("/about", handlers.About)
	repo := handlers.NewRepo(&App, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&App)
	helpers.NewHelpers(&App)

	return db, nil
}

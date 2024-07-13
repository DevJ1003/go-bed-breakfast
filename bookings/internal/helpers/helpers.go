package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/devj1003/bookings/internal/config"
)

var App *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	App = a
}

func ClientError(w http.ResponseWriter, status int) {
	App.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {

	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	App.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func IsAuthenticate(r *http.Request) bool {

	exists := App.Session.Exists(r.Context(), "user_id")
	return exists
}

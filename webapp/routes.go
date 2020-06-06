package webapp

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes ...
func (app *App) Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", app.Index).Methods("GET", "POST")

	fs := http.FileServer(http.Dir(app.StaticDir))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	//r.PathPrefix("/").HandlerFunc(app.DefaultHandler)

	// Return the serve mux.
	return r
}

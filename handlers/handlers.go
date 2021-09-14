package handlers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// templates initialize available html template in templates/ directory.
var templates = template.Must(template.ParseGlob("templates/*"))

// HomePageHandler serves html and render home page for URL Shortener website.
func HomePageHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// NewURLHandler
func NewURLHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

// NotFoundHandler serves html and render 404 Page Not Found to unknown routes.
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := templates.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

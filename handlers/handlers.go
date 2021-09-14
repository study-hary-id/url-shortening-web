package handlers

import (
	"fmt"
	"github.com/study-hary-id/url-shortening-web/base62"
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

// ShortenerHandler serves html and shows the shortened URL on the website.
func ShortenerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		body    = r.FormValue("url")
		encoded = base62.StdEncoding.EncodeToString([]byte(body))
	)
	// TODO: Write data to JSON file
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)
	_, err := fmt.Fprintf(w, "<h3>%s</h3>", encoded)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// NotFoundHandler serves html and render 404 Page Not Found to unknown routes.
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := templates.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

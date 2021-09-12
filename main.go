package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	var (
		router = httprouter.New()
		PORT   = ":8000"
	)

	// GET / <root> show home page of the url-shortening-web.
	//router.GET("/", homePageHandler)

	// GET /shortener show shortened URL and the metadata.
	//router.GET("/shortener", shortenerHandler)

	// POST /new create a shorten URL.
	//router.POST("/new", newURLHandler)

	// GET /:url redirect to a shorten URL.
	//router.GET("/:url", getRedirectHandler)

	s := &http.Server{
		Addr:         PORT,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	fmt.Printf("Server listening at http://localhost%s\n", PORT)
	log.Fatal(s.ListenAndServe())
}

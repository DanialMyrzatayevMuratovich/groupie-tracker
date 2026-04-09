package main

import (
	"groupie-tracker/internal/handlers"
	"groupie-tracker/internal/updater"
	"log"
	"net/http"
	"time"
)

func main() {
	updater.Start(5 * time.Minute)

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)
	http.HandleFunc("/search", handlers.SearchHandler)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	log.Println("Server started at http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

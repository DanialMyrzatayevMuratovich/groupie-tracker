package handlers

import (
	"groupie-tracker/internal/controller"
	"html/template"
	"log"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	results, err := controller.SearchArtists(query)
	if err != nil {
		log.Println("search error:", err)
		RenderError(w, http.StatusInternalServerError, "Search failed")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Template error")
		return
	}

	tmpl.Execute(w, results)
}

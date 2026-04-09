package handlers

import (
	"groupie-tracker/internal/controller"
	"html/template"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, http.StatusNotFound, "Page not found")
		return
	}

	artists, err := controller.GetArtists()
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to load artists")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Template error")
		return
	}

	tmpl.Execute(w, artists)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		RenderError(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	artist, err := controller.GetArtistByID(id)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Failed to load artist")
		return
	}
	if artist == nil {
		RenderError(w, http.StatusNotFound, "Artist not found")
		return
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Template error")
		return
	}

	tmpl.Execute(w, artist)
}

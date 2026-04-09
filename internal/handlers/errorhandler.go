package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)

	tmpl, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		fmt.Fprintf(w, "%d — %s", status, message)
		return
	}

	tmpl.Execute(w, struct {
		Status  int
		Message string
	}{status, message})
}

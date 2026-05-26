package handlers

import (
	"fieldnotes/internal/content"
	"net/http"
)

// HelpHandler devuelve el handler de la página de ayuda.
func HelpHandler(tecnicas []content.Tecnica) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "help", PageData{
			CurrentPage: "help",
			Tecnicas:    tecnicas,
		})
	})
}

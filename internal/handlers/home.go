package handlers

import (
	"fieldnotes/internal/content"
	"net/http"
)

// HomeHandler devuelve el handler de la página de inicio.
// Recibe el slice de técnicas cargadas al arrancar el servidor.
func HomeHandler(tecnicas []content.Tecnica) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			renderNotFound(w, r)
			return
		}
		render(w, r, "home", PageData{
			CurrentPage: "home",
			Tecnicas:    tecnicas,
		})
	})
}

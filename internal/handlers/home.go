package handlers

import (
	"fieldnotes/internal/content"
	"net/http"
)

// HomeHandler devuelve el handler de la página de inicio.
func HomeHandler(tecnicas []content.Tecnica) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			renderNotFound(w, r)
			return
		}

		q := r.URL.Query().Get("q")
		filtros := FiltrosActivos{
			Clasificacion: r.URL.Query()["clasificacion"],
			Fases:         r.URL.Query()["fases"],
			Recursos:      r.URL.Query()["recursos"],
			Duracion:      r.URL.Query()["duracion"],
		}

		resultado := content.Filtrar(tecnicas, q, filtros.Clasificacion, filtros.Fases, filtros.Recursos, filtros.Duracion)

		render(w, r, "home", PageData{
			CurrentPage: "home",
			Tecnicas:    resultado,
			Query:       q,
			Filtros:     filtros,
		})
	})
}

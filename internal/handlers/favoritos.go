package handlers

import (
	"fieldnotes/internal/content"
	"net/http"
	"strings"
)

// FavoritosHandler devuelve el handler de la página de favoritos.
// Recibe los IDs guardados como parámetro GET: /favoritos?ids=card-sorting,think-aloud
// Si no hay IDs, muestra la página vacía.
func FavoritosHandler(tecnicas []content.Tecnica) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idsParam := r.URL.Query().Get("ids")
		var favoritas []content.Tecnica

		if idsParam != "" {
			ids := strings.Split(idsParam, ",")
			for _, id := range ids {
				t := content.FindByID(tecnicas, strings.TrimSpace(id))
				if t != nil {
					favoritas = append(favoritas, *t)
				}
			}
		}

		render(w, r, "favoritos", PageData{
			CurrentPage: "favoritos",
			Tecnicas:    favoritas,
		})
	})
}

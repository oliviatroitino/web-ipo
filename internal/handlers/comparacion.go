package handlers

import (
	"fieldnotes/internal/content"
	"net/http"
	"strings"
)

// ComparacionHandler gestiona /tecnicas/comparar?a={id}&b={id}.
// Con solo ?a muestra el selector de segunda técnica.
// Con ?a y ?b muestra la comparación completa.
func ComparacionHandler(tecnicas []content.Tecnica) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idA := strings.TrimSpace(r.URL.Query().Get("a"))
		idB := strings.TrimSpace(r.URL.Query().Get("b"))

		tecnicaA := content.FindByID(tecnicas, idA)
		if tecnicaA == nil {
			renderNotFound(w, r)
			return
		}

		var tecnicaB *content.Tecnica
		if idB != "" {
			tecnicaB = content.FindByID(tecnicas, idB)
			if tecnicaB == nil {
				renderNotFound(w, r)
				return
			}
		}

		render(w, r, "comparacion", PageData{
			CurrentPage: "",
			Tecnica:     tecnicaA,
			Tecnica2:    tecnicaB,
			Tecnicas:    tecnicas, // para el selector
		})
	})
}

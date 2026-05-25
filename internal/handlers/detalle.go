package handlers

import (
	"fieldnotes/internal/content"
	"net/http"
)

// DetalleHandler devuelve el handler de la ficha de una técnica.
// La ruta debe registrarse como /tecnicas/{id} para que PathValue funcione.
func DetalleHandler(tecnicas []content.Tecnica) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		t := content.FindByID(tecnicas, id)
		if t == nil {
			renderNotFound(w, r)
			return
		}
		render(w, r, "detalle", PageData{
			CurrentPage: "home",
			Tecnica:     t,
			Tecnicas:    tecnicas, // necesario para el selector de comparación
		})
	})
}

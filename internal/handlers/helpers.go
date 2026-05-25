package handlers

import (
	"fieldnotes/internal/content"
	"html/template"
	"net/http"
)

// PageData contiene los datos disponibles en todos los templates.
type PageData struct {
	CurrentPage string           // "home" | "help" | "detalle" | "favoritos"
	Tecnicas    []content.Tecnica  // listado de técnicas (home, favoritos)
	Tecnica     *content.Tecnica   // técnica individual (detalle, comparar)
	Tecnica2    *content.Tecnica   // segunda técnica (comparar)
	Query       string             // texto del buscador
	Filtros     FiltrosActivos     // filtros seleccionados
}

// FiltrosActivos recoge los valores de los cuatro filtros del listado.
type FiltrosActivos struct {
	Clasificacion []string
	Fases         []string
	Recursos      []string
	Duracion      []string
}

// render combina layout.html con la plantilla de página indicada y ejecuta "layout".
// Si algo falla devuelve 500 sin exponer detalles al navegador.
func render(w http.ResponseWriter, r *http.Request, page string, data PageData) {
	tmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/"+page+".html",
	)
	if err != nil {
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
	}
}

// renderNotFound sirve la página 404 con el código de estado correcto.
func renderNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	render(w, r, "404", PageData{})
}

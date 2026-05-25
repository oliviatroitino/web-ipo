package main

import (
	"fieldnotes/internal/content"
	"log"
	"net/http"
)

func main() {
	// Cargar técnicas al arrancar
	tecnicas, err := content.LoadTecnicas("content/tecnicas")
	if err != nil {
		log.Fatalf("Error cargando técnicas: %v", err)
	}
	log.Printf("Técnicas cargadas (%d):", len(tecnicas))
	for _, t := range tecnicas {
		log.Printf("  - %s [%s]", t.Nombre, t.Clasificacion)
	}

	mux := http.NewServeMux()

	// Archivos estáticos
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rutas (handlers a completar en pasos siguientes)
	mux.HandleFunc("/", homeHandler)

	log.Println("Servidor arrancado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Fieldnotes — OK"))
}

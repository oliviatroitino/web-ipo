package main

import (
	"fieldnotes/internal/content"
	"fieldnotes/internal/handlers"
	"log"
	"net/http"
)

func main() {
	// Cargar técnicas al arrancar
	tecnicas, err := content.LoadTecnicas("content/tecnicas")
	if err != nil {
		log.Fatalf("Error cargando técnicas: %v", err)
	}
	log.Printf("Técnicas cargadas (%d)", len(tecnicas))

	mux := http.NewServeMux()

	// Archivos estáticos
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rutas
	mux.Handle("/", handlers.HomeHandler(tecnicas))
	mux.Handle("/tecnicas/{id}", handlers.DetalleHandler(tecnicas))

	log.Println("Servidor arrancado en http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Archivos estáticos
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rutas (handlers a implementar en pasos siguientes)
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

package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func EmailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Email Handler")
}

// SearchHandler maneja las solicitudes de búsqueda de correos electrónicos.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	// Aquí implementarías la lógica para buscar correos electrónicos.
	// Por ahora, simplemente respondemos con el término de búsqueda.
	w.Write([]byte("Buscando correos para: " + query))
}

// Configura las rutas relacionadas con correos electrónicos.
func RegisterRoutes(r chi.Router) {
	r.Get("/search", SearchHandler)
}

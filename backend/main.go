package main

import (
	"backend/handlers"
	"net/http"

	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	// CORS configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler

	r.Use(corsHandler)

	handlers.RegisterRoutes(r)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}

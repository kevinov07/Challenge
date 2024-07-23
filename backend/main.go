package main

import (
	"backend/handlers"
	"net/http"

	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/search", handlers.SearchHandler)
	handlers.RegisterRoutes(r)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}

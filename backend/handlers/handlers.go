package handlers

import (
	"backend/constants"
	"backend/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

// SearchHandler maneja las solicitudes de búsqueda de correos electrónicos.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	term := chi.URLParam(r, "term")
	if term == "" {
		http.Error(w, "Search term is required", http.StatusBadRequest)
		return
	}

	pageStr := chi.URLParam(r, "page")
	page, err := strconv.Atoi(pageStr)

	if err != nil || page <= 0 {
		page = 20
	}

	query := fmt.Sprintf(constants.SEARCH_QUERY, term, page)

	body, StatusCode, err := services.SearchRequest(query)
	if err != nil {
		handleRequestError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCode)
	w.Write(body)
}

func GetEmailHandler(w http.ResponseWriter, r *http.Request) {
	emailId := chi.URLParam(r, "id")
	if emailId == "" {
		http.Error(w, "Email ID is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(constants.EMAIL_QUERY, emailId)

	body, statusCode, err := services.SearchRequest(query)
	if err != nil {
		handleRequestError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)

}

func handleRequestError(w http.ResponseWriter, err error) {
	var status int
	switch {
	case strings.Contains(err.Error(), "timeout"):
		status = http.StatusGatewayTimeout
	case strings.Contains(err.Error(), "invalid URL"):
		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
	}
	http.Error(w, err.Error(), status)
}

// Configura las rutas relacionadas con correos electrónicos.
func RegisterRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the email indexer API"))
	})
	r.Get("/search/{term}-{page}", SearchHandler)
	r.Get("/email/{id}", GetEmailHandler)
}

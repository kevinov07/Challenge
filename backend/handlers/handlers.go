package handlers

import (
	"backend/services"
	"fmt"
	"net/http"
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

	page := r.URL.Query().Get("page")
	if page == "" || page == "0" {
		page = "20"
	}

	query := fmt.Sprintf(`{
        "search_type": "match",
        "query": {
            "term": "%s",
        },
        "from": 0,
        "max_results": %s,
        "_source": []
    }`, term, page)

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
	fmt.Println(emailId)
	if emailId == "" {
		http.Error(w, "Email ID is required", http.StatusBadRequest)
		return
	}

	query := fmt.Sprintf(`{
		"search_type": "matchphrase",
		"query": {
			"term": "%s",
			"field": "message_id"
		},
		"sort_fields": [],
		"from": 0,
		"max_results": 1,
		"_source": []
	}`, emailId)

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

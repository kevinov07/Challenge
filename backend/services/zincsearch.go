package services

import (
	"io"
	"net/http"
	"strings"
)

// SearchRequest realiza una solicitud de búsqueda a ZincSearch.
func SearchRequest(query string) ([]byte, int, error) {
	req, err := http.NewRequest("POST", "http://localhost:4080/api/email-indexer/_search", strings.NewReader(query))
	if err != nil {
		return nil, 0, err
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

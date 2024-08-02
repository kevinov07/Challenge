package services

import (
	"backend/constants"
	"io"
	"net/http"
	"strings"
)

// SearchRequest realiza una solicitud de búsqueda a ZincSearch.
func SearchRequest(query string) ([]byte, int, error) {
	url := constants.ZINCSEARCH_PORT + constants.ZINCSEARCH_ENDPOINT

	req, err := http.NewRequest("POST", url, strings.NewReader(query))
	if err != nil {
		return nil, 0, err
	}
	req.SetBasicAuth(constants.USERNAME, constants.PASSWORD)
	req.Header.Set("Content-Type", "application/json")

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

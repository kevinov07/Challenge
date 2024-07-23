package models

type Email struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Sender  string `json:"sender"`
	// Otros campos necesarios
}

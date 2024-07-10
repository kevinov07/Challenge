package models

import "time"

// Email represents an email structure to be indexed
type Email struct {
	MessageID               string    `json:"message_id"`
	Date                    time.Time `json:"date"`
	DateSubEmail            string    `json:"date_subemail"`
	From                    string    `json:"from"`
	To                      string    `json:"to"`
	Sent                    string    `json:"sent"`
	Cc                      string    `json:"cc"`
	XFrom                   string    `json:"x_from"`
	XTo                     string    `json:"x_to"`
	XCc                     string    `json:"x_cc"`
	XBcc                    string    `json:"x_bcc"`
	XFolder                 string    `json:"c_folder"`
	XOrigin                 string    `json:"x_origin"`
	XFileName               string    `json:"x_file_name"`
	Subject                 string    `json:"subject"`
	MimeVersion             string    `json:"mime_version"`
	ContentType             string    `json:"content_type"`
	ContentTransferEncoding string    `json:"content_transfer_encoding"`
	Body                    string    `json:"body"`
}

package models

import "time"

// Email represents an email structure to be indexed
type Email struct {
	Body                    string    `json:"body"`
	CFolder                 string    `json:"c_folder"`
	CC                      string    `json:"cc"`
	ContentTransferEncoding string    `json:"content_transfer_encoding"`
	ContentType             string    `json:"content_type"`
	Date                    time.Time `json:"date"`
	From                    string    `json:"from"`
	MessageID               string    `json:"message_id"`
	MimeVersion             string    `json:"mime_version"`
	Replies                 []string  `json:"replies"`
	Sent                    time.Time `json:"sent"`
	Subject                 string    `json:"subject"`
	To                      string    `json:"to"`
	XBcc                    string    `json:"x_bcc"`
	XCc                     string    `json:"x_cc"`
	XFileName               string    `json:"x_file_name"`
	XFrom                   string    `json:"x_from"`
	XOrigin                 string    `json:"x_origin"`
	XTo                     string    `json:"x_to"`
}

type Replies struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Send    string `json:"send"`
	Subject string `json:"subject"`
	Cc      string `json:"cc"`
	Content string `json:"content"`
}

package models

type IndexRequestBody struct {
	Name        string   `json:"name"`
	StorageType string   `json:"storage_type"`
	ShardNum    int      `json:"shard_num"`
	Mappings    Mappings `json:"mappings"`
}

type ZincRequestBody struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}

type FieldProperties struct {
	Type           string `json:"type"`
	Index          bool   `json:"index"`
	Store          bool   `json:"store"`
	Sortable       bool   `json:"sortable"`
	Hightlightable bool   `json:"highlightable"`
	Aggregatable   bool   `json:"aggregatable"`
}

type EmailMappings struct {
	MessageID               FieldProperties `json:"message_id"`
	Date                    FieldProperties `json:"date"`
	DateSubEmail            FieldProperties `json:"date_subemail"`
	From                    FieldProperties `json:"from"`
	To                      FieldProperties `json:"to"`
	Sent                    FieldProperties `json:"sent"`
	Cc                      FieldProperties `json:"cc"`
	XFrom                   FieldProperties `json:"x_from"`
	XTo                     FieldProperties `json:"x_to"`
	XCc                     FieldProperties `json:"x_cc"`
	XBcc                    FieldProperties `json:"x_bcc"`
	XFolder                 FieldProperties `json:"x_folder"`
	XOrigin                 FieldProperties `json:"x_origin"`
	XFileName               FieldProperties `json:"x_file_name"`
	Subject                 FieldProperties `json:"subject"`
	MimeVersion             FieldProperties `json:"mime_version"`
	ContentType             FieldProperties `json:"content_type"`
	ContentTransferEncoding FieldProperties `json:"content_transfer_encoding"`
	Body                    FieldProperties `json:"body"`
	CFolder                 FieldProperties `json:"c_folder"`
}

type Mappings struct {
	Properties EmailMappings `json:"properties"`
}

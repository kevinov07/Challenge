package constants

import (
	"email-indexer/models"
)

const FOLDER_PATH = "../../enron_mail_20110402/maildir" //"C:/Users/kevin/Desktop/enron_mail_20110402/maildir"
const DATE_FORMAT = "Mon, 2 Jan 2006 15:04:05 -0700 (MST)"
const DATE_FORMAT2 = "Mon, 02 Jan 2006 15:04:05 -0700 (MST)"
const SERVER = "http://localhost:4080"
const ENDPOINT = "/api/_bulkv2"
const ENDPOINT_INDEX = "/api/index"
const USERNAME = "admin"
const PASSWORD = "Complexpass#123"
const TOTAL_EMAILS = 517424
const FOLDER_PATH_TEST = "../../enron_mail_20110402/maildir/arnold-j"

var IndexConfig = models.IndexRequestBody{
	Name:        "email-indexer",
	StorageType: "disk",
	ShardNum:    3,
	Mappings: models.Mappings{
		Properties: models.EmailMappings{
			MessageID:               PROPERTIES_VALUES,
			Date:                    PROPERTIES_VALUES,
			DateSubEmail:            PROPERTIES_VALUES,
			From:                    PROPERTIES_VALUES,
			To:                      PROPERTIES_VALUES,
			Sent:                    PROPERTIES_VALUES,
			Cc:                      PROPERTIES_VALUES,
			XFrom:                   PROPERTIES_VALUES,
			XTo:                     PROPERTIES_VALUES,
			XCc:                     PROPERTIES_VALUES,
			XBcc:                    PROPERTIES_VALUES,
			XFolder:                 PROPERTIES_VALUES,
			XOrigin:                 PROPERTIES_VALUES,
			XFileName:               PROPERTIES_VALUES,
			Subject:                 PROPERTIES_VALUES,
			MimeVersion:             PROPERTIES_VALUES,
			ContentType:             PROPERTIES_VALUES,
			ContentTransferEncoding: PROPERTIES_VALUES,
			Body:                    PROPERTIES_VALUES,
			CFolder:                 PROPERTIES_VALUES,
		},
	},
}

var PROPERTIES_VALUES = models.FieldProperties{
	Type:           "text",
	Index:          true,
	Store:          false,
	Sortable:       false,
	Hightlightable: false,
	Aggregatable:   false,
}

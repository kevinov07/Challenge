package services

import (
	"email-indexer/constants"
	"email-indexer/models"
	"reflect"
	"strings"
	"time"
)

var HEADER_MAP = map[string]func(*models.Email, string){
	"Message-ID:":                assignToField("MessageID"),
	"Sent:":                      assignToField("Sent"),
	"Cc:":                        assignToField("Cc"),
	"X-From:":                    assignToField("XFrom"),
	"X-To:":                      assignToField("XTo"),
	"X-cc:":                      assignToField("XCc"),
	"X-bcc:":                     assignToField("XBcc"),
	"X-Folder:":                  assignToField("XFolder"),
	"X-Origin:":                  assignToField("XOrigin"),
	"X-FileName:":                assignToField("XFileName"),
	"Date:":                      parseDate,
	"From:":                      assignToField("From"),
	"To:":                        appendToField("To"),
	"Subject:":                   assignToField("Subject"),
	"Mime-Version:":              assignToField("MimeVersion"),
	"Content-Type:":              assignToField("ContentType"),
	"Content-Transfer-Encoding:": assignToField("ContentTransferEncoding"),
}

func assignToField(field string) func(*models.Email, string) {
	return func(email *models.Email, value string) {
		// Removing extra spaces and assigning value to the field
		value = RemoveSpaces(value)
		v := reflect.ValueOf(email).Elem().FieldByName(field)
		if v.IsValid() && v.CanSet() && v.Kind() == reflect.String {
			v.SetString(value)
		}
	}
}

func appendToField(field string) func(*models.Email, string) {
	return func(email *models.Email, value string) {
		// Removing extra spaces and appending value to the field
		value = RemoveSpaces(value)
		v := reflect.ValueOf(email).Elem().FieldByName(field)
		if v.IsValid() && v.CanSet() && v.Kind() == reflect.String {
			v.SetString(v.String() + value + " ")
		}
	}
}

func parseDate(email *models.Email, value string) {
	dateStr := RemoveSpaces(value)
	email.Date, _ = time.Parse(constants.DATE_FORMAT, dateStr)
	email.DateSubEmail = dateStr
}

func RemoveSpaces(str string) string {
	value := strings.Join(strings.Fields(str), " ")
	return value
}

package services

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"email-indexer/constants"
	"email-indexer/models"
)

var totalEmails []models.Email

func CreateZincIndex() {

	jsonData, err := json.Marshal(constants.IndexConfig)

	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return
	}

	//log.Println("JSON:", string(jsonData))

	URL := constants.SERVER + constants.ENDPOINT_INDEX
	req, reqError := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))

	if reqError != nil {
		log.Println(reqError)
		return
	}

	req.SetBasicAuth(constants.USERNAME, constants.PASSWORD)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, resError := client.Do(req)
	if resError != nil {
		log.Println(resError)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error al leer el cuerpo de la respuesta:", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		err := json.Unmarshal(body, &errorResponse)
		if err != nil {
			log.Println("Error unmarshalling error response:", err)
			log.Println("Error indexing Zinc:", string(body))
			return
		}

		errorMessage, ok := errorResponse["error"].(string)
		if !ok {
			log.Println("Error indexing Zinc: respuesta de error no valida")
			return
		}
		log.Println("Error indexing Zinc:", errorMessage)
		return
	}

	fmt.Println("Zinc indexed successfully")
}

func ReadEmail(path string) {

	file, err := os.Open(path)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	email := models.Email{}

	for scanner.Scan() {
		line := scanner.Text()
		ParseLine(&email, line)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading file:", err)
	}

	addEmail(&email)
}

func CheckFolder(folderPath string) {
	if len(totalEmails) > 20000 {
		return
	}
	fmt.Println("Checking folders")

	files, filErr := os.ReadDir(folderPath)
	if filErr != nil {
		log.Println("Error reading folder", filErr)
		return
	}
	for _, file := range files {
		pathToCheck := filepath.Join(folderPath, file.Name())
		if file.IsDir() {
			CheckFolder(pathToCheck)
		} else {
			ReadEmail(pathToCheck)
		}
	}

}

func IndexEmails(emails []models.Email) {
	log.Println("Indexing emails")

	body := models.ZincRequestBody{
		Index:   "email-indexer",
		Records: emails,
	}

	jsonData, err := json.Marshal(body)
	//log.Println(jsonData)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return
	}

	URL := constants.SERVER + constants.ENDPOINT
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	//log.Println("Response:", req.Body)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	req.SetBasicAuth(constants.USERNAME, constants.PASSWORD)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error indexing emails:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Error indexing email:", resp.Status)
		return
	}
	log.Println("Emails indexed successfully", resp.Status)
}

func ParseDate(email *models.Email, value string) {
	dateStr := strings.TrimSpace(value)
	parsedDate, err := time.Parse(constants.DATE_FORMAT, dateStr)
	if err != nil {
		log.Println("Error parsing date:", err)
	}
	email.Date = parsedDate
	email.DateSubEmail = dateStr
}

func ParseLine(email *models.Email, line string) {
	if strings.Contains(line, "=======================================") {
		return
	}
	// else if strings.Contains(line, "---------------------- Forwarded by") {
	// 	return
	// }

	for prefix, action := range HEADER_MAP {
		if strings.HasPrefix(line, prefix) {
			action(email, strings.TrimPrefix(line, prefix))
			return
		}
	}

	email.Body += line + "\n"
}

func addEmail(emailPointer *models.Email) {
	//log.Println("Adding email to the list")
	totalEmails = append(totalEmails, *emailPointer)
	if len(totalEmails) == 20000 {
		IndexEmails(totalEmails)
	}

}

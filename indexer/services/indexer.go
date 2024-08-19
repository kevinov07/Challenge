package services

import (
	"bufio"
	"bytes"
	"email-indexer/constants"
	"email-indexer/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	mu          sync.Mutex // Mutex to protect access to the totalEmails slice
	totalEmails []models.Email
	numEmails   int
	client      = &http.Client{}
	maxEmails   = 50000
)

func CreateZincIndex() error {

	url := constants.SERVER + constants.ENDPOINT_INDEX
	jsonData, err := json.Marshal(constants.IndexConfig)
	if err != nil {
		return fmt.Errorf("failed to marshal index config: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create new request: %v", err)
	}
	req.SetBasicAuth(constants.USERNAME, constants.PASSWORD)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create index: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create index, status: %s", resp.Status)
	}

	log.Println("Index created successfully")
	return nil
}

// CheckFolder recorre el directorio y procesa los archivos utilizando concurrencia
func CheckFolder(folderPath string, wg *sync.WaitGroup, fileChan chan<- string, sem chan struct{}) {
	defer wg.Done()

	files, filErr := os.ReadDir(folderPath)
	if filErr != nil {
		log.Println("Error reading folder", filErr)
		return
	}

	for _, file := range files {
		pathToCheck := filepath.Join(folderPath, file.Name())
		if file.IsDir() {
			wg.Add(1)
			go CheckFolder(pathToCheck, wg, fileChan, sem)
		} else {
			sem <- struct{}{}
			wg.Add(1)
			go func(path string) {
				defer func() {
					wg.Done()
					<-sem
				}()
				ReadEmail(path, wg)
			}(pathToCheck)
		}
	}

}

// ReadEmail procesa un archivo de correo electrónico
func ReadEmail(path string, wg *sync.WaitGroup) {

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

func IndexEmails(emails []models.Email) {
	log.Println("Indexing emails")

	body := models.ZincRequestBody{
		Index:   "email-indexer",
		Records: emails,
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return
	}

	URL := constants.SERVER + constants.ENDPOINT
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	req.SetBasicAuth(constants.USERNAME, constants.PASSWORD)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error indexing emails:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error indexing email: %s\n", resp.Status)
		return
	}
	log.Println("Emails indexed successfully", resp.Status)
}

func addEmail(emailPointer *models.Email) {
	mu.Lock()
	totalEmails = append(totalEmails, *emailPointer)
	mu.Unlock()

	if len(totalEmails)%maxEmails == 0 {
		log.Println("Emails:", len(totalEmails))
	}
}

func ProcessEmails(folderPath string) {
	var wg sync.WaitGroup
	fileChan := make(chan string, 1000)
	sem := make(chan struct{}, 1000)

	// go func() {
	// 	for path := range fileChan {
	// 		sem <- struct{}{}
	// 		wg.Add(1)
	// 		go func(path string) {
	// 			defer func() {
	// 				wg.Done()
	// 				<-sem
	// 			}()
	// 			ReadEmail(path, &wg)
	// 		}(path)
	// 	}
	// }()

	// Explore folder and send files to fileChan
	wg.Add(1)
	go CheckFolder(folderPath, &wg, fileChan, sem)

	// Wait for all
	go func() {
		wg.Wait()
		close(fileChan)
	}()

	// Wait for all files to be processed
	wg.Wait()

	if len(totalEmails) == constants.TOTAL_EMAILS {
		numEmails = len(totalEmails)
		IndexEmails(totalEmails)
	}

	log.Println("All files processed")
	log.Println("Total emails:", numEmails)

}

func ParseLine(email *models.Email, line string) {

	for prefix, action := range HEADER_MAP {
		if strings.HasPrefix(line, prefix) {
			action(email, strings.TrimPrefix(line, prefix))
			return
		}
	}

	email.Body += line + "\n"
}

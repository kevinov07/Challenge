// package main

// import (
// 	"email-indexer/constants"
// 	indexer "email-indexer/services"
// )

// func main() {
// 	indexer.ReadEmail(constants.FOLDER_PATH)
// }

package main

import (
	"email-indexer/constants"
	"email-indexer/services"
	"fmt"
	"time"
)

func main() {
	//services.CreateZincIndex()

	start := time.Now()

	fmt.Println("Indexing emails...")
	services.CheckFolder(constants.FOLDER_PATH)
	elapsed := time.Since(start)

	fmt.Println("Indexing completed in", elapsed)

	// estimatedTime := elapsed * (constants.TOTAL_EMAILS / 20000)
	// fmt.Println("Estimated time to index all emails:", estimatedTime)

}

package main

import (
	"email-indexer/constants"
	"email-indexer/services"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	//services.CreateZincIndex()

	// con 20000 emails tarda aprox 3 minutos

	//con FOLDER_PATH_TEST tarda 1.3 segundos
	//con FOLDER_PATH Indexing completed in 23m50.5626565s

	//Habilitar el perfilado CPU
	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	memProfile, err := os.Create("mem_profile.prof")
	if err != nil {
		fmt.Println("Could not create memory profile: ", err)
	}
	defer memProfile.Close()
	defer pprof.WriteHeapProfile(memProfile)

	start := time.Now()

	services.ProcessEmails(constants.FOLDER_PATH)
	elapsed := time.Since(start)

	fmt.Println("Indexing completed in", elapsed)
}

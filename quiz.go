package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func readCsvFile(filepath string) {
	// Open the file
	csvFile, err := os.Open(filepath) //return a *FILE and an ERROR
	if err != nil {
		log.Fatal("Unable to read the quiz file provided!", err)
	}

	// Parse the file
	reader := csv.NewReader(csvFile) //return a *READER

	// Iterate through the lines
	for {
		arrayLine, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("question : %s ## Answer %s\n", arrayLine[0], arrayLine[1])

	}

}

func main() {
	fmt.Printf("hello\n")
	readCsvFile("problems.csv")

}

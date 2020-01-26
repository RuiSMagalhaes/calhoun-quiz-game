package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer' ")
	flag.Parse()

	readCsvFile(csvFilename)

}

func readCsvFile(filename *string) {
	// Open the file
	csvFile, err := os.Open(*filename) //return a *FILE and an ERROR
	if err != nil {
		log.Fatal("error: ", err)
	}

	// Parse the file
	reader := csv.NewReader(csvFile) //return a *READER

	// Iterate through the lines
	score := 0
	counter := 0

	for {
		arrayLine, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("question : %s ? \n", arrayLine[0])

		var i string
		fmt.Scanf("%s\n", &i)

		if i == strings.TrimSpace(arrayLine[1]) {
			score++
		}

		counter++

	}

	fmt.Printf("Your final score is: %v/%v !!\n", score, counter)

}

// TODO :
//		* add a timer
//		* add flag for timer with default of 30s
// 		* Add an option (a new flag) to shuffle the quiz order each time it is run.

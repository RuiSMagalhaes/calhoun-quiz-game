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
		_, error := fmt.Scanf("%s", &i)

		if error != nil {
			log.Print("Scan for i failed, due to ", error)
			return
		}

		if i == arrayLine[1] {
			score = score + 1
		}

		counter = counter + 1

	}

	fmt.Printf("Your final score is: %v/%v !!\n", score, counter)

}

func main() {
	readCsvFile("problems.csv")

}

// TODO :
//		* clean spaces and capitalizations of the user answer
// 		* add flag for filepath with a default of problems.csv
//		* add a timer
//		* add flag for timer with default of 30s
// 		* Add an option (a new flag) to shuffle the quiz order each time it is run.

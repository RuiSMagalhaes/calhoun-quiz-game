package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format of 'question,answer' ")
	timeLimit := flag.Int("limit", 30, "the time for the quiz in seconds")
	flag.Parse()

	readCsvFile(csvFilename, timeLimit)

}

func readCsvFile(filename *string, timeLimit *int) {
	// Open the file
	csvFile, err := os.Open(*filename) //return a *FILE and an ERROR
	if err != nil {
		log.Fatal("error: ", err)
	}

	// Parse the file
	reader := csv.NewReader(csvFile) //return a *READER

	// Iterate through the lines
	score := 0
	multiSlice, _ := reader.ReadAll()

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, arrayLine := range multiSlice {
		fmt.Printf("question %v: %s ? \n", i+1, arrayLine[0])
		answerChanel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChanel <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("Your final score is: %v/%v !!\n", score, len(multiSlice))
			return
		case answer := <-answerChanel:
			if answer == strings.TrimSpace(arrayLine[1]) {
				score++
			}
		}

	}

	fmt.Printf("Your final score is: %v/%v !!\n", score, len(multiSlice))

}

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type question struct {
	q string
	a string
}

func main() {
	filePath := flag.String("file", "problems.csv", "path to quiz problems")
	timeLimit := flag.Int("limit", 1, "time limit")
	flag.Parse()
	csvFile, err := os.Open(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	numCorrect := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	csvData, err := csv.NewReader(csvFile).ReadAll()
outer:
	for i, line := range csvData {
		data := question{
			q: line[0],
			a: line[1],
		}
		fmt.Printf("Problem #%d: %s\n", i+1, data.q)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			answerCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("You got %d out of %d correct\n", numCorrect, len(csvData))
			break outer
		case answer := <-answerCh:
			if answer == data.a {
				fmt.Println("Correct!")
				numCorrect++
			} else {
				fmt.Println("Incorrect!")
			}
		}
	}

}

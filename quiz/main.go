package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A CSV file containing problems and solutions in the format'question,answer'")
	timeLimit := flag.Int("limit", 30, "Time limit to complete the quiz in seconds")
	flag.Parse()

	csvFile, err := os.Open(*csvFileName)
	if err != nil {
		raiseError(fmt.Sprintf("An error occurred while trying to open file: %s\n", *csvFileName))
	}

	r := csv.NewReader(csvFile)
	lines, err := r.ReadAll()
	if err != nil {
		raiseError("An error occurred while trying to read the CSV.")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		answerChannel := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerChannel <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerChannel:
			if answer == problem.answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, record := range lines {
		problems[i] = problem{
			question: record[0],
			answer:   strings.TrimSpace(record[1]),
		}
	}
	return problems
}

func raiseError(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

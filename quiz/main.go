package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A CSV file containing problems and solutions in the format'question,answer'")
	flag.Parse()
	csvFile, err := os.Open(*csvFileName)
	if err != nil {
		raiseError(fmt.Sprintf("An error occurred while trying to open file: %s\n", *csvFileName))
	}

	lines := csv.NewReader(csvFile)
	records, err := lines.ReadAll()
	if err != nil {
		raiseError("An error occurred while trying to read the CSV.")
	}

	problems := parseLines(records)
	var correct int
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s\n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == problem.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func parseLines(records [][]string) []problem {
	problems := make([]problem, len(records))
	for i, record := range records {
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

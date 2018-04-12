package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type quiz struct {
	questions []question
	grade     float32
}

type question struct {
	text            string
	correctAnswer   string
	attemptedAnswer string
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var reader *csv.Reader

	if file, err := os.Open("./problems.csv"); err != nil {
		fmt.Println(err)
	} else {
		reader = csv.NewReader(file)
	}

	records, err := reader.ReadAll()
	checkError(err)

	// Prep our quiz data structure.
	quiz := quiz{
		questions: make([]question, len(records)),
	}

	for i, line := range records {
		quiz.questions[i] = question{text: line[0], correctAnswer: strings.TrimSpace(line[1])}
	}

	lineReader := bufio.NewReader(os.Stdin)
	answeredCorrectly := 0

	// Now, spin through the quesitons we have, and prompt the user.
	for i, question := range quiz.questions {
		fmt.Printf("%d) %s: ", i+1, question.text)

		// Prompt for their answer.
		answer, err := lineReader.ReadString('\n')
		checkError(err)

		question.attemptedAnswer = strings.TrimSpace(answer)

		if question.attemptedAnswer == question.correctAnswer {
			answeredCorrectly++
		}
	}

	fmt.Printf("You answered %d out of %d questions correctly\n", answeredCorrectly, len(quiz.questions))
}

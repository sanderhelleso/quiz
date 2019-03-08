package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Quiz starts the quiz and print out the new question
// after each guess, accumulating the total score
// returns the total time, score and percentage correct
func Quiz(questions *[]Question) {

	// initialize score
	var score int

	// initialize a new reader
	reader := bufio.NewReader(os.Stdin)

	// iterate over all the questions and print
	// each question one by one and wait for answer
	for i, q := range *questions {
		fmt.Printf("Question %v\n%v = ", i+1, q.question)

		// get question from input (stdIn)
		answer, _ := reader.ReadString('\n')

		// trim input correctly to make it comparable
		answer = strings.ToLower(strings.Trim(answer, " \r\n"))

		// compare answer to input
		if strings.Compare(answer, q.answer) == 0 {
			score++ // increment score if correct
		}

		fmt.Println("\n------------------------------")
	}

	fmt.Printf("Total score: %v\n", score)
}
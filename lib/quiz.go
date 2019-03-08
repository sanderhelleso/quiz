package lib

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const message string = `
////////////////////////////////////////
// Welcome to the amazing quiz game! //
//////////////////////////////////////
`

// Quiz starts the quiz and print out the new question
// after each guess, accumulating the total score
// returns the total time, score and percentage correct
func Quiz(questions *[]Question) {

	// print out welcome screen
	fmt.Println(message + "\n")

	// initialize a new reader
	reader := bufio.NewReader(os.Stdin)

	// initialize score
	var score int

	// ask if questions should be shuffeled
	fmt.Print("Would you like to shuffle the questions? (yes/no): ")
	shouldShuffle, _ := reader.ReadString('\n')
	shouldShuffle = trimString(&shouldShuffle)

	// check answer, and shuffle if "yes"
	if strings.Compare(shouldShuffle, "yes") == 0 {
		shuffleSlice(questions)
	}

	// iterate over all the questions and print
	// each question one by one and wait for answer
	for i, q := range *questions {
		fmt.Printf("Question %v\n%v = ", i+1, q.question)

		// get question from input (stdIn)
		answer, _ := reader.ReadString('\n')

		// trim input correctly to make it comparable
		answer = trimString(&answer)

		// compare answer to input
		if strings.Compare(answer, q.answer) == 0 {
			score++ // increment score if correct
		}

		fmt.Println("\n------------------------------")
	}

	// print out the total score and the score in %
	fmt.Printf(
		"Your score: %v/%v (%v%%)\n",
		score,
		len(*questions),
		getPercentage(score, len(*questions)))
}

func setTimer() int {
	return 2
}

// getPercentage returns the percentage of total by max
func getPercentage(total int, max int) int {
	return int(float64(total) / float64(max) * 100)
}

func trimString(input *string) string {
	return strings.ToLower(strings.Trim(*input, " \r\n"))
}

// shuffle the passed in slice, modifying the
// original slice in 0(n)
func shuffleSlice(q *[]Question) {

	// set the seed for randomnezzz
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < len(*q); i++ {

		// get values by random indexes
		q1 := &(*q)[rand.Int()%len(*q)]
		q2 := &(*q)[rand.Int()%len(*q)]

		// swap values retrieved by random indexes
		swap(q1, q2)
	}
}

// swap swaps two values passed in by memory referance
func swap(q1, q2 *Question) {
	*q1, *q2 = *q2, *q1
}

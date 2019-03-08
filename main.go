package main

import (
	"./lib"
)

const csvFile string = "problems.csv"

func main() {

	// get quiz questions
	quizQuestions := lib.GetQuestionsFromCSV(csvFile)
	lib.Quiz(&quizQuestions)
}

package lib

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Question represent a quiz question
// with its corresponding answer
type Question struct {
	question string
	answer   string
}

// GetQuestionsFromCSV takes in a giveN filepath and reads the csv
// data from the given file, returns a slice of Questions
func GetQuestionsFromCSV(filePath string) []Question {

	// initialize questions array
	var questions []Question

	// open file using recieved filepath & handle potenisal error
	file, err := os.Open(filePath)

	// if any error like "cant find file" or the extension is
	// not csv, exit the program early and alert user
	if err != nil || filepath.Ext(filePath) != ".csv" {
		fmt.Println("Unable to read from file or invalid file fornat, exiting...")
		panic(err)
	}

	// create a new reader to read from file, line by line
	reader := csv.NewReader(bufio.NewReader(file))
	for {
		// read csv record from line
		record, err := reader.Read()

		// stop at end of file
		if err == io.EOF {
			break
		}

		// create new question
		var question = Question{
			record[0],
			record[1],
		}

		// add question to slice of questions
		questions = append(questions, question)
	}

	// return slice of questions retrieved fro array
	return questions
}

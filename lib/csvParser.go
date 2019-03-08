package lib

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ReadCSV takes in a gived filepath and reads the csv
// data from the given file
func ReadCSV(filePath string) {

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

		// print record
		fmt.Println(record)
	}
}

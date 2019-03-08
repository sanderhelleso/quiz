package main

import (
	"./lib"
)

const csvFile string = "problems.csv"

func main() {

	lib.ReadCSV(csvFile)
}

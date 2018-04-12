package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

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

	fmt.Println(records)
	fmt.Println("Hi, this is a go program")
}

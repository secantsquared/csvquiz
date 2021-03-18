package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	csvfile, err := os.Open("problems.csv")

	if err != nil {
		log.Fatalln("Could not open CSV file.")
	}

	r := csv.NewReader(csvfile)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("What is %s?\n", record[0])
	}
}

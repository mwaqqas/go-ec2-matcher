package main

import (
	"encoding/csv"
	"log"
	"os"
)

func CSVReader(filename string) (records [][]string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	// skip the header row
	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	// process the rest of the records
	records, err = r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return
}

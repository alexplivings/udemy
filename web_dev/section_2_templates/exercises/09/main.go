package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("table.csv")
	if err != nil {
		log.Fatal(err)
	}

	input := string(dat)

	r := csv.NewReader(strings.NewReader(input))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}

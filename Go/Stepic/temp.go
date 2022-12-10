package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type group struct {
	ID       int
	Number   string
	Year     int
	Students []student
}

type groupAvg struct {
	Average float64
}

func main() {
	var iGroup group
	var avgRating groupAvg

	iData, err := io.ReadAll(os.Stdin)
	if err != nil && err != io.EOF {
		panic("Cannot read stdin!")
	}

	if err := json.Unmarshal(iData, &iGroup); err != nil {
		panic("Cannot unmarshal json!")
	}

	for _, st := range iGroup.Students {
		avgRating.Average += float64(len(st.Rating))
	}

	avgRating.Average = avgRating.Average / float64(len(iGroup.Students))

	oData, err := json.MarshalIndent(avgRating, "", "    ")
	if err != nil {
		panic("Cannot marshall json!")
	}

	fmt.Printf("%s", oData)
}
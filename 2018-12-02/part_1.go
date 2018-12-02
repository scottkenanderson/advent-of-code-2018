package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func getBoxIds(filename string) []string {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var boxIds []string
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		boxIds = append(boxIds, line[0])
	}
	return boxIds
}

func parseBoxID(boxID string) (bool, bool) {
	var two, three = false, false
	var characterCount = make(map[byte]int)
	for i := 0; i < len(boxID); i++ {
		characterCount[boxID[i]]++
	}

	for _, v := range characterCount {
		if v == 2 && !two {
			two = true
		}
		if v == 3 {
			three = true
		}
	}
	return two, three
}

func main() {
	boxIds := getBoxIds("input_day_2.csv")
	var twoCount, threeCount int
	for _, boxID := range boxIds {
		isTwo, isThree := parseBoxID(boxID)
		if isTwo {
			twoCount++
		}
		if isThree {
			threeCount++
		}
	}
	fmt.Println(twoCount, "*", threeCount, "=", twoCount*threeCount)
}

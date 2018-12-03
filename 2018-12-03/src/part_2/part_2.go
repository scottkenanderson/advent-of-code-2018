package main

import (
	"bufio"
	"bytes"
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

func getStringDifference(boxID, other string) (bool, string) {
	diff := 0
	var b bytes.Buffer
	for i := 0; i < len(boxID); i++ {
		if boxID[i] == other[i] {
			b.WriteByte(boxID[i])
		} else {
			diff++
		}
	}
	return diff <= 1, b.String()
}

func findMatchingString(boxID string, boxIds []string) (bool, string) {
	for _, other := range boxIds {
		hasDifference, commonCharacters := getStringDifference(boxID, other)
		if hasDifference {
			return true, commonCharacters
		}
	}
	return false, ""
}

func getCommonLetters(boxIds []string) string {
	for i, boxID := range boxIds {
		found, matching := findMatchingString(boxID, boxIds[i+1:])
		if found {
			fmt.Println(boxID, matching)
			return matching
		}
	}
	return ""
}

func main() {
	boxIds := getBoxIds("../../input_day_2.csv")
	commonLetters := getCommonLetters(boxIds)
	fmt.Println(commonLetters)
}

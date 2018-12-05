package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func readPolymer(filename string) string {
	csvFile, _ := os.Open(filename)
	reader := bufio.NewScanner(csvFile)
	var lines []string
	for reader.Scan() {
		line := reader.Text()
		lines = append(lines, line)
	}
	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
	return lines[0]
}

type unit struct {
	string
}

func (u unit) isLower() bool {
	return u.string == strings.ToLower(u.string)
}

func (u unit) isUpper() bool {
	return u.string == strings.ToUpper(u.string)
}

func (u unit) isEqual(other unit) bool {
	return strings.ToLower(u.string) == strings.ToLower(other.string)
}

func (u unit) reactsWith(other unit) bool {
	return u.isEqual(other) && ((u.isLower() && other.isUpper()) || u.isUpper() && other.isLower())
}

func explode(polymer string) string {
	stringLength := len(polymer)
	var buffer bytes.Buffer
	for i := 0; i < stringLength; i++ {
		u := unit{string(polymer[i])}
		if i >= stringLength-1 {
			buffer.WriteString(u.string)
			continue
		}
		next := unit{string(polymer[i+1])}
		if u.reactsWith(next) {
			i++
		} else {
			buffer.WriteString(u.string)
		}
	}
	return buffer.String()
}

func getMinExplodedString(polymer string) int {
	l := len(polymer)
	newLen := 0
	for {
		exploded := explode(polymer)
		newLen = len(exploded)
		if l == newLen {
			break
		}
		l = newLen
		polymer = exploded
	}
	return len(explode(polymer))
}

func firstStar(polymer string) int {
	return getMinExplodedString(polymer)
}

func secondStar(polymer string) int {
	uniqueStrings := make(map[string]bool)
	for i := 0; i < len(polymer); i++ {
		s := strings.ToLower(string(polymer[i]))
		uniqueStrings[s] = true
	}
	min := int(^uint(0) >> 1)
	for k := range uniqueStrings {
		str := strings.Replace(polymer, k, "", -1)
		str = strings.Replace(str, strings.ToUpper(k), "", -1)
		explodedString := getMinExplodedString(str)
		if explodedString < min {
			min = explodedString
		}
	}
	return min
}

func main() {
	polymer := readPolymer("../input.csv")
	firstStar := firstStar(polymer)
	fmt.Println(firstStar)
	secondStar := secondStar(polymer)
	fmt.Println(secondStar)
}

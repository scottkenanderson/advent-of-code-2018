package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) string {
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
	return strings.Join(lines, "")
}

type header struct {
	numChildNodes, numMetadataEntries int
}

type node struct {
	header     header
	childNodes []node
	metadata   []int
}

func getNode(input []string, i int) (node, int) {
	numChildNodes, _ := strconv.Atoi(input[i])
	i++
	numMetadataEntries, _ := strconv.Atoi(input[i])
	i++
	header := header{numChildNodes, numMetadataEntries}
	n := node{header: header}

	// fmt.Println(header)
	for c := 0; c < numChildNodes; c++ {
		var child node
		child, i = getNode(input, i)
		n.childNodes = append(n.childNodes, child)
	}
	for m := 0; m < numMetadataEntries; m++ {
		metadata, _ := strconv.Atoi(input[i])
		i++
		n.metadata = append(n.metadata, metadata)
	}
	return n, i
}

func sumMetadata(n node) int {
	sum := 0
	for i := 0; i < n.header.numMetadataEntries; i++ {
		sum += n.metadata[i]
	}
	for i := 0; i < n.header.numChildNodes; i++ {
		sum += sumMetadata(n.childNodes[i])
	}
	return sum
}

func metadataValue(n node) int {
	sum := 0
	if n.header.numChildNodes == 0 {
		return sumMetadata(n)
	}
	for i := 0; i < n.header.numMetadataEntries; i++ {
		m := n.metadata[i] - 1
		if m >= n.header.numChildNodes {
			continue
		}
		sum += metadataValue(n.childNodes[m])
	}
	return sum
}

func firstStar(inputString string) int {
	input := strings.Split(inputString, " ")
	n, _ := getNode(input, 0)
	sum := sumMetadata(n)
	return sum
}

func secondStar(inputString string) int {
	input := strings.Split(inputString, " ")
	n, _ := getNode(input, 0)
	v := metadataValue(n)
	return v
}

func main() {
	schedule := readInput("../input.csv")
	firstStar := firstStar(schedule)
	fmt.Println(firstStar)
	secondStar := secondStar(schedule)
	fmt.Println(secondStar)
}

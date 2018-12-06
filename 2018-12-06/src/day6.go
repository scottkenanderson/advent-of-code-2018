package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readCoOrdinates(filename string) coOrdinates {
	csvFile, _ := os.Open(filename)
	reader := bufio.NewScanner(csvFile)
	var lines []string
	var coOrdinates coOrdinates
	for reader.Scan() {
		line := reader.Text()
		lines = append(lines, line)
		xy := strings.Split(line, ", ")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		coOrdinate := coOrdinate{x: x, y: y}
		coOrdinates.coOrdinates = append(coOrdinates.coOrdinates, coOrdinate)
	}
	if err := reader.Err(); err != nil {
		log.Fatal(err)
	}
	return coOrdinates
}

type coOrdinate struct {
	x, y int
}

type coOrdinates struct {
	coOrdinates []coOrdinate
}

func (c coOrdinates) minXAndY() (int, int) {
	minX := int(^uint(0) >> 1)
	minY := int(^uint(0) >> 1)
	coOrdinatesSlice := c.getCoOrdinates()
	for i := range coOrdinatesSlice {
		if minX > coOrdinatesSlice[i].x {
			minX = coOrdinatesSlice[i].x
		}
		if minY > coOrdinatesSlice[i].y {
			minY = coOrdinatesSlice[i].y
		}
	}
	return minX, minY
}

func (c coOrdinates) findNearestPoint(x, y int) int {
	min := int(^uint(0) >> 1)
	minIndex := -1
	coOrdinatesSlice := c.getCoOrdinates()
	dupe := false
	for i := range coOrdinatesSlice {
		distance := int(math.Abs(float64(x-coOrdinatesSlice[i].x)) + math.Abs(float64(y-coOrdinatesSlice[i].y)))
		if min > distance {
			min = distance
			minIndex = i
			dupe = false
		} else if min == distance {
			dupe = true
		}
	}
	if !dupe {
		return minIndex
	}
	return -1
}

func (c coOrdinates) findDistanceToAllPoints(x, y int) int {
	coOrdinatesSlice := c.getCoOrdinates()
	distance := 0
	for i := range coOrdinatesSlice {
		distance += int(math.Abs(float64(x-coOrdinatesSlice[i].x)) + math.Abs(float64(y-coOrdinatesSlice[i].y)))

	}
	return distance
}

func (c coOrdinates) getCoOrdinates() []coOrdinate {
	return c.coOrdinates
}

func (c coOrdinates) maxXAndY() (int, int) {
	minX := 0
	minY := 0
	coOrdinatesSlice := c.getCoOrdinates()
	for i := range coOrdinatesSlice {
		if minX < coOrdinatesSlice[i].x {
			minX = coOrdinatesSlice[i].x
		}
		if minY < coOrdinatesSlice[i].y {
			minY = coOrdinatesSlice[i].y
		}
	}
	return minX, minY
}

func max(numbers map[int]int) (int, int) {
	var key, maxValue int
	for k, v := range numbers {
		maxValue = v
		key = k
		break
	}
	for k, v := range numbers {
		if v > maxValue {
			maxValue = v
			key = k
		}
	}
	return key, maxValue
}

func isPerimeter(x, y, minX, minY, maxX, maxY int) bool {
	return x == minX || x == maxX || y == minY || y == maxX
}

func firstStar(coOrdinates coOrdinates) int {
	minX, minY := coOrdinates.minXAndY()
	maxX, maxY := coOrdinates.maxXAndY()
	distances := make(map[int]int)
	perimeter := make(map[int]bool)
	for y := minY - 200; y < maxY+200; y++ {
		for x := minX - 200; x < maxX+200; x++ {
			nearestPoint := coOrdinates.findNearestPoint(x, y)
			if isPerimeter(x, y, minX-1, minY-1, maxX+1, maxY+1) {
				perimeter[nearestPoint] = true
			}
			distances[nearestPoint]++
		}
	}
	for k := range perimeter {
		distances[k] = 0
	}
	_, maxValue := max(distances)
	return maxValue
}

func secondStar(coOrdinates coOrdinates, limit int) int {
	minX, minY := coOrdinates.minXAndY()
	maxX, maxY := coOrdinates.maxXAndY()
	size := 0
	for y := minY - 200; y < maxY+200; y++ {
		for x := minX - 200; x < maxX+200; x++ {
			distanceToAllPoints := coOrdinates.findDistanceToAllPoints(x, y)
			if distanceToAllPoints < limit {
				size++
			}
		}
	}
	return size
}

func main() {
	coOrdinates := readCoOrdinates("../input.csv")
	firstStar := firstStar(coOrdinates)
	fmt.Println(firstStar)
	secondStar := secondStar(coOrdinates, 10000)
	fmt.Println(secondStar)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	return strings.Join(lines, "\n")
}

type position struct {
	x, y int
}

type velocity struct {
	x, y int
}

type point struct {
	position
	velocity
}

func getPoint(pointString string) point {
	rexExp, _ := regexp.Compile(`position=< *(-?\d+), *(-?\d+)> velocity=< *(-?\d+), *(-?\d+)>`)
	match := rexExp.FindStringSubmatch(pointString)
	positionX, _ := strconv.Atoi(match[1])
	positionY, _ := strconv.Atoi(match[2])
	velocityX, _ := strconv.Atoi(match[3])
	velocityY, _ := strconv.Atoi(match[4])
	return point{
		position{x: positionX, y: positionY},
		velocity{x: velocityX, y: velocityY},
	}
}

func getPoints(input string) []point {
	inputSplit := strings.Split(input, "\n")
	points := make([]point, 0)
	for i := range inputSplit {
		point := getPoint(inputSplit[i])
		points = append(points, point)
	}
	return points
}

func (p point) move(second int) position {
	xMovement := p.velocity.x * second
	yMovement := p.velocity.y * second
	return position{
		x: p.position.x + xMovement,
		y: p.position.y + yMovement,
	}
}

func (p position) offset(offset position) position {
	return position{
		x: p.x + int(offset.x),
		y: p.y + int(offset.y),
	}
}

func getOutputList(points []point, second int) ([]position, position, position) {
	minX := 1<<31 - 1
	minY := 1<<31 - 1
	var maxX, maxY int
	positions := make([]position, 0)
	for i := range points {
		pos := points[i].move(second)
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.x < minX {
			minX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
		if pos.y < minY {
			minY = pos.y
		}
		positions = append(positions, pos)
	}

	return positions, position{x: minX, y: minY}, position{x: maxX, y: maxY}
}

func getArea(points []point, second int) int {
	_, min, max := getOutputList(points, second)
	return (max.x - min.x) * (max.y - min.y)
}

func (p position) hasNeighbours(points []position) bool {
	for i := range points {
		other := points[i]
		if other.x == p.x && other.y == p.y {
			continue
		}
		if other.x >= p.x-1 && other.x <= p.x+1 && other.y >= p.y-1 && other.y <= p.y+1 {
			return true
		}
	}
	return false
}

func allPointsHaveNeighbours(points []point, second int) bool {
	output, _, _ := getOutputList(points, second)
	for i := range output {
		if !output[i].hasNeighbours(output) {
			return false
		}
	}
	return true
}

func (p position) getOffsetToZero() position {
	return position{
		x: 0 - p.x,
		y: 0 - p.y,
	}
}

func getMessage(points []point, second int) string {
	outputList, min, max := getOutputList(points, second)
	offset := min.getOffsetToZero()
	offsetMax := max.offset(offset)
	message := make([][]bool, offsetMax.y+1)
	for y := 0; y <= offsetMax.y; y++ {
		message[y] = make([]bool, offsetMax.x+1)
	}
	for i := range outputList {
		p := outputList[i].offset(offset)
		message[p.y][p.x] = true
	}
	var sb strings.Builder
	for y := 0; y < offsetMax.y+1; y++ {
		for x := 0; x < offsetMax.x+1; x++ {
			if message[y][x] {
				sb.WriteString("#")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func firstStar(input string) string {
	points := getPoints(input)
	i := 0
	for !allPointsHaveNeighbours(points, i) {
		i++
	}
	return getMessage(points, i)
}

func secondStar(input string) int {
	points := getPoints(input)
	area := getArea(points, 0)
	newArea := area
	i := 0
	for true {
		newArea = getArea(points, i)
		if area < newArea {
			break
		}
		area = newArea
		i++
	}
	return i - 1
}

func main() {
	input := readInput("../input.csv")
	firstStar := firstStar(input)
	fmt.Println(firstStar)
	secondStar := secondStar(input)
	fmt.Println(secondStar)
}

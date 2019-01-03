package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	fmt.Println(inputSplit[0])
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
		x: p.x + int(math.Abs(float64(offset.x))),
		y: p.y + int(math.Abs(float64(offset.y))),
	}
}

func getOutputList(points []point, second int) ([]position, position, position) {
	var minX, maxX, minY, maxY int
	minX = 1<<31 - 1
	minY = 1<<31 - 1
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
	// fmt.Println(min, min.offset(min), max, max.offset(min))
	return max.offset(min).x * max.offset(min).y
}

func (p position) hasNeighbours(points []position) bool {
	for i := range points {
		other := points[i]
		if other.x == p.x && other.y == p.y {
			continue
		}
		if (other.x >= p.x-1 && other.x <= p.x+1) && (other.y >= p.y-1 && other.y <= p.y+1) {
			return true
		}
	}
	return false
}

func allPointsHaveNeighbours(points []point, second int) bool {
	output, _, _ := getOutputList(points, second)
	// fmt.Println(min, min.offset(min), max, max.offset(min))
	for i := range output {
		if !output[i].hasNeighbours(output) {
			return false
		}
	}
	return true
}

func getMessage(points []point, second int) string {
	outputList, min, max := getOutputList(points, second)
	fmt.Println(min, max)

	message := make([][]bool, max.offset(min).y+1)
	for y := 0; y <= max.offset(min).y; y++ {
		message[y] = make([]bool, max.offset(min).x+1)
	}
	fmt.Println(min, min.offset(min), max, max.offset(min))
	// fmt.Println(message)
	for i := range outputList {
		p := outputList[i].offset(min)
		message[p.y][p.x] = true
	}
	var sb strings.Builder
	for y := range message {
		for x := range message[y] {
			char := message[y][x]
			if char {
				sb.WriteString("#")
				fmt.Print("#")
			} else {
				sb.WriteString(".")
				fmt.Print(".")
			}
		}
		sb.WriteString("\n")
		fmt.Print("\n")
	}
	// fmt.Println(sb.String())
	return sb.String()
}

func firstStar(input string) string {
	fmt.Println(input)
	points := getPoints(input)
	// fmt.Println(points)
	// return getMessage(points, 3)
	// area := getArea(points, 0)
	for i := 1; i < 100000; i++ {
		// newArea := getArea(points, i)
		if allPointsHaveNeighbours(points, i) {
			fmt.Println("yes", i)
			getMessage(points, i)
			break
		}
	}
	return ""
}

func secondStar() int {
	return 0
}

func main() {
	input := readInput("../input.csv")

	// rexExp, _ := regexp.Compile(`(\d+) players; last marble is worth (\d+) points`)
	// match := rexExp.FindStringSubmatch(input)
	// players, _ := strconv.Atoi(match[1])
	// lastMarble, _ := strconv.Atoi(match[2])

	firstStar := firstStar(input)
	fmt.Println(firstStar)
	secondStar := secondStar()
	fmt.Println(secondStar)
}

package main

import (
	"fmt"
	"math"
)

type position struct {
	x, y int
}

type fuelCell struct {
	position         position
	gridSerialNumber int
}

func (f fuelCell) powerLevel() int {
	rackID := f.position.x + 10
	powerLevel := rackID * f.position.y
	powerLevel += f.gridSerialNumber
	powerLevel *= rackID
	powerLevel = (powerLevel / 100) % 10
	return powerLevel - 5
}

func findSquarePower(startPos position, gridSize, gridSerialNumber int) int {
	powerLevel := 0
	for y := startPos.y; y < startPos.y+gridSize; y++ {
		for x := startPos.x; x < startPos.x+gridSize; x++ {
			f := fuelCell{position{x, y}, gridSerialNumber}
			powerLevel += f.powerLevel()
		}
	}
	return powerLevel
}

func findPrecalculatedSquarePower(startPos position, gridSize int, grid [][]int) int {
	powerLevel := 0
	for y := startPos.y; y < startPos.y+gridSize; y++ {
		for x := startPos.x; x < startPos.x+gridSize; x++ {
			powerLevel += grid[y][x]
		}
	}
	return powerLevel
}

func firstStar(gridSerialNumber int) (position, int) {
	maxPower := math.MinInt32
	highestPower := position{}
	for y := 1; y <= 297; y++ {
		for x := 1; x <= 297; x++ {
			p := position{x, y}
			squarePower := findSquarePower(p, 3, gridSerialNumber)
			if squarePower > maxPower {
				maxPower = squarePower
				highestPower = p
			}
		}
	}
	return highestPower, maxPower
}

func secondStar(gridSerialNumber, maxGridSize int) string {
	maxPower := math.MinInt32
	var highestkey string
	powers := make([][]int, maxGridSize+1)
	for y := 1; y <= maxGridSize; y++ {
		powers[y] = make([]int, maxGridSize+1)
		for x := 1; x <= maxGridSize; x++ {
			f := fuelCell{position{x, y}, gridSerialNumber}
			powers[y][x] = f.powerLevel()
		}
	}

	for squareSize := 1; squareSize <= maxGridSize; squareSize++ {
		for y := 1; y <= maxGridSize-squareSize+1; y++ {
			for x := 1; x <= maxGridSize-squareSize+1; x++ {
				p := position{x, y}
				key := fmt.Sprintf("%v,%v,%v", p.x, p.y, squareSize)
				power := findPrecalculatedSquarePower(p, squareSize, powers)
				if power > maxPower {
					highestkey = key
					maxPower = power
				}
			}
		}
	}
	return highestkey
}

func main() {
	input := 6303
	firstStar, powerLevel := firstStar(input)
	fmt.Println(firstStar, powerLevel)
	secondStar := secondStar(input, 300)
	fmt.Println(secondStar)
}

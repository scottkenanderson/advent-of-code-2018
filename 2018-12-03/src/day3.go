package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	claimID            string
	x, y, sizeX, sizeY int
}

func getClaims(filename string) []claim {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ' '
	var claims []claim
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		claimID := line[0]
		xy := line[2][:len(line[2])-1]
		xyRaw := strings.Split(xy, ",")
		x, _ := strconv.Atoi(xyRaw[0])
		y, _ := strconv.Atoi(xyRaw[1])
		sizes := line[3]
		rawSizes := strings.Split(sizes, "x")
		sizeX, _ := strconv.Atoi(rawSizes[0])
		sizeY, _ := strconv.Atoi(rawSizes[1])
		claims = append(claims, claim{
			claimID: claimID,
			x:       x,
			y:       y,
			sizeX:   sizeX,
			sizeY:   sizeY,
		})
	}
	return claims
}

func findOverlaps(claims []claim, sizeX, sizeY int) int {
	squareInches := [1000][1000]int{}
	overlaps := 0
	for i := 0; i < len(claims); i++ {
		c := claims[i]
		startX := c.x
		startY := c.y
		fmt.Println(c)

		for y := startY; y < startY+c.sizeY; y++ {
			for x := startX; x < startX+c.sizeX; x++ {
				squareInches[y][x]++
				if squareInches[y][x] == 2 {
					overlaps++
				}
			}
		}
	}
	return overlaps
}

func firstStar(claims []claim) int {
	return findOverlaps(claims, 1000, 1000)
}

func main() {
	claims := getClaims("../input.csv")
	firstStar := firstStar(claims)
	fmt.Println(firstStar)
}

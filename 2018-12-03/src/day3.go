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

type squareInch struct {
	occupied bool
	claimIDs []string
}

func getClaims(filename string) ([]claim, map[string]bool) {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ' '
	var claims []claim
	var allClaimIds = make(map[string]bool)
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		claimID := line[0][1:len(line[0])]
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
		allClaimIds[claimID] = true
	}
	return claims, allClaimIds
}

func findOverlaps(claims []claim, sizeX, sizeY int) int {
	squareInches := [1000][1000]int{}
	overlaps := 0
	for i := 0; i < len(claims); i++ {
		c := claims[i]
		for y := c.y; y < c.y+c.sizeY; y++ {
			for x := c.x; x < c.x+c.sizeX; x++ {
				squareInches[y][x]++
				if squareInches[y][x] == 2 {
					overlaps++
				}
			}
		}
	}
	return overlaps
}

func findUniqueClaim(allClaimIds map[string]bool, overlappingClaims map[string]bool) string {
	for k, _ := range allClaimIds {
		if !overlappingClaims[k] {
			return k
		}
	}
	return "";
}

func findNonOverlappingClaim(claims []claim, allClaimIds map[string]bool) string {
	squareInches := [1000][1000]squareInch{}
	overlaps := 0
	overlappingClaims := make(map[string]bool)
	for i := 0; i < len(claims); i++ {
		c := claims[i]
		for y := c.y; y < c.y+c.sizeY; y++ {
			for x := c.x; x < c.x+c.sizeX; x++ {
				squareInches[y][x].occupied = true
				if len(squareInches[y][x].claimIDs) == 1 {
					overlaps++
					overlappingClaims[squareInches[y][x].claimIDs[0]] = true
				}
				if len(squareInches[y][x].claimIDs) > 0 {
					overlappingClaims[c.claimID] = true
				}
				squareInches[y][x].claimIDs = append(squareInches[y][x].claimIDs, c.claimID)
			}
		}
	}
	diff := findUniqueClaim(allClaimIds, overlappingClaims)
	return diff
}

func firstStar(claims []claim) int {
	return findOverlaps(claims, 1000, 1000)
}

func secondStar(claims []claim, allClaimIds map[string]bool) string {
	return findNonOverlappingClaim(claims, allClaimIds)
}

func main() {
	claims, allClaimIds := getClaims("../input.csv")
	firstStar := firstStar(claims)
	fmt.Println("first star: ", firstStar)
	secondStar := secondStar(claims, allClaimIds)
	fmt.Println("second star: ", secondStar)
}

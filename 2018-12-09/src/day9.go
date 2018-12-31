package main

import (
	"bufio"
	"container/ring"
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
	return strings.Join(lines, "")
}

func max(numbers []int) int {
	maxValue := 0
	for v := range numbers {
		if numbers[v] > maxValue {
			maxValue = numbers[v]
		}
	}
	return maxValue
}

func firstStar(numPlayers, lastMarble int) int {
	players := make([]int, numPlayers)
	board := []int{0}
	currentMarbleIndex := 0
	for i := 1; i <= lastMarble; i++ {
		fmt.Println(i)
		playerNumber := (i - 1) % numPlayers
		if i%23 == 0 {
			players[playerNumber] += i
			currentMarbleIndex -= 7
			if currentMarbleIndex < 0 {
				currentMarbleIndex += len(board)
			}
			players[playerNumber] += board[currentMarbleIndex]

			// fmt.Println("player", playerNumber, "cMI", currentMarbleIndex, "i", i, "boardCMI", board[currentMarbleIndex], "players", players)
			copy(board[currentMarbleIndex:], board[currentMarbleIndex+1:])
			board = board[:len(board)-1]
		} else {
			before := (currentMarbleIndex + 1) % len(board)
			board = append(board, 0)
			copy(board[before+1:], board[before:])
			board[before+1] = i
			currentMarbleIndex = before + 1
		}
		// fmt.Println("board", board, "current", currentMarbleIndex, "i", i, "length", len(board), "current marble", board[currentMarbleIndex])
		// fmt.Println(board)
	}
	return max(players)
}

func secondStar(numPlayers, lastMarble int) int {
	players := make([]int, numPlayers)
	board := ring.New(1)
	board.Value = 0
	for i := 1; i <= lastMarble; i++ {
		if i%23 == 0 {
			playerNumber := (i - 1) % numPlayers
			board = board.Move(-7)
			players[playerNumber] += (i + board.Value.(int))
			prev := board.Prev()
			next := board.Next()
			prev.Link(next)
			board = next
		} else {
			board = board.Next()
			after := board.Next()
			r := ring.New(1)
			r.Value = i
			board.Link(r)
			r.Link(after)
			board = r
		}

	}
	return max(players)
}

func main() {
	input := readInput("../input.csv")

	rexExp, _ := regexp.Compile(`(\d+) players; last marble is worth (\d+) points`)
	match := rexExp.FindStringSubmatch(input)
	players, _ := strconv.Atoi(match[1])
	lastMarble, _ := strconv.Atoi(match[2])

	firstStar := firstStar(players, lastMarble)
	fmt.Println(firstStar)
	secondStar := secondStar(players, lastMarble*100)
	fmt.Println(secondStar)
}

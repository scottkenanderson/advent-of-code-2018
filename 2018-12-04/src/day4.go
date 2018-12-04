package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func readSchedule(filename string) []string {
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
	sort.Strings(lines)
	return lines
}

func shiftStart(record string) bool {
	return strings.Contains(record, "begins")
}

func fallsAsleep(record string) bool {
	return strings.Contains(record, "falls asleep")
}

func wakesUp(record string) bool {
	return strings.Contains(record, "wakes up")
}

func getMatch(record string, regex string) int {
	rexExp, _ := regexp.Compile(regex)
	match := rexExp.FindStringSubmatch(record)[1]
	integerValue, _ := strconv.Atoi(match)
	return integerValue
}

func getGuardID(record string) int {
	return getMatch(record, "Guard #(\\d+)")
}

func getMinute(record string) int {
	return getMatch(record, "\\[....-..-.. ..:(\\d{2})\\]")
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

type sleep struct {
	start, end int
}

type guard struct {
	sleeps []sleep
}

func getSleepiestMinute(guard *guard) (int, int) {
	minutes := make(map[int]int)
	sleeps := guard.sleeps
	for i := range sleeps {
		for m := sleeps[i].start; m < sleeps[i].end; m++ {
			minutes[m]++
		}
	}
	sleepiestMinute, times := max(minutes)
	return sleepiestMinute, times
}

func getGuardRecords(schedule []string) (map[int]*guard, map[int]int) {
	var guardID, minuteStart, minuteEnd int
	guardSleepTime := make(map[int]int)
	guardRecords := make(map[int]*guard)
	for i := range schedule {
		record := schedule[i]
		if shiftStart(record) {
			guardID = getGuardID(record)
		}
		if fallsAsleep(record) {
			minuteStart = getMinute(record)
		}
		if wakesUp(record) {
			minuteEnd = getMinute(record)
			guardSleepTime[guardID] += (minuteEnd - minuteStart)
			if guardRecords[guardID] == nil {
				guardRecords[guardID] = new(guard)
			}
			guardRecords[guardID].sleeps = append(guardRecords[guardID].sleeps,
				sleep{start: minuteStart, end: minuteEnd},
			)
		}
	}
	return guardRecords, guardSleepTime
}

func firstStar(schedule []string) int {
	guardRecords, guardSleepTime := getGuardRecords(schedule)
	sleepiestGuard, _ := max(guardSleepTime)
	sleepiestMinute, _ := getSleepiestMinute(guardRecords[sleepiestGuard])
	return sleepiestGuard * sleepiestMinute
}

func secondStar(schedule []string) int {
	guardRecords, _ := getGuardRecords(schedule)
	var sleepiestMinute, maxTimes, sleepiestGuard int
	for k, v := range guardRecords {
		m, t := getSleepiestMinute(v)
		if t > maxTimes {
			maxTimes = t
			sleepiestMinute = m
			sleepiestGuard = k
		}
	}
	return sleepiestGuard * sleepiestMinute
}

func main() {
	schedule := readSchedule("../input.csv")
	firstStar := firstStar(schedule)
	fmt.Println(firstStar)
	secondStar := secondStar(schedule)
	fmt.Println(secondStar)
}

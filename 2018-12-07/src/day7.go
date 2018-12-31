package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func readInput(filename string) []string {
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
	return lines
}

func getMatch(input string, regex string) (string, string) {
	rexExp, _ := regexp.Compile(regex)
	match := rexExp.FindStringSubmatch(input)
	return match[1], match[2]
}

func getDependencies(dependencyRecord string) (string, string) {
	return getMatch(dependencyRecord, `Step ([A-Z]) must be finished before step ([A-Z]) can begin.`)
}

type dependency struct {
	id                    string
	dependsOn, dependedBy map[string]bool
}

func (d *dependency) addDependency(other string) bool {
	if d.dependsOn == nil {
		d.dependsOn = make(map[string]bool)
	}
	d.dependsOn[other] = true
	return true
}

func (d dependency) removeDependency(other string) bool {
	if d.dependsOn == nil {
		d.dependsOn = make(map[string]bool)
	}
	delete(d.dependsOn, other)
	return true
}

func (d *dependency) addDependedBy(other string) bool {
	if d.dependedBy == nil {
		d.dependedBy = make(map[string]bool)
	}
	d.dependedBy[other] = true
	return true
}

func (d dependency) hasDependencies() bool {
	return len(d.dependsOn) != 0
}

func getTree(input []string) map[string]dependency {
	dependencyMap := make(map[string]dependency)
	for i := range input {
		dependencyRecord := input[i]
		dependsOn, dependedBy := getDependencies(dependencyRecord)
		depOn := dependencyMap[dependsOn]
		depBy := dependencyMap[dependedBy]
		depBy.addDependency(dependsOn)
		depOn.addDependedBy(dependedBy)
		dependencyMap[dependedBy] = depBy
		dependencyMap[dependsOn] = depOn
	}
	return dependencyMap
}

func getNoDependencies(tree map[string]dependency) []string {
	var noDeps []string
	for i := range tree {
		if !tree[i].hasDependencies() {
			noDeps = append(noDeps, i)
		}
	}
	if len(noDeps) > 1 {
		sort.Strings(noDeps)
	}
	return noDeps
}

func getNextDependency(tree map[string]dependency) string {
	noDeps := getNoDependencies(tree)
	if len(noDeps) < 1 {
		return "."
	}
	next := noDeps[0]
	return next
}

func getTimeOfStep(step string, offset int) int {
	const abc = ".ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	return strings.Index(abc, step) + offset
}

func firstStar(input []string) string {
	tree := getTree(input)
	var next string
	var output []string
	for len(tree) > 0 {
		next = getNextDependency(tree)
		for d := range tree[next].dependedBy {
			tree[d].removeDependency(next)
		}
		output = append(output, next)
		delete(tree, next)
	}
	return strings.Join(output, "")
}

type worker struct {
	taskID        string
	timeRemaining int
	working       bool
}

func (w worker) isWorking() bool {
	return w.working
}

func (w *worker) decrementTimeRemaining() bool {
	if !w.isWorking() {
		return false
	}
	w.timeRemaining--
	if w.timeRemaining == 0 {
		w.working = true
		return true
	}
	return false
}

func firstIdleWorker(workers []worker) (int, bool) {
	for i := range workers {
		w := workers[i]
		if !w.isWorking() {
			return i, true
		}
	}
	return -1, false
}

func (d dependency) canStart(workers []worker, taskID string, tree map[string]dependency) bool {
	for w := range workers {
		if workers[w].isWorking() {
			if taskID == workers[w].taskID {
				return false
			}
			for i := range d.dependsOn {
				_, ok := tree[i]
				if ok || i == workers[w].taskID {
					return false
				}
			}
		}
	}
	return true
}

func secondStar(schedule []string, offset, numElves int) int {
	tree := getTree(schedule)
	order := strings.Split(firstStar(schedule), "")
	var workers []worker

	for i := 0; i <= numElves; i++ {
		workers = append(workers, worker{})
	}

	second := 0
	for true {
		for w := range workers {
			if workers[w].isWorking() {
				isFinished := workers[w].decrementTimeRemaining()
				if isFinished {
					workers[w].working = false
					taskID := workers[w].taskID
					delete(tree, taskID)
					workers[w].taskID = ""
				}
			}
		}
		for i := range order {
			taskID := order[i]
			task, ok := tree[order[i]]
			if !ok {
				continue
			}
			if taskID == "" {
				break
			}
			if task.canStart(workers, taskID, tree) {
				firstIdleWorker, anyIdle := firstIdleWorker(workers)
				if !anyIdle {
					break
				} else {
					workers[firstIdleWorker] = worker{
						taskID:        taskID,
						timeRemaining: getTimeOfStep(taskID, offset),
						working:       true,
					}
				}
			} else {
				continue
			}
		}
		if len(tree) > 0 {
			second++
		} else {
			break
		}
	}
	return second
}

func main() {
	schedule := readInput("../input.csv")
	firstStar := firstStar(schedule)
	fmt.Println(firstStar)
	secondStar := secondStar(schedule, 60, 4)
	fmt.Println(secondStar)
}

package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(inputFile string) []int {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), ",")
	values := make([]int, len(inputStrings))
	for i := range values {
		values[i], _ = strconv.Atoi(inputStrings[i])
	}
	return values
}

const (
	term       = 7
	penalty    = 2
	firstCycle = term + penalty
)

var cache = make(map[int]map[int]int)

func simulateFishCount(age int, daysLeft int) int {
	if val, ok := cache[age]; ok {
		if entry, valOk := val[daysLeft]; valOk {
			return entry
		}
	}
	if daysLeft <= age {
		return 0
	}
	value := 1 + simulateFishCount(firstCycle, daysLeft-age) + simulateFishCount(term, daysLeft-age)
	if _, ok := cache[age]; !ok {
		cache[age] = make(map[int]int)
	}
	cache[age][daysLeft] = value
	return value
}

func solveFirst(values []int, days int) int {
	fishCount := len(values)
	for i := range values {
		fishCount += simulateFishCount(values[i], days)
	}
	return fishCount
}

func solveSecond(values []int) int {
	days := 256
	return solveFirst(values, days)
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values, 80))
	println("Part 2: the answer is", solveSecond(values))
}

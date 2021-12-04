package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(inputFile string) []int {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	values := make([]int, len(inputStrings))
	for i := range values {
		values[i], _ = strconv.Atoi(inputStrings[i])
	}
	return values
}

func solveFirst(values []int) int {
	var counter int
	for i := len(values) - 1; i != 0; i-- {
		if (values[i] - values[i-1]) > 0 {
			counter++
		}
	}
	return counter
}

func solveSecond(values []int) int {
	var sumsOfThree = make([]int, len(values)-2)
	for i := 0; i != len(values)-2; i++ {
		sumsOfThree[i] = values[i] + values[i+1] + values[i+2]
	}
	return solveFirst(sumsOfThree)
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: Depth increased", solveFirst(values), "of times.")
	println("Part 2: Depth increased", solveSecond(values), "of times.")
}

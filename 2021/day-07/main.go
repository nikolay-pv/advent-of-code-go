package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func readInput(inputFile string) []int {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), ",")
	values := make([]int, len(inputStrings))
	for i := range values {
		values[i], _ = strconv.Atoi(inputStrings[i])
	}
	return values
}

func solveFirst(values []int) int {
	// minimize sum_i(abs(value[i] - destination)) <=> destination = median
	sort.Ints(values)
	n := len(values)
	destination := values[n/2]
	if n%2 != 1 {
		destination = (values[n/2-1] + values[n/2]) / 2
	}
	fuel := 0
	for i := range values {
		fuel += Abs(values[i] - destination)
		// println("Move from", values[i], "to", destination, ":", Abs(values[i]-destination), "fuel")
	}
	return fuel
}

func solveSecond(values []int) int {
	return values[0]
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
}

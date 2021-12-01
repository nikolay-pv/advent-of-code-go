package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func read_input(input_file string) []int {
	input, _ := ioutil.ReadFile(input_file)
	input_strings := strings.Split(string(input), "\n")
	values := make([]int, len(input_strings))
	for i := range values {
		values[i], _ = strconv.Atoi(input_strings[i])
	}
	return values
}

func solve_first(values []int) int {
	var counter int
	for i := len(values) - 1; i != 0; i-- {
		if (values[i] - values[i-1]) > 0 {
			counter++
		}
	}
	return counter
}

func solve_second(values []int) int {
	var sums_of_three = make([]int, len(values)-2)
	for i := 0; i != len(values)-2; i++ {
		sums_of_three[i] = values[i] + values[i+1] + values[i+2]
	}
	return solve_first(sums_of_three)
}

func main() {
	values := read_input("./input.txt")
	println("Part 1: Depth increased", solve_first(values), "of times.")
	println("Part 2: Depth increased", solve_second(values), "of times.")
}

package main

import (
	"io/ioutil"
	"strings"
)

func read_input(input_file string) []string {
	input, _ := ioutil.ReadFile(input_file)
	input_strings := strings.Split(string(input), "\n")
	return input_strings
}

func solve_first(values []string) int {
	width := len(values[0])
	counters := make([]int, width)
	for i := range values {
		for j := 0; j != width; j++ {
			if values[i][j] == '1' {
				counters[j] += 1
			}
		}
	}
	gamma_rate := 0
	half_of_report := len(values) / 2
	for j := 0; j != width; j++ {
		if counters[j] > half_of_report {
			gamma_rate += 1 << (width - j - 1)
		}
	}
	mask := ^(^0 << width) // first n (=width) of ones
	epsilon_rate := ^gamma_rate & mask
	return gamma_rate * epsilon_rate
}

func solve_second(input []string) int {
	return len(input)
}

func main() {
	values := read_input("./input.txt")
	println("Part 1: the answer is", solve_first(values))
	// println("Part 2: the answer is", solve_second(values))
}

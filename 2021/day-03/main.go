package main

import (
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Input struct {
	values []int
	width  int
}

func read_input(input_file string) Input {
	input, _ := ioutil.ReadFile(input_file)
	input_strings := strings.Split(string(input), "\n")
	values := make([]int, len(input_strings))
	for i := range values {
		tmp, _ := strconv.ParseInt(input_strings[i], 2, 0)
		values[i] = int(tmp)
	}
	return Input{values, len(input_strings[0])}
}

func countOnes(values []int, bitIndex int) int {
	mask := 1 << bitIndex
	counter := 0
	for i := range values {
		counter += (values[i] & mask) >> bitIndex
	}
	return counter
}

func solve_first(values []int, width int) int {
	half_of_report := len(values) / 2
	gamma_rate := 0
	for j := 0; j != width; j++ { // bit iterator
		if countOnes(values, j) > half_of_report {
			gamma_rate += 1 << j
		}
	}
	mask := ^(^0 << width) // first n (=width) of ones
	epsilon_rate := ^gamma_rate & mask
	return gamma_rate * epsilon_rate
}

func solveSecond2(input []int, width int, bitcriteria func(highBitCount, inLength int) bool) int {
	for bitIndex := width - 1; bitIndex != -1 && len(input) != 1; bitIndex-- {
		count := countOnes(input, bitIndex)
		mask := 1 << bitIndex
		sortOrder := bitcriteria(count, len(input))
		if !sortOrder { // we aim for zeros not ones then
			count = len(input) - count
		}
		sort.Slice(input, func(first, second int) bool {
			return ((input[first] & mask) > (input[second] & mask)) == sortOrder
		})
		input = input[:count]
	}
	return input[0]
}

func solve_second(input []int, width int) int {
	first_number := solveSecond2(input, width, func(highBitCount, inLength int) bool {
		return highBitCount >= int(math.Ceil(float64(inLength)/2.0))
	})
	second_number := solveSecond2(input, width, func(highBitCount, inLength int) bool {
		return highBitCount < int(math.Ceil(float64(inLength)/2.0))
	})
	return first_number * second_number
}

func main() {
	input := read_input("./input.txt")
	println("Part 1: the answer is", solve_first(input.values, input.width))
	println("Part 2: the answer is", solve_second(input.values, input.width))
}

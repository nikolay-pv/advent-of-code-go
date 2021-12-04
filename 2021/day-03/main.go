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

func readInput(inputFile string) Input {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	values := make([]int, len(inputStrings))
	for i := range values {
		tmp, _ := strconv.ParseInt(inputStrings[i], 2, 0)
		values[i] = int(tmp)
	}
	return Input{values, len(inputStrings[0])}
}

func countOnes(values []int, bitIndex int) int {
	mask := 1 << bitIndex
	counter := 0
	for i := range values {
		counter += (values[i] & mask) >> bitIndex
	}
	return counter
}

func solveFirst(values []int, width int) int {
	halfOfLength := len(values) / 2
	gammaRate := 0
	for j := 0; j != width; j++ { // bit iterator
		if countOnes(values, j) > halfOfLength {
			gammaRate += 1 << j
		}
	}
	mask := ^(^0 << width) // first n (=width) of ones
	epsilonRate := ^gammaRate & mask
	return gammaRate * epsilonRate
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

func solveSecond(input []int, width int) int {
	first := solveSecond2(input, width, func(highBitCount, inLength int) bool {
		return highBitCount >= int(math.Ceil(float64(inLength)/2.0))
	})
	second := solveSecond2(input, width, func(highBitCount, inLength int) bool {
		return highBitCount < int(math.Ceil(float64(inLength)/2.0))
	})
	return first * second
}

func main() {
	input := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(input.values, input.width))
	println("Part 2: the answer is", solveSecond(input.values, input.width))
}

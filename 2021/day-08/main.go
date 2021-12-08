package main

import (
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func readInput(inputFile string) []string {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	return inputStrings
}

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func solveFirst(entries []string) int {
	sum := 0
	for i := range entries {
		line := strings.Split(entries[i], "|")[1]
		numbers := strings.Split(line, " ")
		for j := range numbers {
			switch len(numbers[j]) {
			case 2, 4, 3, 7:
				sum += 1
			}
		}
	}
	return sum
}

func decodeNumber(s string) int {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	s = string(r)
	switch s {
	case "abcefg":
		return 0
	case "cf":
		return 1
	case "acdeg":
		return 2
	case "acdfg":
		return 3
	case "bcdf":
		return 4
	case "abdfg":
		return 5
	case "abdefg":
		return 6
	case "acf":
		return 7
	case "abcdefg":
		return 8
	case "abcdfg":
		return 9
	}
	return -1
}

func findUniqueDigits(left string) map[int]string {
	// numbers with unique representation
	result := make(map[int]string)
	numbers := strings.Split(left, " ")
	for i := range numbers {
		switch len(numbers[i]) {
		case 2:
			result[1] = numbers[i]
		case 4:
			result[4] = numbers[i]
		case 3:
			result[7] = numbers[i]
		case 7:
			result[8] = numbers[i]
		}
	}
	return result
}

func predictMapping(left string) map[rune]rune {
	runeCounts := make(map[rune]int)
	for _, rune := range left {
		if rune != ' ' {
			runeCounts[rune]++
		}
	}
	mapWrongToRight := make(map[rune]rune)
	for k, v := range runeCounts {
		switch v {
		case 4:
			mapWrongToRight[k] = 'e'
		case 6:
			mapWrongToRight[k] = 'b'
		case 9:
			mapWrongToRight[k] = 'f'
		}
	}
	unique := findUniqueDigits(left)
	guessUnmatchedRune := func(number string, trueRune rune) {
		for _, r := range number {
			if _, ok := mapWrongToRight[r]; !ok {
				mapWrongToRight[r] = trueRune
				break
			}
		}
	}
	guessUnmatchedRune(unique[1], 'c')
	guessUnmatchedRune(unique[4], 'd')
	guessUnmatchedRune(unique[7], 'a')
	guessUnmatchedRune(unique[8], 'g')
	return mapWrongToRight
}

func guessNumber(left string, right string) int {
	mapWrongToRight := predictMapping(left)
	right = strings.Map(func(r rune) rune {
		if r == ' ' {
			return r
		}
		return mapWrongToRight[r]
	}, right)
	digits := strings.Split(right, " ")
	result := 0
	for i := range digits {
		result += decodeNumber(digits[i]) * int(math.Pow(10, float64(len(digits)-1-i)))
	}
	return result
}

func solveSecond(entries []string) int {
	sum := 0
	for i := range entries {
		parts := strings.Split(entries[i], " | ")
		sum += guessNumber(parts[0], parts[1])
	}
	return sum
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
}

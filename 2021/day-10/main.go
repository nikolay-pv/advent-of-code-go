package main

import (
	"io/ioutil"
	"sort"
	"strings"
)

func readInput(inputFile string) []string {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	return inputStrings
}

func getMatching(r rune) rune {
	switch r {
	case '(':
		return ')'
	case ')':
		return '('
	case '[':
		return ']'
	case ']':
		return '['
	case '{':
		return '}'
	case '}':
		return '{'
	case '<':
		return '>'
	case '>':
		return '<'
	}
	return ' '
}

func verifyLine(line string) (rune, bool) {
	// checks if balanced otherwise returns first unbalanced symbol
	stack := make([]rune, 0)
	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			stack = append(stack, r)
		case ')', ']', '}', '>':
			if len(stack) == 0 || stack[len(stack)-1] != getMatching(r) {
				return r, false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return ' ', true
}

func solveFirst(lines []string) int {
	errorScores := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	score := 0
	for _, line := range lines {
		if r, ok := verifyLine(line); !ok {
			score += errorScores[r]
		}
	}
	return score
}

func getMissedClosing(line string) []rune {
	stack := make([]rune, 0)
	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			stack = append(stack, r)
		case ')', ']', '}', '>':
			if len(stack) == 0 || stack[len(stack)-1] != getMatching(r) {
				return []rune{}
			}
			stack = stack[:len(stack)-1]
		}
	}
	for i := range stack {
		stack[i] = getMatching(stack[i])
	}
	return stack
}

func solveSecond(lines []string) int {
	errorScores := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	scores := make([]int, 0, len(lines))
	for _, line := range lines {
		missed := getMissedClosing(line)
		if len(missed) == 0 {
			continue
		}
		score := 0
		for i := range missed {
			r := missed[len(missed)-i-1]
			score *= 5
			score += errorScores[r]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
}

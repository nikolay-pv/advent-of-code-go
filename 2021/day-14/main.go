package main

import (
	"io/ioutil"
	"math"
	"strings"
)

type Input struct {
	formula string
	rules   map[string][2]string
}

func makeRule(in string) (string, [2]string) {
	parts := strings.Split(in, " -> ")
	var res [2]string
	res[0] = parts[0][:1] + parts[1]
	res[1] = parts[1] + parts[0][1:]
	return parts[0], res
}

func splitToPairs(formula string) map[string]int64 {
	res := make(map[string]int64)
	for i := 0; i != len(formula)-1; i++ {
		res[formula[i:i+2]] += 1 // count of appearance
	}
	return res
}

func readInput(inputFile string) Input {
	input, _ := ioutil.ReadFile(inputFile)
	formulaAndRules := strings.Split(string(input), "\n;\n")
	formula := formulaAndRules[0]
	inputStrings := strings.Split(formulaAndRules[1], "\n")
	rules := make(map[string][2]string)
	for _, line := range inputStrings {
		k, v := makeRule(line)
		rules[k] = v
	}
	return Input{formula, rules}
}

func simulate(polymer map[string]int64, rules map[string][2]string, steps int64) map[string]int64 {
	// every pair creates 2 new ones
	for ; steps != 0; steps-- {
		newPolymer := make(map[string]int64)
		for element, count := range polymer {
			outcomes := rules[element]
			newPolymer[outcomes[0]] += count
			newPolymer[outcomes[1]] += count
		}
		polymer = newPolymer
	}
	return polymer
}

func solverHelper(formula string, rules map[string][2]string, steps int64) int64 {
	polymer := splitToPairs(formula)
	polymer = simulate(polymer, rules, steps)
	runesCount := make(map[string]int64)
	for pair, count := range polymer {
		runesCount[pair[:1]] += count // count only first elements
	}
	// last element occurs once more:
	runesCount[formula[len(formula)-1:]]++
	min, max := MinMax(runesCount)
	return max - min
}

func solveFirst(formula string, rules map[string][2]string) int64 {
	return solverHelper(formula, rules, 10)
}

func solveSecond(formula string, rules map[string][2]string) int64 {
	return solverHelper(formula, rules, 40)
}

func main() {
	input := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(input.formula, input.rules))
	input = readInput("./input.txt")
	println("Part 2: the answer is", solveSecond(input.formula, input.rules))
}

// utils
func Max(lhs, rhs int64) int64 {
	if rhs >= lhs {
		return rhs
	}
	return lhs
}

func Min(lhs, rhs int64) int64 {
	if lhs <= rhs {
		return lhs
	}
	return rhs
}

func MinMax(runesCount map[string]int64) (int64, int64) {
	var min int64
	min = math.MaxInt64
	var max int64
	max = math.MinInt64
	for _, v := range runesCount {
		min = Min(min, v)
		max = Max(max, v)
	}
	return min, max
}

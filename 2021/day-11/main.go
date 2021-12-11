package main

import (
	"io/ioutil"
	"strings"
)

func readInput(inputFile string) [][]int {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	values := make([][]int, len(inputStrings))
	for i, line := range inputStrings {
		values[i] = make([]int, len(line))
		for j, r := range line {
			values[i][j] = int(r - '0')
		}
	}
	return values
}

func Max(lhs, rhs int) int {
	if rhs >= lhs {
		return rhs
	}
	return lhs
}

func Min(lhs, rhs int) int {
	if lhs <= rhs {
		return lhs
	}
	return rhs
}

type Pair struct {
	first, second int
}

func step(values [][]int) int {
	queue := make([]Pair, 0)
	flashed := make(map[Pair]struct{}, 0)
	checkValue := func(i, j int) {
		if values[i][j] >= 10 {
			p := Pair{i, j}
			if _, ok := flashed[p]; !ok {
				flashed[p] = struct{}{}
				queue = append(queue, p)
			}
		}
	}
	// increase energy of all by 1
	for i := range values {
		for j := range values[i] {
			values[i][j]++
			checkValue(i, j)
		}
	}
	// update neighbors
	rows := len(values)
	cols := len(values[0])
	for len(queue) != 0 {
		cell := queue[0]
		queue = queue[1:]
		for i := Max(cell.first-1, 0); i != Min(cell.first+2, rows); i++ {
			for j := Max(cell.second-1, 0); j != Min(cell.second+2, cols); j++ {
				if i != cell.first || j != cell.second {
					values[i][j]++
				}
				checkValue(i, j)
			}
		}
	}
	for k := range flashed {
		values[k.first][k.second] = 0
	}
	return len(flashed)
}

func solveFirst(values [][]int, steps int) int {
	flashed := 0
	for i := 0; i != steps; i++ {
		flashed += step(values)
	}
	return flashed
}

func isAllLit(values [][]int) bool {
	for i := range values {
		for _, v := range values[i] {
			if v != 0 {
				return false
			}
		}
	}
	return true
}

func solveSecond(values [][]int) int {
	steps := 0
	for !isAllLit(values) {
		steps++
		step(values)
	}
	return steps
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values, 100))
	values = readInput("./input.txt") // clear input or could add 100 to the result
	println("Part 2: the answer is", solveSecond(values))
}

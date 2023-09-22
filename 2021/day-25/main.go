package main

import (
	"io/ioutil"
	"strings"
)

func readInput(inputFile string) [][]int {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	values := make([][]int, len(inputStrings))
	for i := range inputStrings {
		values[i] = make([]int, len(inputStrings[i]))
		for j, v := range inputStrings[i] {
			switch v {
			case '.':
				values[i][j] = 0
			case '>':
				values[i][j] = 1
			case 'v':
				values[i][j] = 2
			}
		}
	}
	return values
}

func printSeaFloor(values [][]int) {
	rows := len(values)
	cols := len(values[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			v := values[row][col]
			switch v {
			case 0:
				print(".")
			case 1:
				print(">")
			case 2:
				print("v")
			}
		}
		println()
	}
	println()
}

func moveOnce(values [][]int) int {
	rows := len(values)
	cols := len(values[0])
	movedCount := 0
	// left to right, move all 1 to 0
	toBeMoved := make([]int, 0)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if values[row][col] == 1 && values[row][(col+1)%cols] == 0 {
				toBeMoved = append(toBeMoved, col)
			}
		}
		movedCount += len(toBeMoved)
		for _, col := range toBeMoved {
			values[row][col] = 0
			values[row][(col+1)%cols] = 1
		}
		toBeMoved = nil
	}
	// top to bottom, move all 2 to 0
	toBeMoved = nil
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if values[row][col] == 2 && values[(row+1)%rows][col] == 0 {
				toBeMoved = append(toBeMoved, row)
			}
		}
		movedCount += len(toBeMoved)
		for _, row := range toBeMoved {
			values[row][col] = 0
			values[(row+1)%rows][col] = 2
		}
		toBeMoved = nil
	}
	return movedCount
}

func solveFirst(values [][]int) int {
	stepCount := 0
	movedCount := 1
	// printSeaFloor(values)
	for movedCount != 0 {
		movedCount = moveOnce(values)
		stepCount += 1
		// println("Iteration", stepCount)
		// printSeaFloor(values)
	}
	return stepCount
}

func solveSecond(values [][]int) int {
	return values[0][0]
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	// println("Part 2: the answer is", solveSecond(values))
}

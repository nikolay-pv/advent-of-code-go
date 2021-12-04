package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type BingoBoard struct {
	board  []int
	marked []bool
}

func (b BingoBoard) isWinner() bool {
	for offset := 0; offset != 5; offset++ {
		isRowMarked := true
		isColMarked := true
		for increment := 0; increment != 5; increment++ {
			isRowMarked = isRowMarked && b.marked[5*offset+increment]
			isColMarked = isColMarked && b.marked[offset+5*increment]
		}
		if isRowMarked || isColMarked {
			return true
		}
	}
	return false
}

func (b BingoBoard) mark(drawn int) {
	for i := range b.board {
		if b.board[i] == drawn {
			b.marked[i] = true
		}
	}
}

func (b BingoBoard) countUnmarked() int {
	count := 0
	for i := range b.marked {
		if !b.marked[i] {
			count += b.board[i]
		}
	}
	return count
}

func makeBoard(input string) BingoBoard {
	board := make([]int, 25)
	marked := make([]bool, 25)
	inputNumbers := strings.Split(input, " ")
	for i := range inputNumbers {
		board[i], _ = strconv.Atoi(inputNumbers[i])
		marked[i] = false
	}
	return BingoBoard{board, marked}
}

type Input struct {
	values []int
	boards []BingoBoard
}

func readInput(inputFile string) Input {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	// first line is draw numbers
	inputValues := strings.Split(inputStrings[0], ",")
	values := make([]int, len(inputValues))
	for i := range values {
		values[i], _ = strconv.Atoi(inputValues[i])
	}
	// remaining is boards
	boards := make([]BingoBoard, len(inputStrings)-1)
	for i := 1; i != len(inputStrings); i++ {
		boards[i-1] = makeBoard(inputStrings[i])
	}
	return Input{values, boards}
}

func solveFirst(values []int, boards []BingoBoard) int {
	lastDrawn := 0
	winBoardIndex := 0
valuesLoop:
	for i := range values {
		for j := range boards {
			boards[j].mark(values[i])
			if boards[j].isWinner() {
				lastDrawn = values[i]
				winBoardIndex = j
				break valuesLoop
			}
		}
	}
	return lastDrawn * boards[winBoardIndex].countUnmarked()
}

func solveSecond(values []int, boards []BingoBoard) int {
	lastDrawn := 0
	winBoardIndex := 0
	winningBoards := make(map[int]int)
	for i := range values {
		for j := range boards {
			boards[j].mark(values[i])
			if _, wonAlready := winningBoards[j]; boards[j].isWinner() && !wonAlready {
				lastDrawn = values[i]
				winBoardIndex = j
				winningBoards[j] = boards[j].countUnmarked()
			}
		}
	}
	return lastDrawn * winningBoards[winBoardIndex]
}

func main() {
	input := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(input.values, input.boards))
	println("Part 2: the answer is", solveSecond(input.values, input.boards))
}

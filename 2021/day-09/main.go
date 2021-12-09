package main

import (
	"io/ioutil"
	"sort"
	"strings"
)

func readInput(inputFile string) [][]int {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	values := make([][]int, len(inputStrings))
	for i := range inputStrings {
		values[i] = make([]int, len(inputStrings[i]))
		for j, r := range inputStrings[i] {
			values[i][j] = int(r - '0')
		}
	}
	return values
}

type Pair struct {
	first, second int
}

func findLowestPoints(values [][]int) []Pair {
	rows := len(values)
	cols := len(values[0])
	isLowest := func(i, j int) bool {
		x := values[i][j]
		res := (i-1 < 0) || x < values[i-1][j]
		res = res && ((i+1 == rows) || x < values[i+1][j])
		res = res && ((j-1 < 0) || x < values[i][j-1])
		res = res && ((j+1 == cols) || x < values[i][j+1])
		return res
	}
	riskPoints := make([]Pair, 0)
	for i, row := range values {
		for j := range row {
			if isLowest(i, j) {
				riskPoints = append(riskPoints, Pair{i, j})
			}
		}
	}
	return riskPoints
}

func solveFirst(values [][]int) int {
	points := findLowestPoints(values)
	risks := 0
	for _, p := range points {
		risks += values[p.first][p.second] + 1
	}
	return risks
}

func solveSecond(values [][]int) int {
	rows := len(values)
	cols := len(values[0])
	getLowNeighbors := func(i, j int) []Pair {
		points := make([]Pair, 0)
		if i-1 >= 0 && values[i-1][j] != 9 {
			points = append(points, Pair{i - 1, j})
		}
		if i+1 < rows && values[i+1][j] != 9 {
			points = append(points, Pair{i + 1, j})
		}
		if j-1 >= 0 && values[i][j-1] != 9 {
			points = append(points, Pair{i, j - 1})
		}
		if j+1 < cols && values[i][j+1] != 9 {
			points = append(points, Pair{i, j + 1})
		}
		return points
	}
	points := findLowestPoints(values)
	areas := make([]int, 0)
	queue := make([]Pair, 1)
	for _, lowPoint := range points {
		queue = queue[:0]
		queue = append(queue, lowPoint)
		tiles := make(map[Pair]struct{}, 1)
		for len(queue) > 0 {
			p := queue[0]
			tiles[p] = struct{}{}
			queue = queue[1:]
			neighbors := getLowNeighbors(p.first, p.second)
			for j := range neighbors {
				if _, ok := tiles[neighbors[j]]; !ok {
					queue = append(queue, neighbors[j])
				}
			}
		}
		areas = append(areas, len(tiles))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(areas)))
	return areas[0] * areas[1] * areas[2]
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
}

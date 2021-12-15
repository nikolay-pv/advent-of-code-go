package main

import (
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

type pair struct {
	first, second int
}

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

func risk(riskMap [][]int, start, end pair) int {
	rows := len(riskMap)
	cols := len(riskMap[0])
	neighbors := func(i, j int) []pair {
		points := make([]pair, 0)
		if i+1 < rows {
			points = append(points, pair{i + 1, j})
		}
		if j+1 < cols {
			points = append(points, pair{i, j + 1})
		}
		if i-1 >= 0 {
			points = append(points, pair{i - 1, j})
		}
		if j-1 >= 0 {
			points = append(points, pair{i, j - 1})
		}
		// sort by min(manhattan distance to end)
		sort.Slice(points, func(i, j int) bool {
			distI := end.first - points[i].first + end.second - points[i].second
			distJ := end.first - points[j].first + end.second - points[j].second
			return distI < distJ
		})
		return points
	}

	score := make([][]int, rows)
	for i := range score {
		score[i] = make([]int, cols)
		for j := range score[i] {
			score[i][j] = math.MaxInt
		}
	}
	score[start.first][start.second] = 0

	queue := make([]pair, 1, rows*cols)
	queue[0] = start
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]
		for _, p := range neighbors(current.first, current.second) {
			risk := riskMap[p.first][p.second]
			tentativeScore := score[current.first][current.second] + risk
			if tentativeScore < score[p.first][p.second] {
				score[p.first][p.second] = tentativeScore
				queue = append(queue, p)
			}
		}
	}
	return score[end.first][end.second]
}

func solveFirst(values [][]int) int {
	start := pair{0, 0}
	end := pair{len(values[0]) - 1, len(values) - 1}
	return risk(values, start, end)
}

func solveSecond(template [][]int) int {
	tRows := len(template)
	tCols := len(template[0])
	values := make([][]int, 5*tRows)
	for r := range values {
		values[r] = make([]int, 5*tCols)
	}
	for r := range template {
		for c := range template[r] {
			for i := 0; i != 5; i++ {
				for j := 0; j != 5; j++ {
					values[r+i*tRows][c+j*tCols] = (template[r][c]+i+j-1)%9 + 1
				}
			}
		}
	}
	start := pair{0, 0}
	end := pair{len(values[0]) - 1, len(values) - 1}
	return risk(values, start, end)
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
}

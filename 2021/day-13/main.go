package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func Max(i1, i2 int) int {
	if i2 >= i1 {
		return i2
	}
	return i1
}

type Point struct {
	x, y int
}

type Input struct {
	points []Point
	folds  []Point
}

func makePoint(s string) Point {
	coordinates := strings.Split(s, ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	return Point{x, y}
}

func makeFold(s string) Point {
	coordinates := strings.Split(s, "=")
	x := 0
	y := 0
	switch coordinates[0][len(coordinates[0])-1] {
	case 'x':
		x, _ = strconv.Atoi(coordinates[1])
	case 'y':
		y, _ = strconv.Atoi(coordinates[1])
	}
	return Point{x, y}
}

func readInput(inputFile string) Input {
	input, _ := ioutil.ReadFile(inputFile)
	inputParts := strings.Split(string(input), "\n;\n")
	inputStrings := strings.Split(inputParts[0], "\n")
	values := make([]Point, len(inputStrings))
	for i, line := range inputStrings {
		values[i] = makePoint(line)
	}
	inputFolds := strings.Split(inputParts[1], "\n")
	folds := make([]Point, len(inputFolds))
	for i, v := range inputFolds {
		folds[i] = makeFold(v)
	}
	return Input{values, folds}
}

func fold(points []Point, folds []Point) map[Point]struct{} {
	for _, f := range folds {
		for i, p := range points {
			if f.x > 0 && p.x > f.x {
				points[i] = Point{2*f.x - p.x, p.y}
			} else if f.y > 0 && p.y > f.y {
				points[i] = Point{p.x, 2*f.y - p.y}
			}
		}
	}
	unique := make(map[Point]struct{}, 0)
	for _, p := range points {
		unique[p] = struct{}{}
	}
	return unique
}

func printField(unique map[Point]struct{}) {
	maximum := Point{0, 0}
	for p := range unique {
		maximum.x = Max(maximum.x, p.x)
		maximum.y = Max(maximum.y, p.y)
	}
	field := make([][]rune, maximum.y+1)
	for i := range field {
		field[i] = make([]rune, maximum.x+1)
		for j := range field[i] {
			field[i][j] = '.'
		}
	}
	for p := range unique {
		field[p.y][p.x] = '#'
	}
	for i := range field {
		println(string(field[i]))
	}
}

func solveFirst(points []Point, folds []Point) int {
	unique := fold(points, folds[:1])
	return len(unique)
}

func solveSecond(points []Point, folds []Point) int {
	unique := fold(points, folds)
	printField(unique)
	return len(unique)
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values.points, values.folds))
	values = readInput("./input.txt")
	println("Part 2: the answer is")
	solveSecond(values.points, values.folds)
}

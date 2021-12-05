package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// helpers for ints
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

// datastructure for easy hangling of input
type Point struct {
	x, y int
}

type Line struct {
	begin, end Point
}

func (p Point) offset(x int, y int) Point {
	p.x += x
	p.y += y
	return p
}

const (
	vertical   = iota
	horizontal = iota
	other      = iota
)

func (l Line) kind() int {
	if l.begin.x == l.end.x {
		return vertical
	} else if l.begin.y == l.end.y {
		return horizontal
	}
	return other
}

func (l Line) coveredPoints() []Point {
	// always 45 degrees or horizontal or vertical
	lineLength := Max(l.end.x-l.begin.x, l.end.y-l.begin.y) + 1
	points := make([]Point, lineLength)
	deltaX := (l.end.x - l.begin.x)
	if deltaX != 0 {
		deltaX /= Abs(deltaX)
	}
	deltaY := (l.end.y - l.begin.y)
	if deltaY != 0 {
		deltaY /= Abs(deltaY)
	}
	points[0] = l.begin
	for i := 1; i != len(points); i++ {
		points[i] = points[i-1].offset(deltaX, deltaY)
	}
	return points
}

func makePoint(input string) Point {
	coordinates := strings.Split(string(input), ",")
	point := Point{}
	point.x, _ = strconv.Atoi(coordinates[0])
	point.y, _ = strconv.Atoi(coordinates[1])
	return point
}

func makeLine(input string) Line {
	points := strings.Split(string(input), " -> ")
	line := Line{}
	line.begin = makePoint(points[0])
	line.end = makePoint(points[1])
	if line.begin.x >= line.end.x && line.begin.y >= line.end.y {
		line.end, line.begin = line.begin, line.end
	}
	return line
}

// solution
func readInput(inputFile string) []Line {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	lines := make([]Line, len(inputStrings))
	for i := range lines {
		lines[i] = makeLine(inputStrings[i])
	}
	return lines
}

func solveFirst(lines []Line) int {
	coveredPoints := make(map[Point]int)
	var lineCoverage []Point
	for i := range lines {
		switch lines[i].kind() {
		case horizontal:
			fallthrough
		case vertical:
			lineCoverage = lines[i].coveredPoints()
		case other:
			lineCoverage = lineCoverage[:0]
		}
		for j := range lineCoverage {
			coveredPoints[lineCoverage[j]] += 1
		}
	}
	counter := 0
	for k := range coveredPoints {
		if coveredPoints[k] > 1 {
			counter++
		}
	}
	return counter
}

func solveSecond(lines []Line) int {
	coveredPoints := make(map[Point]int)
	var lineCoverage []Point
	for i := range lines {
		lineCoverage = lines[i].coveredPoints()
		for j := range lineCoverage {
			coveredPoints[lineCoverage[j]] += 1
		}
	}
	counter := 0
	for k := range coveredPoints {
		if coveredPoints[k] > 1 {
			counter++
		}
	}
	return counter
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
}

package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type target struct {
	xMin, xMax, yMin, yMax int
}

func (t target) isHit(x, y int) bool {
	return x >= t.xMin && x <= t.xMax && y >= t.yMin && y <= t.yMax
}

func makeTarget(in string) target {
	xy := strings.Split(in, ";")
	xMinMax := strings.Split(xy[0], ",")
	xMin, _ := strconv.Atoi(xMinMax[0])
	xMax, _ := strconv.Atoi(xMinMax[1])
	yMinMax := strings.Split(xy[1], ",")
	yMin, _ := strconv.Atoi(yMinMax[0])
	yMax, _ := strconv.Atoi(yMinMax[1])
	t := target{Min(xMin, xMax), Max(xMin, xMax), Min(yMin, yMax), Max(yMin, yMax)}
	return t
}

func readInput(inputFile string) target {
	input, _ := ioutil.ReadFile(inputFile)
	return makeTarget(string(input))
}

func getXRange(t target) []int {
	res := make(map[int]struct{})
	for dx := 1; dx != t.xMax+1; dx++ {
		x := 0 // do simulation
		for simDx := dx; simDx != 0; simDx-- {
			x += simDx
			if x >= t.xMin && x <= t.xMax {
				res[dx] = struct{}{}
			}
		}
	}
	xes := make([]int, 0, len(res))
	for k := range res {
		xes = append(xes, k)
	}
	return xes
}

func getYRange(t target) (yes []int) {
	for dy := -Abs(t.yMin) - 1; dy <= Abs(t.yMin)+1; dy++ {
		yes = append(yes, dy)
	}
	return
}

func applyDrag(dx int) int {
	if dx > 0 {
		return dx - 1
	} else if dx < 0 {
		return dx + 1
	}
	return dx
}

func fire(dx, dy int, t target) (int, bool) {
	height := 0
	for x, y := 0, 0; x <= t.xMax && y >= t.yMin; x, y, dx, dy = x+dx, y+dy, applyDrag(dx), dy-1 {
		height = Max(y, height)
		if t.isHit(x, y) {
			return height, true
		}
	}
	return 0, false
}

func solveFirst(t target) int {
	maxHeight := math.MinInt
	for _, dx := range getXRange(t) {
		for _, dy := range getYRange(t) {
			if height, hit := fire(dx, dy, t); hit {
				maxHeight = Max(height, maxHeight)
			}
		}
	}
	return maxHeight
}

func solveSecond(t target) int {
	hitCount := 0
	for _, dx := range getXRange(t) {
		for _, dy := range getYRange(t) {
			if _, hit := fire(dx, dy, t); hit {
				hitCount++
			}
		}
	}
	return hitCount
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
}

// utils
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Max(i, j int) int {
	if j >= i {
		return j
	}
	return i
}

func Min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

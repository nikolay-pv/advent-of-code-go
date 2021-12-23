package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type pair [2]int

type cube struct {
	x, y, z pair
	status  int
}

func (c cube) intersect(other cube) (cube, bool) {
	intersection := other
	intersection.x = [2]int{Max(c.x[0], other.x[0]), Min(c.x[1], other.x[1])}
	intersection.y = [2]int{Max(c.y[0], other.y[0]), Min(c.y[1], other.y[1])}
	intersection.z = [2]int{Max(c.z[0], other.z[0]), Min(c.z[1], other.z[1])}
	if c.status == other.status { // otherwise status of other
		intersection.flip()
	}
	return intersection, intersection.isValid()
}

func (c cube) isValid() bool {
	return (c.x[0] <= c.x[1] && c.y[0] <= c.y[1] && c.z[0] <= c.z[1])
}

func (c *cube) flip() {
	c.status *= -1
}

func (c cube) size() int {
	return (c.x[1] - c.x[0] + 1) * (c.y[1] - c.y[0] + 1) * (c.z[1] - c.z[0] + 1)
}

func readInput(inputFile string) []cube {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	values := make([]cube, len(inputStrings))
	for i, line := range inputStrings {
		values[i] = makeCube(line)
	}
	return values
}

func solveFirst(values []cube) int {
	processed := make(map[cube]int)
	for _, c := range values {
		if !isInBounds(c) {
			continue
		}
		news := make(map[cube]int)
		if c.status == 1 {
			news[c] = 1
		}
		for c2, count := range processed {
			if intersection, ok := c2.intersect(c); ok {
				news[intersection] += count
			}
		}
		for c2, v := range news {
			processed[c2] += v
		}
	}
	total := 0
	for c, count := range processed {
		total += c.status * c.size() * count
	}
	return total
}

func solveSecond(values []cube) int {
	processed := make(map[cube]int)
	for _, c := range values {
		news := make(map[cube]int)
		if c.status == 1 {
			news[c] = 1
		}
		for c2, count := range processed {
			if intersection, ok := c2.intersect(c); ok {
				news[intersection] += count
			}
		}
		for c2, v := range news {
			processed[c2] += v
		}
	}
	total := 0
	for c, count := range processed {
		total += c.status * c.size() * count
	}
	return total
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
}

// task's utils
// checks -50 50 for the first task
func isInBounds(c cube) bool {
	checkBounds := func(minMax pair) bool {
		return minMax[0] >= -50 && minMax[1] <= 50
	}
	return checkBounds(c.x) && checkBounds(c.y) && checkBounds(c.z)
}

// factories
func makePair(in string) pair {
	res := strings.Split(in, ",")
	x, _ := strconv.ParseInt(res[0], 10, 0)
	y, _ := strconv.ParseInt(res[1], 10, 0)
	return pair{int(x), int(y)}
}

func makeCube(in string) cube {
	params := strings.Split(in, " ")
	if len(params) != 4 {
		panic("expect 4 parameters for area: status, and 3 coords")
	}
	status, _ := strconv.ParseInt(params[0], 10, 0)
	if status == 0 {
		status = -1
	}
	return cube{makePair(params[1]), makePair(params[2]), makePair(params[3]), int(status)}
}

// utils

func Min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func Max(i, j int) int {
	if j >= i {
		return j
	}
	return i
}

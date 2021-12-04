package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type command struct {
	direction string
	increment int
}

func readInput(inputFile string) []command {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	values := make([]command, len(inputStrings))
	for i := range values {
		raw_command := strings.Split(inputStrings[i], " ")
		value, _ := strconv.Atoi(raw_command[1])
		values[i] = command{raw_command[0], value}
	}
	return values
}

func solveFirst(values []command) int {
	x := 0 // pointing forwards
	depth := 0
	for i := range values {
		switch values[i].direction {
		case "forward":
			x += values[i].increment
		case "down":
			depth += values[i].increment
		case "up":
			depth -= values[i].increment
		}
	}
	return depth * x
}

func solveSecond(values []command) int {
	x := 0
	aim := 0 // pointing forwards
	depth := 0
	for i := range values {
		switch values[i].direction {
		case "down":
			aim += values[i].increment
		case "up":
			aim -= values[i].increment
		case "forward":
			x += values[i].increment
			depth += aim * values[i].increment
		}
	}
	return depth * x
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
}

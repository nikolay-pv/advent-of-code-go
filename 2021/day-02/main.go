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

func read_input(input_file string) []command {
	input, _ := ioutil.ReadFile(input_file)
	input_strings := strings.Split(string(input), "\n")
	values := make([]command, len(input_strings))
	for i := range values {
		raw_command := strings.Split(input_strings[i], " ")
		value, _ := strconv.Atoi(raw_command[1])
		values[i] = command{raw_command[0], value}
	}
	return values
}

func solve_first(values []command) int {
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

func solve_second(values []command) int {
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
	values := read_input("./input.txt")
	println("Part 1: the answer is", solve_first(values))
	println("Part 2: the answer is", solve_second(values))
}

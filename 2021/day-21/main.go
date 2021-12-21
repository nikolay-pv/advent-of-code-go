package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(inputFile string) [2]int {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	values := [2]int{0, 0}
	for i := range values {
		values[i], _ = strconv.Atoi(inputStrings[i])
	}
	return values
}

func deterministicDice(initials [2]int, finalScore int) (int, [2]int) {
	scores := [2]int{0, 0}
	previous := [2]int{initials[0], initials[1]}
	p := 0
	i := 0
	currentRoll := 0
	for i != 400 {
		i += 3
		currentRoll = (i-2)%100 + (i-1)%100 + i%100
		previous[p] = (previous[p]+currentRoll-1)%10 + 1
		scores[p] += previous[p]
		if scores[p] >= finalScore {
			break
		}
		p = 1 - p // flip 1 <-> 0
	}
	return i, scores
}

func solveFirst(values [2]int) int {
	rolls, scores := deterministicDice(values, 1000)
	minScore := Min(scores[0], scores[1])
	return rolls * minScore
}

type gameData struct {
	previous   [2]int
	scores     [2]int
	player     int
	worldCount int
}

func quantumDie(initial [2]int, finalScore int) int {
	queue := make([]gameData, 1)
	queue[0] = gameData{initial, [2]int{0, 0}, 0, 1}
	winCounts := [2]int{0, 0}
	possibleRolls := make(map[int]int)
	for i := 1; i != 4; i++ {
		for j := 1; j != 4; j++ {
			for k := 1; k != 4; k++ {
				possibleRolls[i+j+k]++
			}
		}
	}
	for len(queue) != 0 {
		data := queue[0]
		queue = queue[1:]
		if data.scores[0] >= finalScore {
			winCounts[0] += data.worldCount
			continue
		}
		if data.scores[1] >= finalScore {
			winCounts[1] += data.worldCount
			continue
		}
		for k, v := range possibleRolls {
			gd := data
			gd.previous[gd.player] = (gd.previous[gd.player]+k-1)%10 + 1
			gd.scores[gd.player] += gd.previous[gd.player]
			gd.worldCount *= v
			gd.player = 1 - gd.player
			queue = append(queue, gd)
		}
	}
	return Max(winCounts[0], winCounts[1])
}

func solveSecond(values [2]int) int {
	return quantumDie(values, 21)
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
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

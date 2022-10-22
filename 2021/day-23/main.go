package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type diagram struct {
	positions    [11]int
	rooms        [9][]int // 0 is bottom of the room
	allHome      [9]bool  // 9-sized for simplicity of indexing
	roomCapacity int
}

func (d *diagram) allAtTheirHome(roomId int) bool {
	expected := roomId / 2
	for _, fish := range d.rooms[roomId] {
		if fish != expected {
			return false
		}
	}
	return true
}

func (diag *diagram) canTravel(from, to int) bool {
	// check room capacity locations availability
	switch to {
	case a, b, c, d:
		if !diag.allHome[to] {
			return false
		}
	}
	// verify that nobody blocks the path
	start, end := MinMax(from, to)
	for i := start; i != end+1; i++ {
		// not sure which of start, end is "from", thus check i != from to ignore fish there
		if diag.positions[i] != Empty && i != from {
			return false
		}
	}
	return true
}

// returns energy spend on the travel
func (diag *diagram) travel(from, to int) int {
	steps := Abs(to - from)
	var amphipod int
	switch from {
	case a, b, c, d: // leaving room
		topIndex := len(diag.rooms[from]) - 1
		steps += diag.roomCapacity - topIndex
		amphipod = diag.rooms[from][topIndex]
		diag.rooms[from] = diag.rooms[from][:topIndex]
		diag.allHome[from] = diag.allAtTheirHome(from)
	default:
		amphipod = diag.positions[from]
		diag.positions[from] = Empty
	}
	switch to {
	case a, b, c, d: // entering room
		steps += diag.roomCapacity - len(diag.rooms[to])
		diag.rooms[to] = append(diag.rooms[to], amphipod)
		diag.allHome[to] = diag.allAtTheirHome(to)
	default:
		diag.positions[to] = amphipod
	}
	return steps * fuelPM(amphipod)
}

func (d *diagram) isComplete() bool {
	for _, pos := range allDestinations {
		if !d.allHome[pos] || len(d.rooms[pos]) != d.roomCapacity {
			return false
		}
	}
	return true
}

func hash(d diagram) string {
	return fmt.Sprint(d)
}

func readInput(inputFile string, roomCapacity int) diagram {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	d := makeDiagram(roomCapacity)
	for i := roomCapacity + 1; i != 1; i-- {
		for j, r := range inputStrings[i] {
			switch r {
			case 'A', 'B', 'C', 'D':
				d.rooms[j-1][(roomCapacity+1)-i] = int(r - 'A' + 1)
			}
		}
	}
	for _, v := range allDestinations {
		d.allHome[v] = d.allAtTheirHome(v)
	}
	return d
}

// returns minimum energy to complete the given diagram
// modifies both the visited set and diagram in-place
// returns math.MaxInt if the diagram is not solvable
func findMinEnergyToComplete(visited set, d *diagram) int {
	if d.isComplete() {
		return 0
	}
	// check if that's been traveled already
	h := hash(*d)
	if v, exists := visited[h]; exists {
		return v
	}
	// search for minimum energy needed to complete the map from this point
	localMin := math.MaxInt
	// try to get home
	for pos := 0; pos < len(d.positions); pos++ {
		if d.positions[pos] == Empty {
			continue
		}
		homePos := d.positions[pos] * 2
		if d.canTravel(pos, homePos) {
			delta := d.travel(pos, homePos)
			remainder := findMinEnergyToComplete(visited, d)
			if remainder != math.MaxInt {
				localMin = Min(remainder+delta, localMin)
			}
			d.travel(homePos, pos)
		}
	}
	// get out of the room
	for _, pos := range allDestinations {
		if len(d.rooms[pos]) == 0 || d.allHome[pos] {
			continue
		}
		fish := d.rooms[pos][len(d.rooms[pos])-1]
		for _, dest := range allPositions[fish] {
			if d.canTravel(pos, dest) {
				delta := d.travel(pos, dest)
				remainder := findMinEnergyToComplete(visited, d)
				if remainder != math.MaxInt {
					localMin = Min(remainder+delta, localMin)
				}
				d.travel(dest, pos)
			}
		}
	}
	visited[h] = localMin
	return localMin
}

func solveFirst(d diagram) int {
	visited := set{}
	return findMinEnergyToComplete(visited, &d)
}

func solveSecond(d diagram) int {
	return solveFirst(d)
}

func main() {
	values := readInput("./input.txt", 2)
	println("Part 1: the answer is", solveFirst(values))
	values = readInput("./input_part2.txt", 4)
	println("Part 2: the answer is", solveSecond(values))
}

// utils
const (
	Empty = 0
	A     = 1
	B     = 2
	C     = 3
	D     = 4
)

// destinations
const (
	a = 2
	b = 4
	c = 6
	d = 8
)

type set map[string]int

var allDestinations = [4]int{2, 4, 6, 8}
var allPositions = [5][7]int{ // sorted in abs distance to destination
	{0, 1, 3, 5, 7, 9, 10}, // unused
	{1, 3, 0, 5, 7, 9, 10},
	{3, 5, 1, 7, 0, 9, 10},
	{5, 7, 3, 9, 10, 1, 0},
	{7, 9, 10, 5, 3, 1, 0},
}

// fuel per move
func fuelPM(amphipod int) int {
	switch amphipod {
	case A:
		return 1
	case B:
		return 10
	case C:
		return 100
	case D:
		return 1000
	}
	panic("unknown amphipod")
}

func makeDiagram(roomCapacity int) diagram {
	diag := diagram{}
	diag.positions = [11]int{}
	diag.rooms = [9][]int{}
	for _, v := range allDestinations {
		diag.rooms[v] = make([]int, roomCapacity)
	}
	diag.allHome = [9]bool{}
	diag.roomCapacity = roomCapacity
	return diag
}

// utils
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func MinMax(i, j int) (int, int) {
	if i > j {
		return j, i
	}
	return i, j
}

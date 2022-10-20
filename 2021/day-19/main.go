package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

func (p point) at(i int) int {
	switch i {
	case 0:
		return p.x
	case 1:
		return p.y
	case 2:
		return p.z
	default:
		panic("out of bound")
	}
}

func add(lhs, rhs point) point {
	return point{lhs.x + rhs.x, lhs.y + rhs.y, lhs.z + rhs.z}
}

func subtract(lhs, rhs point) point {
	return point{lhs.x - rhs.x, lhs.y - rhs.y, lhs.z - rhs.z}
}

func (start point) distance(end point) int {
	return (end.x-start.x)*(end.x-start.x) + (end.y-start.y)*(end.y-start.y) + (end.z-start.z)*(end.z-start.z)
}

// represent rotation matrix as a point per row
type rotation struct {
	r1, r2, r3 point
}

func rotate(p point, r rotation) point {
	return point{
		p.at(0)*r.r1.at(0) + p.at(1)*r.r1.at(1) + p.at(2)*r.r1.at(2),
		p.at(0)*r.r2.at(0) + p.at(1)*r.r2.at(1) + p.at(2)*r.r2.at(2),
		p.at(0)*r.r3.at(0) + p.at(1)*r.r3.at(1) + p.at(2)*r.r3.at(2)}
}

type scannerData struct {
	points   []point
	location point
}

func (sd *scannerData) transform(r rotation, translation point) {
	if sd.location.x != 0 || sd.location.y != 0 || sd.location.z != 0 {
		panic("expect the mispositioned scanner")
	}
	sd.location = translation
	for i := 0; i < len(sd.points); i++ {
		sd.points[i] = add(rotate(sd.points[i], r), translation)
	}
}

// will store two pointes per each (ref and source) scanner to find the correct location and rotation of a scanner
type positioningData struct {
	references []point
	source     []point
}

func makePositioner() func(reference, current *scannerData, pd positioningData) {
	rotations := []rotation{
		{point{1, 0, 0}, point{0, 1, 0}, point{0, 0, 1}},
		{point{0, 0, 1}, point{0, 1, 0}, point{-1, 0, 0}},
		{point{-1, 0, 0}, point{0, 1, 0}, point{0, 0, -1}},
		{point{0, 0, -1}, point{0, 1, 0}, point{1, 0, 0}},
		{point{0, -1, 0}, point{1, 0, 0}, point{0, 0, 1}},
		{point{0, 0, 1}, point{1, 0, 0}, point{0, 1, 0}},
		{point{0, 1, 0}, point{1, 0, 0}, point{0, 0, -1}},
		{point{0, 0, -1}, point{1, 0, 0}, point{0, -1, 0}},
		{point{0, 1, 0}, point{-1, 0, 0}, point{0, 0, 1}},
		{point{0, 0, 1}, point{-1, 0, 0}, point{0, -1, 0}},
		{point{0, -1, 0}, point{-1, 0, 0}, point{0, 0, -1}},
		{point{0, 0, -1}, point{-1, 0, 0}, point{0, 1, 0}},
		{point{1, 0, 0}, point{0, 0, -1}, point{0, 1, 0}},
		{point{0, 1, 0}, point{0, 0, -1}, point{-1, 0, 0}},
		{point{-1, 0, 0}, point{0, 0, -1}, point{0, -1, 0}},
		{point{0, -1, 0}, point{0, 0, -1}, point{1, 0, 0}},
		{point{1, 0, 0}, point{0, -1, 0}, point{0, 0, -1}},
		{point{0, 0, -1}, point{0, -1, 0}, point{-1, 0, 0}},
		{point{-1, 0, 0}, point{0, -1, 0}, point{0, 0, 1}},
		{point{0, 0, 1}, point{0, -1, 0}, point{1, 0, 0}},
		{point{1, 0, 0}, point{0, 0, 1}, point{0, -1, 0}},
		{point{0, -1, 0}, point{0, 0, 1}, point{-1, 0, 0}},
		{point{-1, 0, 0}, point{0, 0, 1}, point{0, 1, 0}},
		{point{0, 1, 0}, point{0, 0, 1}, point{1, 0, 0}}}
	return func(reference, current *scannerData, pd positioningData) {
		var d1, d2, s1, s2 point
		for _, r := range rotations {
			s1, s2 = rotate(pd.source[0], r), rotate(pd.source[1], r)
			d1, d2 = subtract(pd.references[0], s1), subtract(pd.references[1], s2)
			if d1 == d2 {
				current.transform(r, d1)
				return
			}
			d1, d2 = subtract(pd.references[0], s2), subtract(pd.references[1], s1)
			if d1 == d2 {
				current.transform(r, d1)
				return
			}
		}
	}
}

type distanceToIndicesMap map[int][]int

func computeCache(distanceToPoints []distanceToIndicesMap, relativePoints []set, scanners []scannerData) {
	d := 0
	for scanId, scanner := range scanners {
		relativePoints[scanId] = make(set)
		distanceToPoints[scanId] = make(distanceToIndicesMap)
		for i := range scanner.points {
			for j := i + 1; j < len(scanner.points); j++ {
				d = scanner.points[i].distance(scanner.points[j])
				relativePoints[scanId].insert(d)
				distanceToPoints[scanId][d] = append(distanceToPoints[scanId][d], i, j)
			}
		}
	}
}

func findPositionData(reference, source scannerData, refMaps, sourceMaps distanceToIndicesMap, overlap set) positioningData {
	pd := positioningData{}
	for k := range overlap {
		if len(refMaps[k]) == 2 && len(sourceMaps[k]) == 2 {
			a, b := refMaps[k][0], refMaps[k][1]
			pd.references = append(pd.references, reference.points[a], reference.points[b])
			a, b = sourceMaps[k][0], sourceMaps[k][1]
			pd.source = append(pd.source, source.points[a], source.points[b])
			break
		}
	}
	return pd
}

func countUniquePoints(scanners []scannerData) int {
	unique := make(map[point]struct{})
	for _, sd := range scanners {
		for _, p := range sd.points {
			unique[p] = struct{}{}
		}
	}
	return len(unique)
}

func readInput(inputFile string) []scannerData {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n\n")
	scanners := make([]scannerData, len(inputStrings))
	for i, data := range inputStrings {
		scanners[i] = makeScannerData(data)
	}
	return scanners
}

func solveFirst(scanners []scannerData, threshold int) int {
	relativePoints := make([]set, len(scanners))
	distanceToPoints := make([]distanceToIndicesMap, len(scanners))
	computeCache(distanceToPoints, relativePoints, scanners)
	position := makePositioner()

	positioned := make([]bool, len(scanners))
	positioned[0] = true
	unalignedCount := len(scanners) - 1
	var overlap set
	for unalignedCount != 0 {
		for unalignedId := range positioned {
			if positioned[unalignedId] == true {
				continue
			}
			for alignedId := range positioned {
				if positioned[alignedId] == false {
					continue
				}
				overlap = intersect(relativePoints[alignedId], relativePoints[unalignedId])
				if len(overlap) >= threshold*(threshold-1)/2 {
					pd := findPositionData(scanners[alignedId], scanners[unalignedId], distanceToPoints[alignedId], distanceToPoints[unalignedId], overlap)
					position(&scanners[alignedId], &scanners[unalignedId], pd)
					// println("Matched scanner", alignedId, "with", unalignedId)
					// println("Positioned scanner", unalignedId, "at", scanners[unalignedId].scanPosition.x, scanners[unalignedId].scanPosition.y, scanners[unalignedId].scanPosition.z)
					positioned[unalignedId] = true
					unalignedCount -= 1
					break
				}
			}
		}
	}

	return countUniquePoints(scanners)
}

func manhattanDistance(lhs, rhs point) int {
	return Abs(lhs.x-rhs.x) + Abs(lhs.y-rhs.y) + Abs(lhs.z-rhs.z)
}

func solveSecond(scanners []scannerData) int {
	distance := 0
	var d int
	for i := 0; i < len(scanners); i++ {
		for j := i + 1; j < len(scanners); j++ {
			d = manhattanDistance(scanners[i].location, scanners[j].location)
			distance = Max(distance, d)
		}
	}
	return distance
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values, 12))
	println("Part 2: the answer is", solveSecond(values))
}

// factories

func makePoint(s string) point {
	coordinates := strings.Split(s, ",")
	x, _ := strconv.Atoi(coordinates[0])
	y, _ := strconv.Atoi(coordinates[1])
	z, _ := strconv.Atoi(coordinates[2])
	return point{x, y, z}
}

func makeScannerData(s string) scannerData {
	lines := strings.Split(s, "\n")
	points := make([]point, len(lines))
	for i, line := range lines {
		points[i] = makePoint(line)
	}
	return scannerData{points, point{}}
}

// utils

type set map[int]struct{}

func (s *set) insert(i int) {
	if s == nil {
		*s = make(set)
	}
	(*s)[i] = struct{}{}
}

func intersect(s1, s2 set) set {
	intersection := make(map[int]struct{})
	if len(s1) > len(s2) {
		s1, s2 = s2, s1 // better to iterate over a shorter set
	}
	for k := range s1 {
		if _, ok := s2[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func Max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

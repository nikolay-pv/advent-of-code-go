package main

import (
	"io/ioutil"
	"strings"
)

type image [][]int8

func (img image) binary(around point) int {
	binary := 0
	neighbors := around.neighboring()
	for j, p := range neighbors {
		x := bound(p[0], 0, len(img)-1)
		y := bound(p[1], 0, len(img[0])-1)
		v := img[x][y]
		binary += int(v) << (len(neighbors) - j - 1)
	}
	return binary
}

func (img image) addPadding(n int) image {
	rows := len(img) + 2*n
	cols := len(img[0]) + 2*n
	out := make([][]int8, len(img)+2*n)
	for i := 0; i != rows; i++ {
		out[i] = make([]int8, cols)
		for j := 0; j != rows; j++ {
			if i >= n && i < n+len(img) && j >= n && j < n+len(img[0]) {
				out[i][j] = img[i-n][j-n]
			}
		}
	}
	return out
}

func (img image) countLit() int {
	count := 0
	for _, row := range img {
		for _, v := range row {
			if v == 1 {
				count++
			}
		}
	}
	return count
}

// algorithm
type imgAlgorithm map[int]int

func (img imgAlgorithm) outputForBinary(binary int) int {
	if v, ok := img[binary]; ok {
		return v
	}
	return 0
}

// utility struct to easy test
type input struct {
	a   imgAlgorithm
	img image
}

func readInput(inputFile string) input {
	in, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(in), "\n\n")
	a := makeAlgorithm(inputStrings[0])
	img := makeImage(inputStrings[1])
	return input{a, img}
}

// solution
func enchanceImage(a imgAlgorithm, img image, times int) image {
	img = img.addPadding(times * 2)
	for ; times != 0; times-- {
		outputImg := makeNewImage(len(img), len(img[0]))
		for i, row := range img {
			for j := range row {
				binary := img.binary(point{i, j})
				if a.outputForBinary(binary) != 0 {
					outputImg[i][j] = 1
				}
			}
		}
		img = outputImg
	}
	return img
}

func solveFirst(a imgAlgorithm, img image) int {
	img = enchanceImage(a, img, 2)
	// img.print()
	return img.countLit()
}

func solveSecond(a imgAlgorithm, img image) int {
	img = enchanceImage(a, img, 50)
	// img.print()
	return img.countLit()
}

func main() {
	in := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(in.a, in.img))
	in = readInput("./input.txt")
	println("Part 2: the answer is", solveSecond(in.a, in.img))
}

// factories

func makeImage(in string) image {
	lines := strings.Split(in, "\n")
	img := make([][]int8, len(lines))
	for i, line := range lines {
		img[i] = make([]int8, len(line))
		for j, r := range line {
			if r == '#' {
				img[i][j] = 1
			} else {
				img[i][j] = 0
			}
		}
	}
	return img
}

func makeNewImage(row, col int) image {
	img := make([][]int8, row)
	for i := 0; i != row; i++ {
		img[i] = make([]int8, col)
		for j := range img[i] {
			img[i][j] = 0
		}
	}
	return img
}

func makeAlgorithm(in string) imgAlgorithm {
	a := imgAlgorithm{}
	for i, v := range in {
		if v == '#' {
			a[i] = 1
		}
	}
	return a
}

// utils

// func (i image) print() {
// 	for _, row := range i {
// 		for _, val := range row {
// 			if val == 1 {
// 				print("#")
// 			} else {
// 				print(".")
// 			}
// 		}
// 		print("\n")
// 	}
// 	print("\n")
// }

func bound(i, min, max int) int {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}

// poor point type
type point [2]int

// includes itself
func (p point) neighboring() []point {
	points := make([]point, 0, 9)
	for i := -1; i != 2; i++ {
		for j := -1; j != 2; j++ {
			points = append(points, point{p[0] + i, p[1] + j})
		}
	}
	return points
}

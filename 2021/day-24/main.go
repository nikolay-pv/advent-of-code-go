package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readInput(inputFile string) []string {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	return inputStrings
}

func compile(line string) string {
	switch line[:3] {
	case "inp":
		return fmt.Sprintf(`%c = input[lastConsumedInputIdx]
    lastConsumedInputIdx++
`, line[4])
	case "add":
		return fmt.Sprintf("%c += %s", line[4], line[6:])
	case "mul":
		return fmt.Sprintf("%c *= %s", line[4], line[6:])
	case "div":
		return fmt.Sprintf("%c /= %s", line[4], line[6:])
	case "mod":
		return fmt.Sprintf("%c = %c %% %s", line[4], line[4], line[6:])
	case "eql":
		return fmt.Sprintf("%c = btoi(%c == %s)", line[4], line[4], line[6:])
	}
	return ""
}

func generateGo(values []string) {
	header := `package main

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func helper(input [14]int) int {
	lastConsumedInputIdx := 0
	w, x, y, z := 0, 0, 0, 0
`

	body := ""
	for _, v := range values {
		body += "     "
		body += compile(v)
		body += "\n"
	}

	footer := `
	return z
}
`
	ioutil.WriteFile("helper.go", []byte(header+body+footer), 0644)
}

func solveFirst(values []string) int {
	generateGo(values)
	return 0
}

func solveSecond(values []string) int {
	return len(values)
}

func main() {
	values := readInput("./input.txt")
	// println("Part 1: the answer is", solveFirst(values))
	solveFirst(values)
	// println("Part 2: the answer is", solveSecond(values))
}

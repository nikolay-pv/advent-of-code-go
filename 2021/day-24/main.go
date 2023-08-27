package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
		return fmt.Sprintf(`
	%c = input[lastConsumedInputIdx]
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

func decrement(input *[14]int) {
	for i := 13; i != -1; i-- {
		if input[i] != 1 {
			input[i]--
			for j := i + 1; j < 14; j++ {
				input[j] = 9
			}
			return
		}
	}
	panic("all numbers are ones in input")
}

func printAsNumber(values *[14]int) {
	for _, v := range values {
		print(v)
	}
	println()
}

func solveFirst() int {
	values := [14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	print("Searching for the max value:")
	for helper(values) != 0 {
		// printAsNumber(&values)
		decrement(&values)
	}
	println("---")
	printAsNumber(&values)
	return 0
}

func solveSecond(values []string) int {
	return len(values)
}

func printHelp() {
	print("pass --generate to generate helper file with your input (auto-reads input.txt file), you will have to recompile the program")
	print("or pass --find to search for largest number for that input")
	return
}

func main() {
	values := readInput("./input.txt")
	if len(os.Args) == 1 {
		printHelp()
		return
	}
	if os.Args[1] == "--generate" {
		generateGo(values)
	} else if os.Args[1] == "--find" {
		solveFirst()
	} else {
		printHelp()
	}
	// println("Part 1: the answer is", solveFirst(values))
	// println("Part 2: the answer is", solveSecond(values))
}

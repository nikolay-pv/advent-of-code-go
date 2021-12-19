package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readInput(inputFile string) []*node {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	nodes := make([]*node, len(inputStrings))
	for i, line := range inputStrings {
		nodes[i] = makeNodes(line)
	}
	return nodes
}

type node struct {
	number              int
	left, right, parent *node
}

func (n *node) magnitude() int {
	if n.left == nil && n.right == nil {
		return n.number
	}
	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func (n *node) depth() int {
	depth := -1
	for p := n; p != nil; p = p.parent {
		if p.isGrouping() {
			depth++
		}
	}
	return depth
}

func (n *node) shouldExplode() bool {
	return n.depth() == 4
}

func (n *node) isGrouping() bool {
	return n.number == -1
}

func (n *node) print() string {
	if n.isGrouping() {
		return fmt.Sprintf("[%s,%s]", n.left.print(), n.right.print())
	} else {
		return fmt.Sprintf("%d", n.number)
	}
}

const (
	left  = 0
	right = 1
)

// finds a node in the graph which appears left or right as if the graph was
// written out to a string
func findAdjacent(n *node, side int) *node {
	target := n.parent
	for previous := n; target != nil; previous, target = target, target.parent {
		if side == left {
			if target.right == previous {
				break
			}
		} else {
			if target.left == previous {
				break
			}
		}
	}
	if target == nil {
		return nil
	}
	if side == left {
		target = target.left
		for target.isGrouping() {
			target = target.right
		}
	} else {
		target = target.right
		for target.isGrouping() {
			target = target.left
		}
	}
	return target
}

func explode(n *node) *node {
	adjLeft := findAdjacent(n, left)
	if adjLeft != nil {
		adjLeft.number += n.left.number
	}
	if adjRight := findAdjacent(n, right); adjRight != nil {
		adjRight.number += n.right.number
	}
	// nullify itself
	n.number = 0
	n.left = nil
	n.right = nil
	return adjLeft
}

func split(n *node) *node {
	n.left = makeLeafNode(n.number / 2)
	n.left.parent = n
	n.right = makeLeafNode((n.number-1)/2 + 1)
	n.right.parent = n
	n.number = -1
	return n
}

func reduce(parent *node) {
	// println(parent.print())
	queue := make([]*node, 1)
	queue[0] = parent
	for len(queue) != 0 {
		n := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if n.shouldExplode() {
			explode(n)
			// println(parent.print())
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
		if n.left != nil {
			queue = append(queue, n.left)
		}
	}
	// println(parent.print())
	// split
	queue = make([]*node, 1)
	queue[0] = parent
	for len(queue) != 0 {
		n := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if n.number >= 10 {
			split(n)
			// println(parent.print())
			if n.shouldExplode() {
				explode(n)
				queue = queue[:1]
				queue[0] = parent
				// println(parent.print())
				continue
			}
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
		if n.left != nil {
			queue = append(queue, n.left)
		}
	}
}

func solveFirst(nodes []*node) int {
	parent := nodes[0]
	for i := 1; i != len(nodes); i++ {
		reduce(parent)
		reduce(nodes[i])
		parent = makeGroupingNode(parent, nodes[i])
	}
	reduce(parent)
	return parent.magnitude()
}

func solveSecond(nodes []*node) int {
	lines := make([]string, len(nodes))
	for i, n := range nodes {
		reduce(n)
		lines[i] = n.print()
	}
	max := 0
	for i := 0; i != len(lines)-1; i++ {
		for j := i + 1; j != len(lines); j++ {
			n := makeGroupingNode(makeNodes(lines[i]), makeNodes(lines[j]))
			reduce(n)
			max = Max(max, n.magnitude())
			n = makeGroupingNode(makeNodes(lines[j]), makeNodes(lines[i]))
			reduce(n)
			max = Max(max, n.magnitude())
		}
	}
	return max
}

func main() {
	p := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(p))
	p = readInput("./input.txt")
	println("Part 2: the answer is", solveSecond(p))
}

// factories

func makeLeafNode(number int) *node {
	n := node{}
	n.number = number
	return &n
}

func makeGroupingNode(left, right *node) *node {
	n := makeLeafNode(-1)
	n.left = left
	n.left.parent = n
	n.right = right
	n.right.parent = n
	return n
}

func makeNodes(in string) (p *node) {
	orphans := make([]*node, 0)
	for _, r := range in {
		switch r {
		case '[', ',':
			continue
		case ']':
			left := orphans[len(orphans)-2]
			right := orphans[len(orphans)-1]
			orphans = orphans[:len(orphans)-2]
			groupNode := makeGroupingNode(left, right)
			p = groupNode
			orphans = append(orphans, groupNode)
		default:
			orphans = append(orphans, makeLeafNode(int(r-'0')))
		}
	}
	return
}

// utils

func Max(i, j int) int {
	if j >= i {
		return j
	}
	return i
}

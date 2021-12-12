package main

import (
	"io/ioutil"
	"strings"
	"unicode"
)

func IsLower(label string) bool {
	for _, r := range label {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

type Node struct {
	label string
}

type Graph struct {
	nodes      []Node
	edges      map[Node][]*Node
	start, end *Node
}

func (g *Graph) AddNode(n Node) *Node {
	g.nodes = append(g.nodes, n)
	res := &g.nodes[len(g.nodes)-1]
	if n.label == "start" {
		g.start = res
	}
	if n.label == "end" {
		g.end = res
	}
	return res
}

func (g *Graph) AddEdge(n0, n1 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n0] = append(g.edges[*n0], n1)
	g.edges[*n1] = append(g.edges[*n1], n0)
}

func readInput(inputFile string) Graph {
	input, _ := ioutil.ReadFile(inputFile)
	inputStrings := strings.Split(string(input), "\n")
	nodesMap := make(map[string]*Node)
	g := Graph{}
	makeNode := func(label string) *Node {
		n, ok := nodesMap[label]
		if !ok {
			n = g.AddNode(Node{label})
			nodesMap[label] = n
		}
		return n
	}
	for _, line := range inputStrings {
		nodes := strings.Split(line, "-")
		n0 := makeNode(nodes[0])
		n1 := makeNode(nodes[1])
		g.AddEdge(n0, n1)
	}
	return g
}

func didVisit(n *Node, nodes []*Node) bool {
	for i := len(nodes) - 1; i != -1; i-- {
		if nodes[i] == n {
			return true
		}
	}
	return false
}

func makeVisited(n *Node) []*Node {
	v := make([]*Node, 1)
	v[0] = n
	return v
}

func solveFirst(g Graph) int {
	paths := 0
	queue := make([][]*Node, 0)
	queue = append(queue, makeVisited(g.start))
	for len(queue) != 0 {
		path := queue[0]
		queue = queue[1:]
		for _, n := range g.edges[*path[len(path)-1]] {
			if n == g.end {
				paths++
				continue
			}
			if IsLower(n.label) && didVisit(n, path) {
				continue
			}
			v := make([]*Node, len(path))
			copy(v, path)
			queue = append(queue, append(v, n))
		}
	}
	return paths
}

func solveSecond(values Graph) int {
	return len(values.nodes[0].label)
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	// println("Part 2: the answer is", solveSecond(values))
}

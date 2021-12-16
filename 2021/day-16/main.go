package main

import (
	"io/ioutil"
	"math"
	"strconv"
)

type offset int

type packet struct {
	version, id int
	literal     int
	subpackets  []packet
}

func preprocessMessage(hex string) string {
	mapping := map[rune]string{'0': "0000", '1': "0001", '2': "0010", '3': "0011",
		'4': "0100", '5': "0101", '6': "0110", '7': "0111", '8': "1000", '9': "1001",
		'A': "1010", 'B': "1011", 'C': "1100", 'D': "1101", 'E': "1110", 'F': "1111"}
	binary := make([]byte, 4*len(hex))
	for i, r := range hex {
		hexToBin := mapping[r]
		for j := range hexToBin {
			binary[i*4+j] = hexToBin[j]
		}
	}
	return string(binary)
}

func decodeHeader(binary string, start offset) (int, int, offset) {
	version, _ := strconv.ParseInt(binary[start:start+3], 2, 0)
	start += 3
	id, _ := strconv.ParseInt(binary[start:start+3], 2, 0)
	start += 3
	return int(version), int(id), start
}

// start points to the first bit of literal
func decodeLiteral(binary string, start offset) (int, offset) {
	res := make([]byte, 0)
	for binary[start] == byte('1') {
		res = append(res, binary[start+1:start+5]...)
		start += 5
	}
	// one more time as we met 0
	res = append(res, binary[start+1:start+5]...)
	value, _ := strconv.ParseInt(string(res), 2, 0)
	return int(value), start + 5
}

func decodeSubpackets(binary string, start offset) ([]packet, offset) {
	subpackets := make([]packet, 0)
	switch binary[start] {
	case byte('0'):
		start++
		length, _ := strconv.ParseInt(binary[start:start+15], 2, 0)
		start += 15
		for length != 0 {
			subP, offset := decodePacket(binary, start)
			subpackets = append(subpackets, subP)
			length -= int64(offset - start)
			start = offset
		}
	case byte('1'):
		start++
		count, _ := strconv.ParseInt(binary[start:start+11], 2, 0)
		start += 11
		for count != 0 {
			subP, offset := decodePacket(binary, start)
			subpackets = append(subpackets, subP)
			count--
			start = offset
		}
	}
	return subpackets, start
}

func decodeLiteralPacket(binary string, start offset) (packet, offset) {
	p := packet{}
	p.version, p.id, start = decodeHeader(binary, start)
	p.literal, start = decodeLiteral(binary, start)
	return p, start
}

func decodeOperatorPacket(binary string, start offset) (packet, offset) {
	p := packet{}
	p.version, p.id, start = decodeHeader(binary, start)
	p.subpackets, start = decodeSubpackets(binary, start)
	return p, start
}

func decodePacket(binary string, start offset) (packet, offset) {
	p := packet{}
	p.version, p.id, start = decodeHeader(binary, start)
	switch p.id {
	case literal:
		p.literal, start = decodeLiteral(binary, start)
	default:
		p.subpackets, start = decodeSubpackets(binary, start)
	}
	return p, start
}

func readInput(inputFile string) string {
	input, _ := ioutil.ReadFile(inputFile)
	binary := preprocessMessage(string(input))
	return binary
}

func countVersions(p packet, versions *int) {
	*versions += p.version
	for _, p2 := range p.subpackets {
		countVersions(p2, versions)
	}
}

func readPackets(binary string) []packet {
	start := offset(0)
	packets := make([]packet, 0)
	for len(binary)-int(start) > 10 {
		packet, offset := decodePacket(binary, start)
		packets = append(packets, packet)
		start = offset
	}
	return packets
}

func solveFirst(binary string) int {
	packets := readPackets(binary)
	versionSum := 0
	for _, p := range packets {
		countVersions(p, &versionSum)
	}
	return versionSum
}

const (
	sum     = 0
	product = 1
	minimum = 2
	maximum = 3
	literal = 4
	greater = 5
	less    = 6
	equal   = 7
)

func (p packet) evaluate() int {
	res := 0
	switch p.id {
	case sum:
		for _, p2 := range p.subpackets {
			res += p2.evaluate()
		}
	case product:
		res = 1
		for _, p2 := range p.subpackets {
			res *= p2.evaluate()
		}
	case minimum:
		res = math.MaxInt
		for _, p2 := range p.subpackets {
			res = Min(res, p2.evaluate())
		}
	case maximum:
		res = math.MinInt
		for _, p2 := range p.subpackets {
			res = Max(res, p2.evaluate())
		}
	case literal:
		res = p.literal
	case greater:
		if p.subpackets[0].evaluate() > p.subpackets[1].evaluate() {
			res = 1
		}
	case less:
		if p.subpackets[0].evaluate() < p.subpackets[1].evaluate() {
			res = 1
		}
	case equal:
		if p.subpackets[0].evaluate() == p.subpackets[1].evaluate() {
			res = 1
		}
	}
	return res
}

func solveSecond(binary string) int {
	packets := readPackets(binary)
	if len(packets) != 1 {
		panic("must be single packet")
	}
	return packets[0].evaluate()
}

func main() {
	values := readInput("./input.txt")
	println("Part 1: the answer is", solveFirst(values))
	println("Part 2: the answer is", solveSecond(values))
}

// utils
func Max(lhs, rhs int) int {
	if rhs >= lhs {
		return rhs
	}
	return lhs
}

func Min(lhs, rhs int) int {
	if lhs <= rhs {
		return lhs
	}
	return rhs
}

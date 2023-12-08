package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := Part1GetCnt(scanner)
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := Part2GetCnt(scanner)
	fmt.Println(sum)
}

type Turn int8

const (
	L Turn = iota
	R
)

type Node struct {
	left      string
	right     string
	hasSuffix bool
}

func Part1GetCnt(scanner *bufio.Scanner) int {
	turns := Part1ParseTurns(scanner)
	m := Part1ParseNodes(scanner, "ZZZ")
	current := m["AAA"]
	return getCnt(turns, current, m)
}

func getCnt(turns []Turn, current Node, m map[string]Node) int {
	stepCnt := 0
	for {
		for _, turn := range turns {
			stepCnt++
			switch turn {
			case L:
				current = m[current.left]
			case R:
				current = m[current.right]
			}
			if current.hasSuffix {
				return stepCnt
			}
		}
	}
	panic("not reachable")
}

func Part2GetCnt(scanner *bufio.Scanner) int {
	turns := Part1ParseTurns(scanner)
	m := Part1ParseNodes(scanner, "Z")

	currents := []Node{}
	for k, v := range m {
		if strings.HasSuffix(k, "A") {
			currents = append(currents, v)
		}
	}
	paths := []int{}

	for _, current := range currents {
		paths = append(paths, getCnt(turns, current, m))
	}

	return LCM(paths...)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Part1ParseTurns(scanner *bufio.Scanner) []Turn {
	turns := []Turn{}
	scanner.Scan()
	s := scanner.Text()
	for _, turn := range s {
		switch turn {
		case 'R':
			turns = append(turns, R)
		case 'L':
			turns = append(turns, L)
		}
	}
	scanner.Scan()
	return turns
}

func Part1ParseNodes(scanner *bufio.Scanner, suffix string) map[string]Node {
	m := map[string]Node{}

	for scanner.Scan() {
		s := strings.ReplaceAll(scanner.Text(), "  ", " ")
		nameAndNode := strings.Split(s, " = ")
		name := nameAndNode[0]
		nodes := strings.Split(strings.Trim(nameAndNode[1], "()"), ", ")
		m[name] = Node{nodes[0], nodes[1], strings.HasSuffix(name, suffix)}
	}
	return m
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1()
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

type Turn int8

const (
	L Turn = iota
	R
)

type Node struct {
	left  string
	right string
}

func Part1GetCnt(scanner *bufio.Scanner) int {
	turns := Part1ParseTurns(scanner)
	m := map[string]Node{}
	stepCnt := 0

	for scanner.Scan() {
		s := strings.ReplaceAll(scanner.Text(), "  ", " ")
		nameAndNode := strings.Split(s, " = ")
		name := nameAndNode[0]
		nodes := strings.Split(strings.Trim(nameAndNode[1], "()"), ", ")
		m[name] = Node{nodes[0], nodes[1]}
	}
	current := m["AAA"]
	for {
		for _, turn := range turns {
			stepCnt++
			selected := ""
			switch turn {
			case L:
				selected = current.left
				current = m[selected]
			case R:
				selected = current.right
				current = m[selected]
			}
			if selected == "ZZZ" {
				return stepCnt
			}
		}
	}
	panic("not reachable")
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

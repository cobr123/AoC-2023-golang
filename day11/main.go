package main

import (
	"bufio"
	"fmt"
	"os"
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

	sum := Part1GetSum(scanner)
	fmt.Println(sum)
}

func Part1GetSum(scanner *bufio.Scanner) int {
	sum := 0
	universe := [][]rune{}
	for scanner.Scan() {
		s := scanner.Text()
		universe = append(universe, Part1ParseLine(s))
	}
	expandedUniverse := Part1ExpandUniverse(universe)
	fmt.Println(expandedUniverse)
	return sum
}

func Part1ParseUniverse(scanner *bufio.Scanner) [][]rune {
	universe := [][]rune{}
	for scanner.Scan() {
		s := scanner.Text()
		universe = append(universe, Part1ParseLine(s))
	}
	return universe
}

func Part1ParseLine(s string) []rune {
	line := []rune{}
	for _, ch := range s {
		line = append(line, ch)
	}
	return line
}

func Part1ExpandUniverse(universe [][]rune) [][]rune {
	return Part1ExpandUniverseV(Part1ExpandUniverseH(universe))
}

func Part1ExpandUniverseH(universe [][]rune) [][]rune {
	expandedUniverse := [][]rune{}
	for _, line := range universe {
		noGalaxies := true
		for _, ch := range line {
			if ch == '#' {
				noGalaxies = false
			}
		}
		if noGalaxies {
			expandedUniverse = append(expandedUniverse, line)
		}
		expandedUniverse = append(expandedUniverse, line)
	}
	return expandedUniverse
}

func Part1ExpandUniverseV(universe [][]rune) [][]rune {
	return transpose(Part1ExpandUniverseH(transpose(universe)))
}

func transpose(slice [][]rune) [][]rune {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]rune, xl)
	for i := range result {
		result[i] = make([]rune, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

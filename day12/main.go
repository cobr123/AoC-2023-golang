package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	sum := Part1GetSum(scanner)
	fmt.Println(sum)
}

type Line struct {
	Runes  []rune
	Counts []int
}

func Part1GetSum(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		line := Part1ParseLine(s)
		sum += Part1FindVariants(line)
	}
	return sum
}

func Part1ParseLine(s string) Line {
	runesAndCounts := strings.Split(s, " ")
	runes := Part1ParseRunes(runesAndCounts[0])
	counts := Part1ParseCounts(runesAndCounts[1])
	return Line{runes, counts}
}

func Part1ParseRunes(s string) []rune {
	runes := []rune{}
	for _, item := range s {
		runes = append(runes, item)
	}
	return runes
}

func Part1ParseCounts(s string) []int {
	numbers := []int{}
	arr := strings.Split(s, ",")
	for _, item := range arr {
		n, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}

func Part1FindVariants(line Line) int {
	return 1
}

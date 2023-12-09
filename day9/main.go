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

func Part1GetSum(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		numbers := Part1ParseNumbers(s)
		sum += Part1FindNext(numbers)
	}
	return sum
}

func Part1ParseNumbers(s string) []int {
	numbers := []int{}
	arr := strings.Split(s, " ")
	for _, item := range arr {
		n, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}

func Part1FindNext(numbers []int) int {
	lines := Part1GetDiffLines(numbers)
	for i := len(lines) - 1; i > 0; i-- {
		prevLastIdx := len(lines[i-1]) - 1
		lastIdx := len(lines[i]) - 1
		newLastValue := lines[i-1][prevLastIdx] + lines[i][lastIdx]
		lines[i-1] = append(lines[i-1], newLastValue)
	}
	return lines[0][len(lines[0])-1]
}

func Part1GetDiffLines(numbers []int) [][]int {
	lines := [][]int{}
	lines = append(lines, numbers)
	for {
		diffs, allZeros := Part1GetDiffs(numbers)
		numbers = diffs
		lines = append(lines, diffs)
		if allZeros {
			break
		}
	}
	return lines
}

func Part1GetDiffs(numbers []int) ([]int, bool) {
	diffs := []int{}
	allZeros := true
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if diff != 0 {
			allZeros = false
		}
		diffs = append(diffs, diff)
	}
	return diffs, allZeros
}

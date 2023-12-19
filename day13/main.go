package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

	sum := Part1GetSum(scanner)
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := Part2GetSum(scanner)
	fmt.Println(sum)
}

type Mirror struct {
	lines [][]rune
}

func Part1GetSum(scanner *bufio.Scanner) int {
	mirrors := []Mirror{}
	mirror := Mirror{}
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			mirrors = append(mirrors, mirror)
			mirror = Mirror{}
		} else {
			line := []rune(s)
			mirror.lines = append(mirror.lines, line)
		}
	}
	mirrors = append(mirrors, mirror)

	sum := 0
	for _, mirror := range mirrors {
		numberOfRowsBeforeReflection, numberOfColumnsBeforeReflection := Part1GetRowColBeforeReflection(mirror)
		sum += numberOfColumnsBeforeReflection + 100*numberOfRowsBeforeReflection
	}
	return sum
}

func Part1GetRowColBeforeReflection(mirror Mirror) (int, int) {
	numberOfRowsBeforeReflection, err := findReflection(mirror.lines)
	if err == nil {
		return numberOfRowsBeforeReflection, 0
	}
	numberOfColumnsBeforeReflection, err := findReflection(transpose(mirror.lines))
	if err == nil {
		return 0, numberOfColumnsBeforeReflection
	}
	for _, line := range mirror.lines {
		fmt.Println(string(line))
	}
	panic("reflection not found")
}

func findReflection(lines [][]rune) (int, error) {
loop:
	for i := 1; i < len(lines); i++ {
		if string(lines[i-1]) == string(lines[i]) {
			numberOfBeforeReflection := i
			for k := 2; i-k >= 0 && i+k-1 < len(lines); k++ {
				if string(lines[i-k]) != string(lines[i+k-1]) {
					continue loop
				}
			}
			return numberOfBeforeReflection, nil
		}
	}
	return 0, errors.New("reflection not found")
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

func Part2GetSum(scanner *bufio.Scanner) int {
	mirrors := []Mirror{}
	mirror := Mirror{}
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			mirrors = append(mirrors, mirror)
			mirror = Mirror{}
		} else {
			line := []rune(s)
			mirror.lines = append(mirror.lines, line)
		}
	}
	mirrors = append(mirrors, mirror)

	sum := 0
	for _, mirror := range mirrors {
		numberOfRowsBeforeReflection, numberOfColumnsBeforeReflection := Part2GetRowColBeforeReflection(mirror)
		sum += numberOfColumnsBeforeReflection + 100*numberOfRowsBeforeReflection
	}
	return sum
}

func Part2GetRowColBeforeReflection(mirror Mirror) (int, int) {
	numberOfRowsBeforeReflection, err := findReflection2(mirror.lines)
	if err == nil {
		return numberOfRowsBeforeReflection, 0
	}
	numberOfColumnsBeforeReflection, err := findReflection2(transpose(mirror.lines))
	if err == nil {
		return 0, numberOfColumnsBeforeReflection
	}
	for _, line := range mirror.lines {
		fmt.Println(string(line))
	}
	panic("reflection not found")
}

func findReflection2(lines [][]rune) (int, error) {
loop:
	for i := 1; i < len(lines); i++ {
		swapFound := false
		if compareLinesWithOneSwap(lines[i-1], lines[i], &swapFound) {
			numberOfBeforeReflection := i
			for k := 2; i-k >= 0 && i+k-1 < len(lines); k++ {
				if !compareLinesWithOneSwap(lines[i-k], lines[i+k-1], &swapFound) {
					continue loop
				}
			}
			if swapFound {
				return numberOfBeforeReflection, nil
			}
		}
	}
	return 0, errors.New("reflection not found")
}

func compareLinesWithOneSwap(line1 []rune, line2 []rune, swapFound *bool) bool {
	for i := 0; i < len(line1); i++ {
		if line1[i] != line2[i] {
			if *swapFound {
				return false
			} else {
				*swapFound = true
			}
		}
	}
	return true
}

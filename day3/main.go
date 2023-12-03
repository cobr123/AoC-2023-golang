package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
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

	arr := Part1GetLines(scanner)
	sum := Part1GetSum(arr)
	fmt.Println(sum)
}

func Part1GetLines(scanner *bufio.Scanner) [][]rune {
	arr := [][]rune{}
	for scanner.Scan() {
		s := []rune(scanner.Text())
		arr = append(arr, s)
	}
	return arr
}

func Part1GetSum(arr [][]rune) int {
	sum := 0
	for j, line := range arr {
		number := ""
		charFound := false
		for k, ch := range line {
			if unicode.Is(unicode.Digit, ch) {
				number += string(ch)
				if !charFound {
					charFound = isCharAround(arr, j, k)
				}
			} else {
				if number != "" && charFound {
					n, err := strconv.Atoi(number)
					if err != nil {
						panic(err)
					}
					sum += n
				}
				number = ""
				charFound = false
			}
		}
		if number != "" && charFound {
			n, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			sum += n
		}
	}
	return sum
}

func isCharAround(arr [][]rune, line int, pos int) bool {
	chars := []rune{}
	if line > 0 {
		if pos > 0 {
			chars = append(chars, arr[line-1][pos-1])
		}
		chars = append(chars, arr[line-1][pos])
		if pos < len(arr[line])-1 {
			chars = append(chars, arr[line-1][pos+1])
		}
	}
	if pos > 0 {
		chars = append(chars, arr[line][pos-1])
	}
	if pos < len(arr[line])-1 {
		chars = append(chars, arr[line][pos+1])
	}
	if line < len(arr)-1 {
		if pos > 0 {
			chars = append(chars, arr[line+1][pos-1])
		}
		chars = append(chars, arr[line+1][pos])
		if pos < len(arr[line])-1 {
			chars = append(chars, arr[line+1][pos+1])
		}
	}
	for _, ch := range chars {
		if ch != '.' && !unicode.Is(unicode.Digit, ch) {
			return true
		}
	}
	return false
}

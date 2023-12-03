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
	part2()
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

func part2() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	arr := Part1GetLines(scanner)
	sum := Part2GetSum(arr)
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

func Part2GetSum(arr [][]rune) int {
	sum := 0
	m := map[string]int{}
	for j, line := range arr {
		number := ""
		key := ""
		charFound := false
		for k, ch := range line {
			if unicode.Is(unicode.Digit, ch) {
				number += string(ch)
				if !charFound {
					tmp := ""
					charFound, tmp = isStarAround(arr, j, k)
					if charFound {
						key = tmp
					}
				}
			} else {
				if number != "" && charFound {
					n, err := strconv.Atoi(number)
					if err != nil {
						panic(err)
					}
					if prev, found := m[key]; found {
						delete(m, key)
						sum += prev * n
					} else {
						m[key] = n
					}
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
			if prev, found := m[key]; found {
				delete(m, key)
				sum += prev * n
			} else {
				m[key] = n
			}
		}
	}
	return sum
}

type Char struct {
	line int
	pos  int
}

func isCharAround(arr [][]rune, line int, pos int) bool {
	chars := []Char{}
	if line > 0 {
		if pos > 0 {
			chars = append(chars, Char{line - 1, pos - 1})
		}
		chars = append(chars, Char{line - 1, pos})
		if pos < len(arr[line])-1 {
			chars = append(chars, Char{line - 1, pos + 1})
		}
	}
	if pos > 0 {
		chars = append(chars, Char{line, pos - 1})
	}
	if pos < len(arr[line])-1 {
		chars = append(chars, Char{line, pos + 1})
	}
	if line < len(arr)-1 {
		if pos > 0 {
			chars = append(chars, Char{line + 1, pos - 1})
		}
		chars = append(chars, Char{line + 1, pos})
		if pos < len(arr[line])-1 {
			chars = append(chars, Char{line + 1, pos + 1})
		}
	}
	for _, char := range chars {
		ch := arr[char.line][char.pos]
		if ch != '.' && !unicode.Is(unicode.Digit, ch) {
			return true
		}
	}
	return false
}

func isStarAround(arr [][]rune, line int, pos int) (bool, string) {
	chars := []Char{}
	if line > 0 {
		if pos > 0 {
			chars = append(chars, Char{line - 1, pos - 1})
		}
		chars = append(chars, Char{line - 1, pos})
		if pos < len(arr[line])-1 {
			chars = append(chars, Char{line - 1, pos + 1})
		}
	}
	if pos > 0 {
		chars = append(chars, Char{line, pos - 1})
	}
	if pos < len(arr[line])-1 {
		chars = append(chars, Char{line, pos + 1})
	}
	if line < len(arr)-1 {
		if pos > 0 {
			chars = append(chars, Char{line + 1, pos - 1})
		}
		chars = append(chars, Char{line + 1, pos})
		if pos < len(arr[line])-1 {
			chars = append(chars, Char{line + 1, pos + 1})
		}
	}
	for _, char := range chars {
		ch := arr[char.line][char.pos]
		if ch == '*' {
			return true, strconv.Itoa(char.line) + ":" + strconv.Itoa(char.pos)
		}
	}
	return false, ""
}

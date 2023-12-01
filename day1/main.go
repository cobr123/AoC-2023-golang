package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("./input1.txt")
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
		first, last := Part1ParseDigits(s)
		n, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		if err != nil {
			panic(err)
		}
		sum += n
	}
	return sum
}

func Part1ParseDigits(s string) (int, int) {
	first := 0
	last := 0
	for _, ch := range s {
		if unicode.Is(unicode.Digit, ch) {
			i, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}
			if first == 0 {
				first = i
			}
			last = i
		}
	}
	return first, last
}

func part2() {
	f, err := os.Open("./input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := Part2GetSum(scanner)
	fmt.Println(sum)
}

func Part2GetSum(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		first, last := Part2ParseDigits(s)
		n, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		if err != nil {
			panic(err)
		}
		sum += n
	}
	return sum
}

func Part2ParseDigits(s string) (int, int) {
	digits := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	first := 0
	last := 0
	acc := ""
	for _, ch := range s {
		if unicode.Is(unicode.Digit, ch) {
			i, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}
			if first == 0 {
				first = i
			}
			last = i
			acc = ""
		} else {
			acc += string(ch)
			if v, exists := digits[acc]; exists {
				if first == 0 {
					first = v
				}
				last = v
				acc = string(ch)
			} else {
				for len(acc) > 0 {
					foundPrefix := false
					for key := range digits {
						if strings.HasPrefix(key, acc) {
							foundPrefix = true
							break
						}
					}
					if !foundPrefix {
						_, i := utf8.DecodeRuneInString(acc)
						acc = acc[i:]
					} else {
						break
					}
				}
			}
		}
	}
	return first, last
}

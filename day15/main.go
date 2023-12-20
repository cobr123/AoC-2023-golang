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

	sum := Part1GetSum(scanner)
	fmt.Println(sum)
}

func Part1GetSum(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, ",")
		for _, part := range parts {
			sum += Part1Hash(part)
		}
	}
	return sum
}

func Part1Hash(s string) int {
	result := 0
	for _, ch := range s {
		result += int(ch)
		result *= 17
		result %= 256
	}
	return result
}

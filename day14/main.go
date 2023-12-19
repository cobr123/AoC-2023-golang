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
	dish := Part1ParseDish(scanner)
	Part1SlideNorth(dish)
	return Part1GetTotalLoad(dish)
}

func Part1ParseDish(scanner *bufio.Scanner) [][]rune {
	dish := [][]rune{}
	for scanner.Scan() {
		s := scanner.Text()
		line := []rune(s)
		dish = append(dish, line)
	}
	return dish
}

func Part1SlideNorth(dish [][]rune) {
	isFound := true
	for isFound {
		isFound = false
		for row := 1; row < len(dish); row++ {
			for col := 0; col < len(dish[0]); col++ {
				if dish[row-1][col] == '.' && dish[row][col] == 'O' {
					dish[row-1][col] = 'O'
					dish[row][col] = '.'
					isFound = true
				}
			}
		}
	}
}

func Part1GetTotalLoad(dish [][]rune) int {
	sum := 0
	coef := len(dish)
	for row := 0; row < len(dish); row++ {
		for col := 0; col < len(dish[0]); col++ {
			if dish[row][col] == 'O' {
				sum += coef
			}
		}
		coef--
	}
	return sum
}

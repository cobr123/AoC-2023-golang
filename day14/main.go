package main

import (
	"bufio"
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

func Part1GetSum(scanner *bufio.Scanner) int {
	dish := Part1ParseDish(scanner)
	Part1SlideNorth(dish)
	return Part1GetTotalLoad(dish)
}

func Part2GetSum(scanner *bufio.Scanner) int {
	dish := Part1ParseDish(scanner)
	for i := 0; i < 1e3; i++ {
		Part2SpinCycle(dish)
	}
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

func Part2SpinCycle(dish [][]rune) {
	Part1SlideNorth(dish)
	Part2SlideWest(dish)
	Part2SlideSouth(dish)
	Part2SlideEast(dish)
}

func Part1SlideNorth(dish [][]rune) {
	isFound := true
	for isFound {
		isFound = false
		for row := 1; row < len(dish); row++ {
			for col := 0; col < len(dish[0]); col++ {
				if dish[row][col] == 'O' && dish[row-1][col] == '.' {
					dish[row-1][col] = 'O'
					dish[row][col] = '.'
					isFound = true
				}
			}
		}
	}
}

func Part2SlideWest(dish [][]rune) {
	isFound := true
	for isFound {
		isFound = false
		for row := 0; row < len(dish); row++ {
			for col := 1; col < len(dish[0]); col++ {
				if dish[row][col] == 'O' && dish[row][col-1] == '.' {
					dish[row][col-1] = 'O'
					dish[row][col] = '.'
					isFound = true
				}
			}
		}
	}
}

func Part2SlideSouth(dish [][]rune) {
	isFound := true
	for isFound {
		isFound = false
		for row := 0; row < len(dish)-1; row++ {
			for col := 0; col < len(dish[0]); col++ {
				if dish[row][col] == 'O' && dish[row+1][col] == '.' {
					dish[row+1][col] = 'O'
					dish[row][col] = '.'
					isFound = true
				}
			}
		}
	}
}

func Part2SlideEast(dish [][]rune) {
	isFound := true
	for isFound {
		isFound = false
		for row := 0; row < len(dish); row++ {
			for col := 0; col < len(dish[0])-1; col++ {
				if dish[row][col] == 'O' && dish[row][col+1] == '.' {
					dish[row][col+1] = 'O'
					dish[row][col] = '.'
					isFound = true
				}
			}
		}
	}
}

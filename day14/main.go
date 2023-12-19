package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
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

var cacheIdxToDish = map[uint8]string{}
var cacheDishToIdx = map[string]uint8{}
var cacheDishIdx uint8 = 0

func cacheDish(dish [][]rune) uint8 {
	dishStr := str(dish)
	if idx, ok := cacheDishToIdx[dishStr]; ok {
		return idx
	} else {
		idx = cacheDishIdx
		cacheIdxToDish[idx] = dishStr
		cacheDishToIdx[dishStr] = idx
		cacheDishIdx++
		return idx
	}
}

func Part2GetSum(scanner *bufio.Scanner) int {
	dish := Part1ParseDish(scanner)
	cacheFromTo := [255]uint8{}

	key := cacheDish(dish)

	for i := 0; i < 1e9; i++ {
		if value := cacheFromTo[key]; value > 0 {
			key = value
		} else {
			dish = Part1ParseDish(bufio.NewScanner(bytes.NewReader([]byte(cacheIdxToDish[key]))))
			Part2SpinCycle(dish)
			value = cacheDish(dish)
			cacheFromTo[key] = value
			key = value
		}
	}
	dish = Part1ParseDish(bufio.NewScanner(bytes.NewReader([]byte(cacheIdxToDish[key]))))
	return Part1GetTotalLoad(dish)
}

func str(dish [][]rune) string {
	s := ""
	for _, line := range dish {
		s += string(line)
		s += "\n"
	}
	return strings.Trim(s, "\n")
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

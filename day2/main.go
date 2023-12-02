package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

type Cards struct {
	red   int
	green int
	blue  int
}

func part1() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	cardsMaxValues := Cards{red: 12, green: 13, blue: 14}
	sum := Part1GetSum(scanner, cardsMaxValues)
	fmt.Println(sum)
}

func Part1GetSum(scanner *bufio.Scanner, cardsMaxValues Cards) int {
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		sum += Part1ParseGame(s, cardsMaxValues)
	}
	return sum
}

func Part1ParseGame(str string, cardsMaxValues Cards) int {
	gameAndSteps := strings.Split(str, ":")
	gameIdArr := strings.Split(strings.Trim(gameAndSteps[0], " "), " ")
	gameId, err := strconv.Atoi(gameIdArr[1])
	if err != nil {
		panic(err)
	}
	acc := Cards{}
	steps := strings.Split(strings.Trim(gameAndSteps[1], " "), ";")
	for _, step := range steps {
		cards := strings.Split(strings.Trim(step, " "), ",")
		for _, card := range cards {
			cntAndCardName := strings.Split(strings.Trim(card, " "), " ")
			name := strings.Trim(cntAndCardName[1], " ")
			cnt, err := strconv.Atoi(cntAndCardName[0])
			if err != nil {
				panic(err)
			}
			switch name {
			case "blue":
				acc.blue = int(math.Max(float64(acc.blue), float64(cnt)))
				if acc.blue > cardsMaxValues.blue {
					return 0
				}
			case "green":
				acc.green = int(math.Max(float64(acc.green), float64(cnt)))
				if acc.green > cardsMaxValues.green {
					return 0
				}
			case "red":
				acc.red = int(math.Max(float64(acc.red), float64(cnt)))
				if acc.red > cardsMaxValues.red {
					return 0
				}
			}
		}
	}
	return gameId
}

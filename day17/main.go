package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
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

	sum := Part1FindPath(scanner)
	fmt.Println(sum)
}

func Part1FindPath(scanner *bufio.Scanner) int {
	tiles := Part1GetTiles(scanner)
	maxSteps := len(tiles) * len(tiles[0])
	return Part1FindPathWithMinHeatLoss(tiles, Step{Pos{0, 0}, 'r', 3, maxSteps, 0, []Pos{}})
}

type Pos struct {
	row int
	col int
}

type Step struct {
	pos             Pos
	direction       rune
	stepsBeforeTurn int
	stepsLeft       int
	heatLoss        int
	visited         []Pos
}

func Part1FindPathWithMinHeatLoss(tiles [][]int, initStep Step) int {
	var minLoss int = 1e9
	endRow := len(tiles) - 1
	endCol := len(tiles[endRow]) - 1

	steps := []Step{}
	steps = append(steps, initStep)

	for len(steps) > 0 {
		newSteps := []Step{}
		for _, step := range steps {
			if step.pos.row == endRow && step.pos.col == endCol {
				minLoss = step.heatLoss
				//if step.heatLoss < 115 {
				//Part1PrintVisited(visited, heatLoss)
				Part1PrintVisited(tiles, step)
				//}
			} else {
				isInTheBox := step.pos.row >= 0 && step.pos.row < len(tiles) && step.pos.col >= 0 && step.pos.col < len(tiles[0])
				if isInTheBox && step.stepsLeft > 0 && !slices.Contains(step.visited, step.pos) {
					newDirections := Part1GetNewDirections(step.stepsBeforeTurn, step.direction)

					for _, newDirection := range newDirections {
						newVisited := append(step.visited, step.pos)
						row, col := Part1GetNewPosByDirection(step, newDirection)
						newSteps = append(newSteps, Step{
							Pos{row, col},
							newDirection,
							step.stepsBeforeTurn - 1,
							step.stepsLeft - 1,
							step.heatLoss + tiles[step.pos.row][step.pos.col],
							newVisited},
						)
					}
				}
			}
		}
		slices.SortFunc(newSteps, func(a, b Step) int {
			return a.heatLoss - b.heatLoss
		})
		Part1PrintVisited(tiles, newSteps[0])
		fmt.Println(len(newSteps))
		fmt.Println("-------------------")
		steps = newSteps
	}
	return minLoss
}

func Part1PrintVisited(tiles [][]int, step Step) {
	newTiles := Part1CopyTiles(tiles)
	for _, pos := range step.visited {
		newTiles[pos.row][pos.col] = 0
	}
	for _, line := range newTiles {
		for _, ch := range line {
			fmt.Print(strconv.Itoa(ch))
		}
		fmt.Println("")
	}
	fmt.Println("-------------------", step.heatLoss)
}

func Part1CopyTiles(tiles [][]int) [][]int {
	newTiles := make([][]int, len(tiles))
	for i := 0; i < len(newTiles); i++ {
		newTiles[i] = make([]int, len(tiles[i]))
		for k := 0; k < len(newTiles[i]); k++ {
			newTiles[i][k] = tiles[i][k]
		}
	}
	return newTiles
}

func Part1GetNewDirections(stepsBeforeTurn int, direction rune) []rune {
	newDirections := []rune{}
	if stepsBeforeTurn == 1 {
		stepsBeforeTurn = 3
		switch direction {
		case 't', 'd':
			newDirections = append(newDirections, 'r', 'l')
		case 'l', 'r':
			newDirections = append(newDirections, 'd', 't')
		}
	} else {
		switch direction {
		case 't':
			newDirections = append(newDirections, 't', 'l', 'r')
		case 'd':
			newDirections = append(newDirections, 'd', 'l', 'r')
		case 'l':
			newDirections = append(newDirections, 'l', 't', 'd')
		case 'r':
			newDirections = append(newDirections, 'r', 't', 'd')
		}
	}
	rand.Shuffle(len(newDirections), func(i, j int) {
		newDirections[i], newDirections[j] = newDirections[j], newDirections[i]
	})
	return newDirections
}

func Part1GetNewPosByDirection(step Step, direction rune) (int, int) {
	switch direction {
	case 't':
		return step.pos.row - 1, step.pos.col
	case 'd':
		return step.pos.row + 1, step.pos.col
	case 'l':
		return step.pos.row, step.pos.col - 1
	case 'r':
		return step.pos.row, step.pos.col + 1
	}
	panic("unknown direction: " + string(direction))
}

func Part1GetTiles(scanner *bufio.Scanner) [][]int {
	tiles := [][]int{}
	for scanner.Scan() {
		s := scanner.Text()
		numbers := []int{}
		notVisited := []rune{}
		for _, ch := range s {
			number, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, number)
			notVisited = append(notVisited, '.')
		}
		tiles = append(tiles, numbers)
	}
	return tiles
}

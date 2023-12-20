package main

import (
	"bufio"
	"fmt"
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

func Part1FindPathWithMinHeatLoss(tiles [][]int, step Step) int {
	var loss int = 1e9
	endRow := len(tiles) - 1
	endCol := len(tiles[endRow]) - 1

	newDirections, stepsBeforeTurn := Part1GetNewDirections(step.stepsBeforeTurn-1, step.direction)
	for _, newDirection := range newDirections {
		newVisited := append(step.visited, step.pos)
		row, col := Part1GetNewPosByDirection(step, newDirection)
		isInTheBox := row >= 0 && row < len(tiles) && col >= 0 && col < len(tiles[0])
		pos := Pos{row, col}
		newStep := Step{
			pos,
			newDirection,
			stepsBeforeTurn,
			step.stepsLeft - 1,
			step.heatLoss + tiles[step.pos.row][step.pos.col],
			newVisited}
		if row == endRow && col == endCol {
			newStep.heatLoss += tiles[row][col]
			newStep.visited = append(newStep.visited, pos)
			if newStep.heatLoss < loss {
				loss = newStep.heatLoss
				if loss < 400 {
					Part1PrintVisited(tiles, newStep)
				}
			}
		} else if isInTheBox && newStep.stepsLeft > 0 && !slices.Contains(step.visited, pos) {
			if tmp := Part1FindPathWithMinHeatLoss(tiles, newStep); tmp < loss {
				loss = tmp
			}
		}
	}
	return loss
}

func Part1PrintVisited(tiles [][]int, step Step) {
	newTiles := Part1CopyTiles(tiles)
	loss := 0
	for _, pos := range step.visited {
		loss += newTiles[pos.row][pos.col]
		newTiles[pos.row][pos.col] = 0
	}
	for _, line := range newTiles {
		for _, ch := range line {
			fmt.Print(strconv.Itoa(ch))
		}
		fmt.Println("")
	}
	fmt.Println("-------------------", step.heatLoss, loss)
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

func Part1GetNewDirections(stepsBeforeTurn int, direction rune) ([]rune, int) {
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
	//rand.Shuffle(len(newDirections), func(i, j int) {
	//	newDirections[i], newDirections[j] = newDirections[j], newDirections[i]
	//})
	return newDirections, stepsBeforeTurn
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

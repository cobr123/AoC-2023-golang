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
	maxSteps := 35 //len(tiles) * len(tiles[0])
	return Part1FindPathWithMinHeatLoss(tiles, Step{Pos{0, 0}, 'r', 3, maxSteps, 0, []Pos{}})
}

type Pos struct {
	row int
	col int
}

func (p *Pos) isInTheBox(tiles [][]int) bool {
	return p.row >= 0 && p.row < len(tiles) && p.col >= 0 && p.col < len(tiles[0])
}

type Step struct {
	pos             Pos
	direction       rune
	stepsBeforeTurn int
	stepsLeft       int
	heatLoss        int
	visited         []Pos
}

var loss int = 1e9

func Part1FindPathWithMinHeatLoss(tiles [][]int, step Step) int {
	endRow := len(tiles) - 1
	endCol := len(tiles[endRow]) - 1

	newDirections, stepsBeforeTurn := Part1GetNewDirections(tiles, step.pos, step.stepsBeforeTurn, step.direction)
	for _, newDirection := range newDirections {
		newVisited := append(step.visited, step.pos)
		pos := Part1GetNewPosByDirection(step.pos, newDirection)
		newStep := Step{
			pos,
			newDirection,
			stepsBeforeTurn,
			step.stepsLeft - 1,
			step.heatLoss + tiles[step.pos.row][step.pos.col],
			newVisited}
		if pos.row == endRow && pos.col == endCol {
			newStep.heatLoss += tiles[pos.row][pos.col]
			newStep.visited = append(newStep.visited, pos)
			if newStep.heatLoss < loss {
				loss = newStep.heatLoss
				if loss < 150 {
					Part1PrintVisited(tiles, newStep)
				}
			}
		} else if pos.isInTheBox(tiles) && newStep.stepsLeft > 0 && !slices.Contains(newStep.visited, pos) {
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

func Part1GetNewDirections(tiles [][]int, pos Pos, stepsBeforeTurn int, direction rune) ([]rune, int) {
	newDirections := []rune{}
	stepsBeforeTurn = stepsBeforeTurn - 1
	if stepsBeforeTurn == 1 {
		stepsBeforeTurn = 3
		switch direction {
		case 't', 'd':
			newDirections = append(newDirections, 'l', 'r')
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
	tmpDirections := []rune{}
	for _, d := range newDirections {
		if newPos := Part1GetNewPosByDirection(pos, d); newPos.isInTheBox(tiles) {
			tmpDirections = append(tmpDirections, d)
		}
	}
	newDirections = tmpDirections
	slices.SortFunc(newDirections, func(a, b rune) int {
		posA := Part1GetNewPosByDirection(pos, a)
		posB := Part1GetNewPosByDirection(pos, b)
		return tiles[posA.row][posA.col] - tiles[posB.row][posB.col]
	})
	//slices.SortFunc(newDirections, func(a, b rune) int {
	//	posA := Part1GetNewPosByDirection(pos, a)
	//	posB := Part1GetNewPosByDirection(pos, b)
	//	if posA.row == posB.row {
	//		return posB.col - posA.col
	//	} else {
	//		return posB.row - posA.row
	//	}
	//})

	//rand.Shuffle(len(newDirections), func(i, j int) {
	//	newDirections[i], newDirections[j] = newDirections[j], newDirections[i]
	//})
	return newDirections, stepsBeforeTurn
}

func Part1GetNewPosByDirection(pos Pos, direction rune) Pos {
	switch direction {
	case 't':
		return Pos{pos.row - 1, pos.col}
	case 'd':
		return Pos{pos.row + 1, pos.col}
	case 'l':
		return Pos{pos.row, pos.col - 1}
	case 'r':
		return Pos{pos.row, pos.col + 1}
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

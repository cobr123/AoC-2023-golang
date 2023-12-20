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

type Turn struct {
	pos       Pos
	direction rune
	loss      int
}

var loss int = 1e9

func Part1FindPathWithMinHeatLoss(tiles [][]int, step Step) int {
	newTurns, stepsBeforeTurn := Part1GetNewTurns(tiles, step)
	for _, newTurn := range newTurns {
		newVisited := append(step.visited, step.pos)
		newStep := Step{
			newTurn.pos,
			newTurn.direction,
			stepsBeforeTurn,
			step.stepsLeft - 1,
			step.heatLoss + tiles[step.pos.row][step.pos.col],
			newVisited}
		if newTurn.pos.row == len(tiles)-1 && newTurn.pos.col == len(tiles[0])-1 {
			newStep.heatLoss += tiles[newTurn.pos.row][newTurn.pos.col]
			newStep.visited = append(newStep.visited, newTurn.pos)
			if newStep.heatLoss < loss {
				loss = newStep.heatLoss
				if loss < 150 {
					Part1PrintVisited(tiles, newStep)
				}
			}
		} else {
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

func Part1GetNewDirections(direction rune, stepsBeforeTurn int) ([]rune, int) {
	newDirections := []rune{}
	stepsBeforeTurn = stepsBeforeTurn - 1
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
			newDirections = append(newDirections, 'r', 't', 'l')
		case 'd':
			newDirections = append(newDirections, 'r', 'd', 'l')
		case 'l':
			newDirections = append(newDirections, 'd', 'l', 't')
		case 'r':
			newDirections = append(newDirections, 'r', 'd', 't')
		}
	}
	return newDirections, stepsBeforeTurn
}

func Part1GetNewTurns(tiles [][]int, step Step) ([]Turn, int) {
	newDirections, stepsBeforeTurn := Part1GetNewDirections(step.direction, step.stepsBeforeTurn)
	newTurns := []Turn{}
	for _, d := range newDirections {
		if newPos := Part1GetNewPosByDirection(step.pos, d); newPos.isInTheBox(tiles) && step.stepsLeft > 0 && !slices.Contains(step.visited, newPos) {
			newTurns = append(newTurns, Turn{newPos, d, tiles[newPos.row][newPos.col]})
		}
	}
	//slices.SortFunc(newTurns, func(a, b Turn) int {
	//	return a.loss - b.loss
	//})
	return newTurns, stepsBeforeTurn
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

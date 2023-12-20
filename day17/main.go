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
	return Part1FindPathWithMinHeatLoss(tiles, Step{Pos{0, 0}, 'r', 3, -tiles[0][0], []Turn{}})
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
	heatLoss        int
	visited         []Turn
}

type Turn struct {
	pos             Pos
	direction       rune
	stepsBeforeTurn int
}

var loss int = 1e9

func Part1FindPathWithMinHeatLoss(tiles [][]int, step Step) int {
	newTurns := Part1GetNewTurns(tiles, step)
	for _, newTurn := range newTurns {
		newVisited := append(step.visited, Turn{step.pos, step.direction, step.stepsBeforeTurn})
		newStep := Step{
			newTurn.pos,
			newTurn.direction,
			newTurn.stepsBeforeTurn,
			step.heatLoss + tiles[step.pos.row][step.pos.col],
			newVisited}
		if newTurn.pos.row == len(tiles)-1 && newTurn.pos.col == len(tiles[0])-1 {
			newStep.heatLoss += tiles[newTurn.pos.row][newTurn.pos.col]
			newStep.visited = append(newStep.visited, newTurn)
			if newStep.heatLoss < loss {
				loss = newStep.heatLoss
				Part1PrintVisited(tiles, newStep)
			}
		} else if newStep.heatLoss < loss {
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
	for i := 1; i < len(step.visited); i++ {
		pos := step.visited[i].pos
		loss += tiles[pos.row][pos.col]
		newTiles[pos.row][pos.col] = step.visited[i].direction
	}
	for _, line := range newTiles {
		for _, ch := range line {
			fmt.Print(string(ch))
		}
		fmt.Println("")
	}
	fmt.Println("-------------------", step.heatLoss, loss)
}

func Part1CopyTiles(tiles [][]int) [][]rune {
	newTiles := make([][]rune, len(tiles))
	for i := 0; i < len(newTiles); i++ {
		newTiles[i] = make([]rune, len(tiles[i]))
		for k := 0; k < len(newTiles[i]); k++ {
			newTiles[i][k] = '.'
		}
	}
	return newTiles
}

func Part1GetNewDirections(direction rune) []rune {
	newDirections := []rune{}
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
	return newDirections
}

func Part1GetNewTurns(tiles [][]int, step Step) []Turn {
	newDirections := Part1GetNewDirections(step.direction)
	newTurns := []Turn{}
	for _, d := range newDirections {
		if newPos := Part1GetNewPosByDirection(step.pos, d); newPos.isInTheBox(tiles) && !slices.ContainsFunc(step.visited, func(turn Turn) bool {
			return turn.pos == newPos
		}) {
			stepsBeforeTurn := step.stepsBeforeTurn
			if step.direction != d {
				stepsBeforeTurn = 3
			} else {
				stepsBeforeTurn--
			}
			if stepsBeforeTurn > 0 {
				newTurns = append(newTurns, Turn{newPos, d, stepsBeforeTurn})
			}
		}
	}
	return newTurns
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

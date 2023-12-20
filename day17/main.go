package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
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
	tiles, visited := Part1GetTiles(scanner)
	maxSteps := len(tiles) * len(tiles[0])
	return Part1FindPathWithMinHeatLoss(tiles, visited, 'r', 3, maxSteps, 0, Step{0, 0})
}

type Step struct {
	row int
	col int
}

func Part1FindPathWithMinHeatLoss(tiles [][]int, visited [][]rune, direction rune, stepsBeforeTurn int, stepsLeft int, heatLoss int, step Step) int {
	var minLoss int = 1e9
	endRow := len(tiles) - 1
	endCol := len(tiles[endRow]) - 1

	if step.row == endRow && step.col == endCol {
		minLoss = heatLoss
		if heatLoss < 110 {
			Part1PrintVisited(visited, heatLoss)
		}
	} else {
		isInTheBox := step.row >= 0 && step.row < len(tiles) && step.col >= 0 && step.col < len(tiles[0])
		if isInTheBox && stepsLeft > 0 && visited[step.row][step.col] == '.' {
			newDirections := Part1GetNewDirections(stepsBeforeTurn, direction)
			n := len(newDirections)
			losses := make([]int, n, n)
			wg := new(sync.WaitGroup)
			wg.Add(n)
			for i, newDirection := range newDirections {
				go func(idx int, wg *sync.WaitGroup, arr []int, newDirection rune, visited [][]rune) {
					newVisited := Part1CopyVisited(visited)
					newVisited[step.row][step.col] = newDirection
					row, col := Part1GetNewPosByDirection(step, newDirection)
					arr[idx] = Part1FindPathWithMinHeatLoss(tiles, newVisited, newDirection, stepsBeforeTurn-1, stepsLeft-1, heatLoss+tiles[step.row][step.col], Step{row, col})
				}(i, wg, losses, newDirection, visited)
			}
			wg.Wait()
			for _, loss := range losses {
				if loss < minLoss {
					minLoss = loss
				}
			}
		}
	}
	return minLoss
}

func Part1CopyVisited(visited [][]rune) [][]rune {
	newVisited := make([][]rune, len(visited))
	for i := 0; i < len(newVisited); i++ {
		newVisited[i] = make([]rune, len(visited[i]))
		for k := 0; k < len(newVisited[i]); k++ {
			newVisited[i][k] = visited[i][k]
		}
	}
	return newVisited
}

func Part1PrintVisited(visited [][]rune, heatLoss int) {
	for _, line := range visited {
		for _, ch := range line {
			switch ch {
			case 't':
				fmt.Print("^")
			case 'd':
				fmt.Print("v")
			case 'l':
				fmt.Print("<")
			case 'r':
				fmt.Print(">")
			default:
				fmt.Print(string(ch))
			}
		}
		fmt.Println("")
	}
	fmt.Println("-------------------", heatLoss)
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
		return step.row - 1, step.col
	case 'd':
		return step.row + 1, step.col
	case 'l':
		return step.row, step.col - 1
	case 'r':
		return step.row, step.col + 1
	}
	panic("unknown direction: " + string(direction))
}

func Part1GetTiles(scanner *bufio.Scanner) ([][]int, [][]rune) {
	tiles := [][]int{}
	visited := [][]rune{}
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
		visited = append(visited, notVisited)
	}
	return tiles, visited
}

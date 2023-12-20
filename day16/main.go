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

	sum := Part1CountEnergized(scanner)
	fmt.Println(sum)
}

func Part1CountEnergized(scanner *bufio.Scanner) int {
	tiles := [][]rune{}
	for scanner.Scan() {
		s := scanner.Text()
		line := []rune(s)
		tiles = append(tiles, line)
	}
	energized := Part1GetEnergized(tiles)
	cnt := 0
	for _, line := range energized {
		for _, power := range line {
			if power > 0 {
				//fmt.Print("#")
				cnt++
			} else {
				//fmt.Print(".")
			}
		}
		//fmt.Println("")
	}
	return cnt
}

type Step struct {
	row       int
	col       int
	direction rune
}

func Part1GetEnergized(tiles [][]rune) [][]int {
	energized := [][]int{}
	for _, line := range tiles {
		energized = append(energized, make([]int, len(line)))
	}
	steps := []Step{}
	steps = append(steps, Step{0, 0, 'r'})

	for len(steps) > 0 {
		newSteps := []Step{}
		for _, step := range steps {
			isInTheBox := step.row >= 0 && step.row < len(tiles) && step.col >= 0 && step.col < len(tiles[0])
			if isInTheBox && energized[step.row][step.col] < 10000 {
				energized[step.row][step.col]++

				switch tiles[step.row][step.col] {
				case '.':
					newStep := Part1GetNewPosByDirection(step)
					newSteps = append(newSteps, newStep)
				case '|':
					switch step.direction {
					case 't':
						newStep := Part1GetNewPosByDirection(step)
						newSteps = append(newSteps, newStep)
					case 'd':
						newStep := Part1GetNewPosByDirection(step)
						newSteps = append(newSteps, newStep)
					case 'l':
						newStep1 := Part1GetNewPosByDirection(Step{step.row, step.col, 't'})
						newSteps = append(newSteps, newStep1)
						newStep2 := Part1GetNewPosByDirection(Step{step.row, step.col, 'd'})
						newSteps = append(newSteps, newStep2)
					case 'r':
						newStep1 := Part1GetNewPosByDirection(Step{step.row, step.col, 't'})
						newSteps = append(newSteps, newStep1)
						newStep2 := Part1GetNewPosByDirection(Step{step.row, step.col, 'd'})
						newSteps = append(newSteps, newStep2)
					}
				case '-':
					switch step.direction {
					case 't':
						newStep1 := Part1GetNewPosByDirection(Step{step.row, step.col, 'l'})
						newSteps = append(newSteps, newStep1)
						newStep2 := Part1GetNewPosByDirection(Step{step.row, step.col, 'r'})
						newSteps = append(newSteps, newStep2)
					case 'd':
						newStep1 := Part1GetNewPosByDirection(Step{step.row, step.col, 'l'})
						newSteps = append(newSteps, newStep1)
						newStep2 := Part1GetNewPosByDirection(Step{step.row, step.col, 'r'})
						newSteps = append(newSteps, newStep2)
					case 'l':
						newStep := Part1GetNewPosByDirection(step)
						newSteps = append(newSteps, newStep)
					case 'r':
						newStep := Part1GetNewPosByDirection(step)
						newSteps = append(newSteps, newStep)
					}
				case '/':
					switch step.direction {
					case 't':
						newStep := Part1GetNewPosByDirection(Step{step.row, step.col, 'r'})
						newSteps = append(newSteps, newStep)
					case 'd':
						newStep := Part1GetNewPosByDirection(Step{step.row, step.col, 'l'})
						newSteps = append(newSteps, newStep)
					case 'l':
						newStep := Part1GetNewPosByDirection(Step{step.row, step.col, 'd'})
						newSteps = append(newSteps, newStep)
					case 'r':
						newStep := Part1GetNewPosByDirection(Step{step.row, step.col, 't'})
						newSteps = append(newSteps, newStep)
					}
				case '\\':
					switch step.direction {
					case 't':
						newStep := Part1GetNewPosByDirection(Step{step.row, step.col, 'l'})
						newSteps = append(newSteps, newStep)
					case 'd':
						newStep := Part1GetNewPosByDirection(Step{step.row, step.col, 'r'})
						newSteps = append(newSteps, newStep)
					case 'l':
						newStep := Part1GetNewPosByDirection(Step{step.row, step.col, 't'})
						newSteps = append(newSteps, newStep)
					case 'r':
						newStep := Part1GetNewPosByDirection(Step{step.row, step.col, 'd'})
						newSteps = append(newSteps, newStep)
					}
				}
			}
		}
		steps = newSteps
	}
	return energized
}

func Part1GetNewPosByDirection(step Step) Step {
	switch step.direction {
	case 't':
		step.row = step.row - 1
	case 'd':
		step.row = step.row + 1
	case 'l':
		step.col = step.col - 1
	case 'r':
		step.col = step.col + 1
	}
	return step
}

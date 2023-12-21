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

type Direction int8

const (
	Left Direction = iota
	Right
	Up
	Down
)

type Move struct {
	direction Direction
	steps     int
}

func (m *Move) String() string {
	switch m.direction {
	case Up:
		return fmt.Sprintf("U %d", m.steps)
	case Down:
		return fmt.Sprintf("D %d", m.steps)
	case Left:
		return fmt.Sprintf("L %d", m.steps)
	case Right:
		return fmt.Sprintf("R %d", m.steps)
	}
	return ""
}

func Part1GetSum(scanner *bufio.Scanner) int {
	moves := Part1ParseMoves(scanner)
	field := Part1ParseField(moves)
	Part1PrintField(field)
	Part1DigField(field, moves)
	Part1PrintField(field)
	return 0
}

func Part1DigField(field [][]rune, moves []Move) {
	rowCurr := len(field)/2 - 1
	colCurr := len(field[0])/2 - 1
	for _, move := range moves {
		switch move.direction {
		case Up:
			rowPrev := rowCurr
			rowCurr -= move.steps
			for i := rowCurr; i <= rowPrev; i++ {
				field[i][colCurr] = '#'
			}
		case Down:
			rowPrev := rowCurr
			rowCurr += move.steps
			for i := rowPrev; i <= rowCurr; i++ {
				field[i][colCurr] = '#'
			}
		case Left:
			colPrev := colCurr
			colCurr -= move.steps
			for i := colCurr; i <= colPrev; i++ {
				field[rowCurr][i] = '#'
			}
		case Right:
			colPrev := colCurr
			colCurr += move.steps
			for i := colPrev; i <= colCurr; i++ {
				field[rowCurr][i] = '#'
			}
		}
	}
}

func Part1PrintField(field [][]rune) {
	for _, line := range field {
		for _, ch := range line {
			fmt.Print(string(ch))
		}
		fmt.Println("")
	}
	fmt.Println("---------------")
}

func Part1ParseField(moves []Move) [][]rune {
	var rowMin int = 1e9
	var rowMax int = -1e9
	var colMin int = 1e9
	var colMax int = -1e9
	colCurr := 0
	rowCurr := 0
	for _, move := range moves {
		switch move.direction {
		case Up:
			rowCurr -= move.steps
			if rowCurr < rowMin {
				rowMin = rowCurr
			}
		case Down:
			rowCurr += move.steps
			if rowCurr > rowMax {
				rowMax = rowCurr
			}
		case Left:
			colCurr -= move.steps
			if colCurr < colMin {
				colMin = colCurr
			}
		case Right:
			colCurr += move.steps
			if colCurr > colMax {
				colMax = colCurr
			}
		}
	}

	rowsCnt := int(math.Abs(float64(rowMin))+math.Abs(float64(rowMax))) * 2
	colsCnt := int(math.Abs(float64(colMin))+math.Abs(float64(colMax))) * 2
	fmt.Println(rowsCnt, colsCnt)
	field := make([][]rune, rowsCnt)
	for r := 0; r < rowsCnt; r++ {
		line := make([]rune, colsCnt)
		for c := 0; c < colsCnt; c++ {
			line[c] = '.'
		}
		field[r] = line
	}
	return field
}

func Part1ParseMoves(scanner *bufio.Scanner) []Move {
	moves := []Move{}
	for scanner.Scan() {
		s := scanner.Text()
		directionStepsAndColor := strings.Split(s, " ")
		move := Move{}
		switch directionStepsAndColor[0] {
		case "U":
			move.direction = Up
		case "D":
			move.direction = Down
		case "L":
			move.direction = Left
		case "R":
			move.direction = Right
		default:
			panic("direction not found")

		}
		steps, err := strconv.Atoi(directionStepsAndColor[1])
		if err != nil {
			panic(err)
		}
		move.steps = steps
		moves = append(moves, move)
	}
	return moves
}

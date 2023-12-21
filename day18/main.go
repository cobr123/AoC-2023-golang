package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

type Direction int8

const (
	Left Direction = iota
	Right
	Up
	Down
)

type Move struct {
	direction Direction
	steps     int64
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

func Part1GetSum(scanner *bufio.Scanner) int64 {
	moves := Part1ParseMoves(scanner)
	field := Part1ParseField(moves)
	vertices := Part1DGetVertices(field, moves)
	return getPolygonArea(vertices)
}

func Part2GetSum(scanner *bufio.Scanner) int64 {
	moves := Part2ParseMoves(scanner)
	field := Part1ParseField(moves)
	vertices := Part1DGetVertices(field, moves)
	return getPolygonArea(vertices)
}

type Field struct {
	rowsCnt int64
	colsCnt int64
}

func Part1DGetVertices(field Field, moves []Move) []Pos {
	vertices := []Pos{}
	rowCurr := field.rowsCnt/2 - 1
	colCurr := field.colsCnt/2 - 1
	for _, move := range moves {
		switch move.direction {
		case Up:
			rowPrev := rowCurr
			rowCurr -= move.steps
			for i := rowCurr; i <= rowPrev; i++ {
				vertices = append(vertices, Pos{i, colCurr})
			}
		case Down:
			rowPrev := rowCurr
			rowCurr += move.steps
			for i := rowPrev; i <= rowCurr; i++ {
				vertices = append(vertices, Pos{i, colCurr})
			}
		case Left:
			colPrev := colCurr
			colCurr -= move.steps
			for i := colCurr; i <= colPrev; i++ {
				vertices = append(vertices, Pos{rowCurr, i})
			}
		case Right:
			colPrev := colCurr
			colCurr += move.steps
			for i := colPrev; i <= colCurr; i++ {
				vertices = append(vertices, Pos{rowCurr, i})
			}
		}
	}
	return vertices
}

type Pos struct {
	row int64
	col int64
}

// A function to apply the Shoelace algorithm
func getPolygonArea(vertices []Pos) int64 {
	if len(vertices) == 0 {
		return 0
	}
	var sum1 int64 = 0
	var sum2 int64 = 0

	for i := 0; i < len(vertices)-1; i++ {
		sum1 = sum1 + vertices[i].row*vertices[i+1].col
		sum2 = sum2 + vertices[i].col*vertices[i+1].row
	}

	sum1 = sum1 + vertices[len(vertices)-1].row*vertices[0].col
	sum2 = sum2 + vertices[0].row*vertices[len(vertices)-1].col

	if sum1 > sum2 {
		return (sum1 - sum2) / 2
	} else {
		return (sum2 - sum1) / 2
	}
}

func Part1ParseField(moves []Move) Field {
	var rowMin int64 = 1e9
	var rowMax int64 = -1e9
	var colMin int64 = 1e9
	var colMax int64 = -1e9
	var colCurr int64 = 0
	var rowCurr int64 = 0
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

	var rowsCnt int64 = 0
	if rowMin < 0 {
		rowsCnt += -rowMin
	} else {
		rowsCnt += rowMin
	}
	if rowMax < 0 {
		rowsCnt += -rowMax
	} else {
		rowsCnt += rowMax
	}
	rowsCnt = rowsCnt/2 + 1

	var colsCnt int64 = 0
	if colMin < 0 {
		colsCnt += -colMin
	} else {
		colsCnt += colMin
	}
	if colMax < 0 {
		colsCnt += -colMax
	} else {
		colsCnt += colMax
	}
	colsCnt = colsCnt/2 + 1
	fmt.Println(rowMin, colMin, rowMax, colMax, rowsCnt, colsCnt)

	return Field{rowsCnt, colsCnt}
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
		steps, err := strconv.ParseInt(directionStepsAndColor[1], 10, 64)
		if err != nil {
			panic(err)
		}
		move.steps = steps
		moves = append(moves, move)
	}
	return moves
}

func Part2ParseMoves(scanner *bufio.Scanner) []Move {
	moves := []Move{}
	for scanner.Scan() {
		s := scanner.Text()
		prefixAndColor := strings.Split(s, "#")
		hex := prefixAndColor[1][:len(prefixAndColor[1])-1]
		color := hex[:len(hex)-1]
		direction := hex[len(hex)-1:]
		move := Move{}
		switch direction {
		case "3":
			move.direction = Up
		case "1":
			move.direction = Down
		case "2":
			move.direction = Left
		case "0":
			move.direction = Right
		default:
			panic("direction not found")

		}
		steps, err := strconv.ParseInt(color, 16, 64)
		if err != nil {
			panic(err)
		}
		move.steps = steps
		moves = append(moves, move)
	}
	return moves
}

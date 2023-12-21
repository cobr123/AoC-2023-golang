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
	Part1DigField(field, moves)
	Part1PrintField(field)
	Part1MarkMarkOutsideTiles(field)
	Part1PrintField(field)
	return Part1Count(field)
}

func Part2GetSum(scanner *bufio.Scanner) int {
	moves := Part2ParseMoves(scanner)
	field := Part1ParseField(moves)
	Part1DigField(field, moves)
	Part1PrintField(field)
	Part1MarkMarkOutsideTiles(field)
	Part1PrintField(field)
	return Part1Count(field)
}

func Part1Count(field [][]rune) int {
	cnt := 0
	for _, line := range field {
		for _, ch := range line {
			if ch == '#' || ch == '.' {
				cnt++
			}
		}
	}
	return cnt
}

type Pos struct {
	x int
	y int
}

func Part1MarkMarkOutsideTiles(tiles [][]rune) {
	w := len(tiles[0])
	h := len(tiles)
	for i := 0; i < w; i++ {
		MarkOutsideGround(Pos{i, 0}, tiles, w, h)
	}
	for i := 0; i < w; i++ {
		MarkOutsideGround(Pos{i, h - 1}, tiles, w, h)
	}
	for i := 0; i < h; i++ {
		MarkOutsideGround(Pos{0, i}, tiles, w, h)
	}
	for i := 0; i < h; i++ {
		MarkOutsideGround(Pos{w - 1, i}, tiles, w, h)
	}
}

func MarkOutsideGround(pos Pos, tiles [][]rune, w, h int) {
	toLeft := Pos{pos.x - 1, pos.y}
	toRight := Pos{pos.x + 1, pos.y}
	toUp := Pos{pos.x, pos.y - 1}
	toDown := Pos{pos.x, pos.y + 1}

	if pos.x >= 0 && tiles[pos.y][pos.x] == '.' {
		tiles[pos.y][pos.x] = '0'
	}
	if toLeft.x >= 0 && tiles[toLeft.y][toLeft.x] == '.' {
		tiles[toLeft.y][toLeft.x] = '0'
		MarkOutsideGround(toLeft, tiles, w, h)
	}
	if toRight.x < w && tiles[toRight.y][toRight.x] == '.' {
		tiles[toRight.y][toRight.x] = '0'
		MarkOutsideGround(toRight, tiles, w, h)
	}
	if toUp.y >= 0 && tiles[toUp.y][toUp.x] == '.' {
		tiles[toUp.y][toUp.x] = '0'
		MarkOutsideGround(toUp, tiles, w, h)
	}
	if toDown.y < h && tiles[toDown.y][toDown.x] == '.' {
		tiles[toDown.y][toDown.x] = '0'
		MarkOutsideGround(toDown, tiles, w, h)
	}
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

	rowsCnt := int(math.Abs(float64(rowMin))+math.Abs(float64(rowMax)))*2 + 1
	colsCnt := int(math.Abs(float64(colMin))+math.Abs(float64(colMax)))*2 + 1
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
		steps, err := strconv.ParseInt(color, 16, 32)
		if err != nil {
			panic(err)
		}
		move.steps = int(steps)
		moves = append(moves, move)
	}
	return moves
}

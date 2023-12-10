package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
}

type Tile int8

const (
	// | is a vertical pipe connecting north and south.
	// - is a horizontal pipe connecting east and west.
	// L is a 90-degree bend connecting north and east.
	// J is a 90-degree bend connecting north and west.
	// 7 is a 90-degree bend connecting south and west.
	// F is a 90-degree bend connecting south and east.
	// . is ground; there is no pipe in this tile.
	// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
	Pipe Tile = iota
	Minus
	L
	J
	Seven
	F
	Dot
	S
)

func part1() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	res := Part1GetMid(scanner)
	fmt.Println(res)
}

func Part1GetMid(scanner *bufio.Scanner) int {
	tiles := [][]Tile{}
	currentPos := Pos{0, 0}
	x := 0
	for scanner.Scan() {
		s := scanner.Text()
		line := []Tile{}
		for y, ch := range s {
			switch ch {
			case '|':
				line = append(line, Pipe)
			case '-':
				line = append(line, Minus)
			case 'L':
				line = append(line, L)
			case 'J':
				line = append(line, J)
			case '7':
				line = append(line, Seven)
			case 'F':
				line = append(line, F)
			case '.':
				line = append(line, Dot)
			case 'S':
				line = append(line, S)
				currentPos = Pos{x, y}
			}
		}
		tiles = append(tiles, line)
		x++
	}
	stepCnt := Part1FindPaths(tiles, currentPos)
	return stepCnt / 2
}

type Pos struct {
	x int
	y int
}

func Part1FindPaths(tiles [][]Tile, pos Pos) int {
	width := len(tiles[0])
	height := len(tiles)

	toLeft := Pos{pos.x - 1, pos.y}
	toRight := Pos{pos.x + 1, pos.y}
	toUp := Pos{pos.x, pos.y - 1}
	toDown := Pos{pos.x, pos.y + 1}

	if toLeft.x >= 0 && tiles[toLeft.y][toLeft.x] == Minus {
		return stepToTheLeft(tiles, pos, 0)
	} else if toLeft.x >= 0 && tiles[toLeft.y][toLeft.x] == F {
		return stepToTheLeftAndDownward(tiles, pos, 0)
	} else if toLeft.x >= 0 && tiles[toLeft.y][toLeft.x] == L {
		return stepToTheLeftAndUpward(tiles, pos, 0)
	} else if toRight.x < width && tiles[toRight.y][toRight.x] == Minus {
		return stepToTheRight(tiles, pos, 0)
	} else if toRight.x < width && tiles[toRight.y][toRight.x] == Seven {
		return stepToTheRightAndDownward(tiles, pos, 0)
	} else if toRight.x < width && tiles[toRight.y][toRight.x] == J {
		return stepToTheRightAndUpward(tiles, pos, 0)
	} else if toUp.y >= 0 && tiles[toUp.y][toUp.x] == Pipe {
		return stepToTheUpward(tiles, pos, 0)
	} else if toUp.y >= 0 && tiles[toUp.y][toUp.x] == Seven {
		return stepToTheUpwardAndLeft(tiles, pos, 0)
	} else if toUp.y >= 0 && tiles[toUp.y][toUp.x] == F {
		return stepToTheUpwardAndRight(tiles, pos, 0)
	} else if toDown.y < height && tiles[toDown.y][toDown.x] == Pipe {
		return stepToTheDownward(tiles, pos, 0)
	} else if toDown.y < height && tiles[toDown.y][toDown.x] == L {
		return stepToTheDownwardAndRight(tiles, pos, 0)
	} else if toDown.y < height && tiles[toDown.y][toDown.x] == J {
		return stepToTheDownwardAndLeft(tiles, pos, 0)
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheLeft(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.x -= 2
	switch tiles[pos.y][pos.x] {
	case F:
		return stepToTheDownward(tiles, pos, stepCount+2)
	case L:
		return stepToTheUpward(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheLeftAndUpward(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.x -= 1
	pos.y -= 1
	switch tiles[pos.y][pos.x] {
	case F:
		return stepToTheDownward(tiles, pos, stepCount+2)
	case L:
		return stepToTheUpward(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheLeftAndDownward(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.x -= 1
	pos.y += 1
	switch tiles[pos.y][pos.x] {
	case F:
		return stepToTheDownward(tiles, pos, stepCount+2)
	case L:
		return stepToTheUpward(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheRight(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.x += 2
	switch tiles[pos.y][pos.x] {
	case Seven:
		return stepToTheDownward(tiles, pos, stepCount+2)
	case J:
		return stepToTheUpward(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheRightAndUpward(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.x += 1
	pos.y -= 1
	switch tiles[pos.y][pos.x] {
	case Seven:
		return stepToTheDownward(tiles, pos, stepCount+2)
	case J:
		return stepToTheUpward(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheRightAndDownward(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.x += 1
	pos.y += 1
	switch tiles[pos.y][pos.x] {
	case Seven:
		return stepToTheDownward(tiles, pos, stepCount+2)
	case J:
		return stepToTheUpward(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheUpward(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.y -= 2
	switch tiles[pos.y][pos.x] {
	case F:
		return stepToTheRight(tiles, pos, stepCount+2)
	case Seven:
		return stepToTheLeft(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheUpwardAndLeft(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.y -= 1
	pos.x -= 1
	switch tiles[pos.y][pos.x] {
	case F:
		return stepToTheRight(tiles, pos, stepCount+2)
	case Seven:
		return stepToTheLeft(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheUpwardAndRight(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.y -= 1
	pos.x += 1
	switch tiles[pos.y][pos.x] {
	case F:
		return stepToTheRight(tiles, pos, stepCount+2)
	case Seven:
		return stepToTheLeft(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheDownward(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.y += 2
	switch tiles[pos.y][pos.x] {
	case L:
		return stepToTheRight(tiles, pos, stepCount+2)
	case J:
		return stepToTheLeft(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheDownwardAndLeft(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.y += 1
	pos.x -= 1
	switch tiles[pos.y][pos.x] {
	case L:
		return stepToTheRight(tiles, pos, stepCount+2)
	case J:
		return stepToTheLeft(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

func stepToTheDownwardAndRight(tiles [][]Tile, pos Pos, stepCount int) int {
	pos.y += 1
	pos.x += 1
	switch tiles[pos.y][pos.x] {
	case L:
		return stepToTheRight(tiles, pos, stepCount+2)
	case J:
		return stepToTheLeft(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2
	}
	panic("undefined step: " + strconv.Itoa(int(tiles[pos.y][pos.x])))
}

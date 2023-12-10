package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	part1()
}

type Tile string

const (
	// | is a vertical pipe connecting north and south.
	// - is a horizontal pipe connecting east and west.
	// L is a 90-degree bend connecting north and east.
	// J is a 90-degree bend connecting north and west.
	// 7 is a 90-degree bend connecting south and west.
	// F is a 90-degree bend connecting south and east.
	// . is ground; there is no pipe in this tile.
	// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
	Pipe  Tile = "|"
	Minus      = "-"
	L          = "L"
	J          = "J"
	Seven      = "7"
	F          = "F"
	Dot        = "."
	S          = "S"
)

func part1() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := Part1GetMid(scanner)
	fmt.Println(result)
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

	result := -1
	if toLeft.x >= 0 {
		switch tiles[toLeft.y][toLeft.x] {
		case Minus:
			tmp, err := stepToTheLeft(tiles, toLeft, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", tmp, "onTheLeft")
				result = int(math.Max(float64(result), float64(tmp)))
			}
		case F:
			tmp, err := stepToTheLeftAndDownward(tiles, toLeft, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", tmp, "onTheDownward")
				result = int(math.Max(float64(result), float64(tmp)))
			}
		case L:
			tmp, err := stepToTheLeftAndUpward(tiles, toLeft, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", tmp, "onTheUpward")
				result = int(math.Max(float64(result), float64(tmp)))
			}
		}
	}
	if toRight.x < width {
		switch tiles[toRight.y][toRight.x] {
		case Minus:
			tmp, err := stepToTheRight(tiles, toRight, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", tmp, "onTheRight")
				result = int(math.Max(float64(result), float64(tmp)))
			}
		case Seven:
			tmp, err := stepToTheRightAndDownward(tiles, toRight, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", tmp, "onTheRightAndDownward")
				result = int(math.Max(float64(result), float64(tmp)))
			}
		case J:
			tmp, err := stepToTheRightAndUpward(tiles, toRight, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", tmp, "onTheRightAndUpward")
				result = int(math.Max(float64(result), float64(tmp)))
			}
		}
	}
	if toUp.y >= 0 {
		switch tiles[toUp.y][toUp.x] {
		case Pipe:
			tmp, err := stepToTheUpward(tiles, toUp, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				result = int(math.Max(float64(result), float64(tmp)))
			}
		case Seven:
			tmp, err := stepToTheUpwardAndLeft(tiles, toUp, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				result = int(math.Max(float64(result), float64(tmp)))
			}
		case F:
			tmp, err := stepToTheUpwardAndRight(tiles, toUp, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				result = int(math.Max(float64(result), float64(tmp)))
			}
		}
	}
	if toDown.y < height {
		switch tiles[toDown.y][toDown.x] {
		case Pipe:
			tmp, err := stepToTheDownward(tiles, toDown, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				result = int(math.Max(float64(result), float64(tmp)))
			}
		case L:
			tmp, err := stepToTheDownwardAndRight(tiles, toDown, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				result = int(math.Max(float64(result), float64(tmp)))
			}
		case J:
			tmp, err := stepToTheDownwardAndLeft(tiles, toDown, 1)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				result = int(math.Max(float64(result), float64(tmp)))
			}
		}
	}
	return result
}

func stepToTheLeft(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheLeft", tiles[pos.y][pos.x])
	pos.x -= 1
	switch tiles[pos.y][pos.x] {
	case Minus:
		return stepToTheLeft(tiles, pos, stepCount+1)
	case F:
		return stepToTheDownward(tiles, pos, stepCount+1)
	case L:
		return stepToTheUpward(tiles, pos, stepCount+1)
	case S:
		return stepCount + 1, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheLeftAndUpward(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheLeftAndUpward", tiles[pos.y][pos.x])
	pos.x -= 1
	pos.y -= 1
	switch tiles[pos.y][pos.x] {
	case F:
		return stepToTheDownward(tiles, pos, stepCount+2)
	case L:
		return stepToTheUpward(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheLeftAndDownward(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheLeftAndDownward", tiles[pos.y][pos.x])
	pos.x -= 1
	pos.y += 1
	switch tiles[pos.y][pos.x] {
	case F:
		return stepToTheDownward(tiles, pos, stepCount+2)
	case L:
		return stepToTheUpward(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheRight(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheRight", tiles[pos.y][pos.x])
	pos.x += 1
	switch tiles[pos.y][pos.x] {
	case Minus:
		return stepToTheRight(tiles, pos, stepCount+1)
	case Seven:
		return stepToTheDownward(tiles, pos, stepCount+1)
	case J:
		return stepToTheUpward(tiles, pos, stepCount+1)
	case S:
		return stepCount + 1, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheRightAndUpward(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheRightAndUpward", tiles[pos.y][pos.x])
	pos.x += 1
	pos.y -= 1
	switch tiles[pos.y][pos.x] {
	case Seven:
		return stepToTheDownward(tiles, pos, stepCount+2)
	case J:
		return stepToTheUpward(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheRightAndDownward(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheRightAndDownward", tiles[pos.y][pos.x])
	pos.x += 1
	pos.y += 1
	switch tiles[pos.y][pos.x] {
	case Seven:
		return stepToTheDownward(tiles, pos, stepCount+2)
	case J:
		return stepToTheUpward(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheUpward(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheUpward", tiles[pos.y][pos.x])
	pos.y -= 1
	switch tiles[pos.y][pos.x] {
	case Pipe:
		return stepToTheUpward(tiles, pos, stepCount+1)
	case F:
		return stepToTheRight(tiles, pos, stepCount+1)
	case Seven:
		return stepToTheLeft(tiles, pos, stepCount+1)
	case S:
		return stepCount + 1, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheUpwardAndLeft(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheUpwardAndLeft", tiles[pos.y][pos.x])
	pos.y -= 1
	pos.x -= 1
	switch tiles[pos.y][pos.x] {
	case F:
		return stepToTheRight(tiles, pos, stepCount+2)
	case Seven:
		return stepToTheLeft(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheUpwardAndRight(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheUpwardAndRight", tiles[pos.y][pos.x])
	pos.y -= 1
	pos.x += 1
	switch tiles[pos.y][pos.x] {
	case F:
		return stepToTheRight(tiles, pos, stepCount+2)
	case Seven:
		return stepToTheLeft(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheDownward(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheDownward", tiles[pos.y][pos.x])
	pos.y += 1
	switch tiles[pos.y][pos.x] {
	case Pipe:
		return stepToTheDownward(tiles, pos, stepCount+1)
	case L:
		return stepToTheRight(tiles, pos, stepCount+1)
	case J:
		return stepToTheLeft(tiles, pos, stepCount+1)
	case S:
		return stepCount + 1, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheDownwardAndLeft(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheDownwardAndLeft", tiles[pos.y][pos.x])
	pos.y += 1
	pos.x -= 1
	switch tiles[pos.y][pos.x] {
	case L:
		return stepToTheRight(tiles, pos, stepCount+2)
	case J:
		return stepToTheLeft(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheDownwardAndRight(tiles [][]Tile, pos Pos, stepCount int) (int, error) {
	fmt.Println("stepToTheDownwardAndRight", tiles[pos.y][pos.x])
	pos.y += 1
	pos.x += 1
	switch tiles[pos.y][pos.x] {
	case L:
		return stepToTheRight(tiles, pos, stepCount+2)
	case J:
		return stepToTheLeft(tiles, pos, stepCount+2)
	case S:
		return stepCount + 2, nil
	}
	return 0, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

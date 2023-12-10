package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
)

func main() {
	part1()
	part2()
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
	ZERO       = "0"
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

func part2() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := Part2GetEnclosedTilesCnt(scanner)
	fmt.Println(result)
}

func Part1GetMid(scanner *bufio.Scanner) int {
	tiles, currentPos := Part1GetTiles(scanner)
	steps := Part1FindPaths(tiles, currentPos)
	return len(steps) / 2
}

func Part1GetTiles(scanner *bufio.Scanner) ([][]Tile, Pos) {
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
				currentPos = Pos{y, x}
			}
		}
		tiles = append(tiles, line)
		x++
	}
	return tiles, currentPos
}

type Pos struct {
	x int
	y int
}

func Part1FindPaths(tiles [][]Tile, pos Pos) []Pos {
	fmt.Println("Part1FindPaths", tiles[pos.y][pos.x])
	width := len(tiles[0])
	height := len(tiles)

	toLeft := Pos{pos.x - 1, pos.y}
	toRight := Pos{pos.x + 1, pos.y}
	toUp := Pos{pos.x, pos.y - 1}
	toDown := Pos{pos.x, pos.y + 1}

	if toLeft.x >= 0 {
		switch tiles[toLeft.y][toLeft.x] {
		case Minus:
			steps, err := stepToTheLeft(tiles, toLeft, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheLeft")
				return steps
			}
		case F:
			steps, err := stepToTheDownward(tiles, toLeft, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheDownward")
				return steps
			}
		case L:
			steps, err := stepToTheUpward(tiles, toLeft, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheUpward")
				return steps
			}
		}
	}
	if toRight.x < width {
		switch tiles[toRight.y][toRight.x] {
		case Minus:
			steps, err := stepToTheRight(tiles, toRight, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheRight")
				return steps
			}
		case Seven:
			steps, err := stepToTheDownward(tiles, toRight, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheDownward")
				return steps
			}
		case J:
			steps, err := stepToTheUpward(tiles, toRight, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheUpward")
				return steps
			}
		}
	}
	if toUp.y >= 0 {
		switch tiles[toUp.y][toUp.x] {
		case Pipe:
			steps, err := stepToTheUpward(tiles, toUp, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheUpward")
				return steps
			}
		case Seven:
			steps, err := stepToTheLeft(tiles, toUp, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheLeft")
				return steps
			}
		case F:
			steps, err := stepToTheRight(tiles, toUp, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheRight")
				return steps
			}
		}
	}
	if toDown.y < height {
		switch tiles[toDown.y][toDown.x] {
		case Pipe:
			steps, err := stepToTheDownward(tiles, toDown, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheDownward")
				return steps
			}
		case L:
			steps, err := stepToTheRight(tiles, toDown, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheRight")
				return steps
			}
		case J:
			steps, err := stepToTheLeft(tiles, toDown, []Pos{pos})
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("found", len(steps), "onTheLeft")
				return steps
			}
		}
	}
	panic("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheLeft(tiles [][]Tile, pos Pos, steps []Pos) ([]Pos, error) {
	steps = append(steps, pos)
	//fmt.Println("stepToTheLeft", tiles[pos.y][pos.x])
	pos.x -= 1
	switch tiles[pos.y][pos.x] {
	case Minus:
		return stepToTheLeft(tiles, pos, steps)
	case F:
		return stepToTheDownward(tiles, pos, steps)
	case L:
		return stepToTheUpward(tiles, pos, steps)
	case S:
		return steps, nil
	}
	return nil, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheRight(tiles [][]Tile, pos Pos, steps []Pos) ([]Pos, error) {
	steps = append(steps, pos)
	//fmt.Println("stepToTheRight", tiles[pos.y][pos.x])
	pos.x += 1
	switch tiles[pos.y][pos.x] {
	case Minus:
		return stepToTheRight(tiles, pos, steps)
	case Seven:
		return stepToTheDownward(tiles, pos, steps)
	case J:
		return stepToTheUpward(tiles, pos, steps)
	case S:
		return steps, nil
	}
	return nil, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheUpward(tiles [][]Tile, pos Pos, steps []Pos) ([]Pos, error) {
	steps = append(steps, pos)
	//fmt.Println("stepToTheUpward", tiles[pos.y][pos.x])
	pos.y -= 1
	switch tiles[pos.y][pos.x] {
	case Pipe:
		return stepToTheUpward(tiles, pos, steps)
	case F:
		return stepToTheRight(tiles, pos, steps)
	case Seven:
		return stepToTheLeft(tiles, pos, steps)
	case S:
		return steps, nil
	}
	return nil, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func stepToTheDownward(tiles [][]Tile, pos Pos, steps []Pos) ([]Pos, error) {
	steps = append(steps, pos)
	//fmt.Println("stepToTheDownward", tiles[pos.y][pos.x])
	pos.y += 1
	switch tiles[pos.y][pos.x] {
	case Pipe:
		return stepToTheDownward(tiles, pos, steps)
	case L:
		return stepToTheRight(tiles, pos, steps)
	case J:
		return stepToTheLeft(tiles, pos, steps)
	case S:
		return steps, nil
	}
	return nil, errors.New("undefined step: " + string(tiles[pos.y][pos.x]))
}

func Part2GetEnclosedTilesCnt(scanner *bufio.Scanner) int {
	tiles, currentPos := Part1GetTiles(scanner)
	steps := Part1FindPaths(tiles, currentPos)
	ground := 0
	for y, row := range tiles {
		for x, col := range row {
			if slices.Contains(steps, Pos{x, y}) {
				col = col
				fmt.Print(col)
			} else {
				tiles[y][x] = Dot
				fmt.Print(".")
				ground += 1
			}
		}
		fmt.Println("")
	}

	for i := 0; i < len(tiles[0]); i++ {
		ground -= MarkOutsideGround(Pos{i, 0}, &tiles)
	}
	for i := 0; i < len(tiles[0]); i++ {
		ground -= MarkOutsideGround(Pos{i, len(tiles) - 1}, &tiles)
	}
	for i := 0; i < len(tiles); i++ {
		ground -= MarkOutsideGround(Pos{0, i}, &tiles)
	}
	for i := 0; i < len(tiles); i++ {
		ground -= MarkOutsideGround(Pos{len(tiles[0]) - 1, i}, &tiles)
	}
	fmt.Println("---")
	for _, row := range tiles {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println("")
	}

	return ground
}

func MarkOutsideGround(pos Pos, tiles *[][]Tile) int {
	cnt := 0
	toLeft := Pos{pos.x - 1, pos.y}
	toRight := Pos{pos.x + 1, pos.y}
	toUp := Pos{pos.x, pos.y - 1}
	toUpLeft := Pos{pos.x - 1, pos.y - 1}
	toUpRight := Pos{pos.x + 1, pos.y - 1}
	toDown := Pos{pos.x, pos.y + 1}

	if pos.y < len((*tiles))-1 && toUp.y > 0 && (*tiles)[toUp.y][toUp.x] == L && toUpLeft.y > 0 && toUpLeft.x > 0 && (*tiles)[toUpLeft.y][toUpLeft.x] != Dot && (*tiles)[toUpLeft.y][toUpLeft.x] != ZERO {
		cnt -= AddCol(pos.x, tiles)
	}
	if toUp.y > 0 && (*tiles)[toUp.y][toUp.x] == J && toUpRight.y > 0 && toUpRight.x < len((*tiles)[0])-1 && (*tiles)[toUpRight.y][toUpRight.x] != Dot && (*tiles)[toUpRight.y][toUpRight.x] != ZERO {
		cnt -= AddCol(pos.x+1, tiles)
	}
	if pos.x >= 0 && (*tiles)[pos.y][pos.x] == Dot {
		(*tiles)[pos.y][pos.x] = ZERO
		cnt += 1
	}
	if toLeft.x >= 0 && (*tiles)[toLeft.y][toLeft.x] == Dot {
		(*tiles)[toLeft.y][toLeft.x] = ZERO
		cnt += 1 + MarkOutsideGround(toLeft, tiles)
	}
	if toRight.x < len((*tiles)[0]) && (*tiles)[toRight.y][toRight.x] == Dot {
		(*tiles)[toRight.y][toRight.x] = ZERO
		cnt += 1 + MarkOutsideGround(toRight, tiles)
	}
	if toUp.y >= 0 && (*tiles)[toUp.y][toUp.x] == Dot {
		(*tiles)[toUp.y][toUp.x] = ZERO
		cnt += 1 + MarkOutsideGround(toUp, tiles)
	}
	if toDown.y < len(*tiles) && (*tiles)[toDown.y][toDown.x] == Dot {
		(*tiles)[toDown.y][toDown.x] = ZERO
		cnt += 1 + MarkOutsideGround(toDown, tiles)
	}
	return cnt
}

func AddCol(colIdx int, tiles *[][]Tile) int {
	dotsAdded := 0
	for y, row := range *tiles {
		line := []Tile{}
		for x, col := range row {
			if x == colIdx {
				switch col {
				case Minus:
					line = append(line, Minus)
				default:
					line = append(line, Dot)
					dotsAdded++
				}
			}
			line = append(line, col)
		}
		(*tiles)[y] = line
	}
	return dotsAdded
}

func AddRow(rowIdx int, tiles *[][]Tile) {

}

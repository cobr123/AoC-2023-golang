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
	BottomToTop   Tile = "↑" // "|"
	TopToBottom        = "↓" // "|"
	RightToLeft        = "←" // "-"
	LeftToRight        = "→" // "-"
	TopToRight         = "→" // "L"
	RightToTop         = "↑" // "L"
	TopToLeft          = "←" // "J"
	LeftToTop          = "↑" // "J"
	LeftToBottom       = "↓" // "7"
	BottomToLeft       = "←" // "7"
	RightToBottom      = "↓" // "F"
	BottomToRight      = "→" // "F"
	Dot                = "."
	S                  = "S"
	ZERO               = "0"
	Pipe               = "|"
	Minus              = "-"
	L                  = "L"
	J                  = "J"
	Seven              = "7"
	F                  = "F"
)

type Direction int8

const (
	toTheTop Direction = iota
	toTheRight
	toTheBottom
	toTheLeft
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
	leftTopStep := slices.MinFunc(steps, func(a, b Pos) int {
		if a.y != b.y {
			return a.y - b.y
		} else {
			return a.x - b.x
		}
	})
	repeatTil := 0
	for i, step := range steps {
		if step == leftTopStep {
			repeatTil = i
			break
		}
	}
	fmt.Println(leftTopStep, repeatTil, tiles[leftTopStep.y][leftTopStep.x])
	direction := initDirection(tiles, steps[repeatTil], steps[repeatTil+1])
	for i := repeatTil; i < len(steps); i++ {
		direction = markDirections(tiles, steps[i], direction)
	}
	for i := 0; i < repeatTil; i++ {
		direction = markDirections(tiles, steps[i], direction)
	}
	fmt.Println("---with direction")
	for _, row := range tiles {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println("")
	}
	width := len(tiles[0])
	height := len(tiles)

	for y, row := range tiles {
		for x, col := range row {
			switch col {
			case BottomToTop:
				onLeft := Pos{x - 1, y}
				ground -= MarkOutsideGround(onLeft, tiles, width, height)
			case TopToBottom:
				onRight := Pos{x + 1, y}
				ground -= MarkOutsideGround(onRight, tiles, width, height)
			case LeftToRight:
				onTop := Pos{x, y - 1}
				ground -= MarkOutsideGround(onTop, tiles, width, height)
			case RightToLeft:
				onBottom := Pos{x, y + 1}
				ground -= MarkOutsideGround(onBottom, tiles, width, height)
			}
		}
	}
	//for i := 0; i < width; i++ {
	//	ground -= MarkOutsideGround(Pos{i, 0}, tiles, width, height)
	//}
	//for i := 0; i < width; i++ {
	//	ground -= MarkOutsideGround(Pos{i, height - 1}, tiles, width, height)
	//}
	//for i := 0; i < height; i++ {
	//	ground -= MarkOutsideGround(Pos{0, i}, tiles, width, height)
	//}
	//for i := 0; i < height; i++ {
	//	ground -= MarkOutsideGround(Pos{width - 1, i}, tiles, width, height)
	//}
	fmt.Println("---with zero")
	for _, row := range tiles {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println("")
	}
	return ground
}

func initDirection(tiles [][]Tile, step Pos, nextStep Pos) Direction {
	switch tiles[step.y][step.x] {
	case S:
		return initDirection(tiles, nextStep, Pos{})
	case Minus:
		return toTheRight
	case L:
		return toTheRight
	case F:
		return toTheTop
	}
	fmt.Println(step)
	panic("undefined step:" + tiles[step.y][step.x])
}

func markDirections(tiles [][]Tile, step Pos, direction Direction) Direction {
	switch tiles[step.y][step.x] {
	case Minus:
		if direction == toTheRight {
			tiles[step.y][step.x] = LeftToRight
		} else {
			tiles[step.y][step.x] = RightToLeft
		}
		return direction
	case Seven:
		if direction == toTheRight {
			tiles[step.y][step.x] = LeftToBottom
			return toTheBottom
		} else {
			tiles[step.y][step.x] = BottomToLeft
			return toTheLeft
		}
	case Pipe:
		if direction == toTheBottom {
			tiles[step.y][step.x] = TopToBottom
		} else {
			tiles[step.y][step.x] = BottomToTop
		}
	case J:
		if direction == toTheRight {
			tiles[step.y][step.x] = LeftToTop
			return toTheTop
		} else {
			tiles[step.y][step.x] = TopToLeft
			return toTheLeft
		}
	case L:
		if direction == toTheLeft {
			tiles[step.y][step.x] = RightToTop
			return toTheTop
		} else {
			tiles[step.y][step.x] = TopToRight
			return toTheRight
		}
	case F:
		if direction == toTheTop {
			tiles[step.y][step.x] = BottomToRight
			return toTheRight
		} else {
			tiles[step.y][step.x] = RightToBottom
			return toTheBottom
		}
	}
	return direction
}

func MarkOutsideGround(pos Pos, tiles [][]Tile, w, h int) int {
	cnt := 0

	if pos.x >= 0 && pos.x < w && pos.y >= 0 && pos.y < h && tiles[pos.y][pos.x] == Dot {
		tiles[pos.y][pos.x] = ZERO
		cnt += 1

		toLeft := Pos{pos.x - 1, pos.y}
		toRight := Pos{pos.x + 1, pos.y}
		toUp := Pos{pos.x, pos.y - 1}
		toDown := Pos{pos.x, pos.y + 1}
		cnt += MarkOutsideGround(toLeft, tiles, w, h)
		cnt += MarkOutsideGround(toRight, tiles, w, h)
		cnt += MarkOutsideGround(toUp, tiles, w, h)
		cnt += MarkOutsideGround(toDown, tiles, w, h)
	}
	return cnt
}

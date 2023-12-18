package main

import (
	"bufio"
	"fmt"
	"os"
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

	sum := Part2GetSum(scanner, 1e6)
	fmt.Println(sum)
}

func Part1GetSum(scanner *bufio.Scanner) int {
	universe := Part1ParseUniverse(scanner)
	Part1ExpandUniverse(universe, 2)
	galaxies := Part1GetGalaxies(universe)
	paths := Part1GetPaths(galaxies, universe)
	sum := 0
	for _, path := range paths {
		sum += path.stepsCnt
	}
	return sum
}

func Part2GetSum(scanner *bufio.Scanner, emptySpaceSize int) int {
	universe := Part1ParseUniverse(scanner)
	Part1ExpandUniverse(universe, emptySpaceSize)
	galaxies := Part1GetGalaxies(universe)
	paths := Part1GetPaths(galaxies, universe)
	sum := 0
	for _, path := range paths {
		sum += path.stepsCnt
	}
	return sum
}

type Space struct {
	ch    rune
	hsize int
	vsize int
}

type Pos struct {
	row int
	col int
}

type Path struct {
	from     Pos
	to       Pos
	stepsCnt int
}

func Part1GetPath(from Pos, to Pos, universe [][]Space) Path {
	path := Path{from, to, 0}
	minCol := 0
	maxCol := 0
	minRow := 0
	maxRow := 0
	if path.from.col < path.to.col {
		minCol = path.from.col
		maxCol = path.to.col
	} else {
		minCol = path.to.col
		maxCol = path.from.col
	}
	if path.from.row < path.to.row {
		minRow = path.from.row
		maxRow = path.to.row
	} else {
		minRow = path.to.row
		maxRow = path.from.row
	}
	for row := minRow; row < maxRow; row++ {
		path.stepsCnt += universe[row][minCol].hsize
	}
	for col := minCol; col < maxCol; col++ {
		path.stepsCnt += universe[maxRow][col].vsize
	}
	return path
}

func Part1GetPaths(galaxies []Pos, universe [][]Space) []Path {
	paths := []Path{}
	for i := 0; i < len(galaxies); i++ {
		for k := i + 1; k < len(galaxies); k++ {
			path := Part1GetPath(galaxies[i], galaxies[k], universe)
			paths = append(paths, path)
		}
	}
	return paths
}

func Part1GetGalaxies(universe [][]Space) []Pos {
	galaxies := []Pos{}
	for row, line := range universe {
		for col, item := range line {
			if item.ch == '#' {
				galaxies = append(galaxies, Pos{row, col})
			}
		}
	}
	return galaxies
}

func Part1ParseUniverse(scanner *bufio.Scanner) [][]Space {
	universe := [][]Space{}
	for scanner.Scan() {
		s := scanner.Text()
		line := Part1ParseLine(s)
		universe = append(universe, line)
	}
	return universe
}

func Part1ParseLine(s string) []Space {
	line := []Space{}
	for _, ch := range s {
		line = append(line, Space{ch, 1, 1})
	}
	return line
}

func Part1ExpandUniverse(universe [][]Space, emptySpaceSize int) {
	Part1ExpandUniverseH(universe, emptySpaceSize)
	Part1ExpandUniverseV(universe, emptySpaceSize)
}

func Part1ExpandUniverseH(universe [][]Space, emptySpaceSize int) {
	for row := 0; row < len(universe); row++ {
		noGalaxies := true
		for col := 0; col < len(universe[row]); col++ {
			if universe[row][col].ch == '#' {
				noGalaxies = false
				break
			}
		}
		if noGalaxies {
			for col := 0; col < len(universe[row]); col++ {
				universe[row][col].hsize = emptySpaceSize
			}
		}
	}
}

func Part1ExpandUniverseV(universe [][]Space, emptySpaceSize int) {
	for col := 0; col < len(universe[0]); col++ {
		noGalaxies := true
		for row := 0; row < len(universe); row++ {
			if universe[row][col].ch == '#' {
				noGalaxies = false
				break
			}
		}
		if noGalaxies {
			for row := 0; row < len(universe); row++ {
				universe[row][col].vsize = emptySpaceSize
			}
		}
	}
}

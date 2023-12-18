package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part1GetSum(scanner)
	expected := 374
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1GetGalaxies(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	universe := Part1ParseUniverse(scanner)
	got := len(Part1GetGalaxies(universe))
	expected := 9
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1GetPaths(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	universe := Part1ParseUniverse(scanner)
	galaxies := Part1GetGalaxies(universe)
	got := len(Part1GetPaths(galaxies, universe))
	expected := 36
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1ExpandUniverse1(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	universe := Part1ParseUniverse(scanner)
	Part1ExpandUniverse(universe, 2)
	got := universe[0][2]
	expected := Space{'.', 1, 2}
	if got != expected {
		t.Errorf("got = %v; want %v", got, expected)
	}
}

func TestPart1ExpandUniverse2(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	universe := Part1ParseUniverse(scanner)
	Part1ExpandUniverse(universe, 2)
	got := universe[3][1]
	expected := Space{'.', 2, 1}
	if got != expected {
		t.Errorf("got = %v; want %v", got, expected)
	}
}

func TestPart1ExpandUniverse3(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	universe := Part1ParseUniverse(scanner)
	Part1ExpandUniverse(universe, 2)
	got := universe[3][2]
	expected := Space{'.', 2, 2}
	if got != expected {
		t.Errorf("got = %v; want %v", got, expected)
	}
}

func TestPart1GetPath1(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	universe := Part1ParseUniverse(scanner)
	from := Pos{0, 3}
	to := Pos{2, 0}
	got := Part1GetPath(from, to, universe)
	expected := Path{from, to, 5}
	if got != expected {
		t.Errorf("got = %v; want %v", got, expected)
	}
}

func TestPart1GetPath2(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	universe := Part1ParseUniverse(scanner)
	Part1ExpandUniverse(universe, 2)
	from := Pos{0, 4}
	to := Pos{2, 0}
	got := Part1GetPath(from, to, universe)
	expected := Path{from, to, 7}
	if got != expected {
		t.Errorf("got = %v; want %v", got, expected)
	}
}

func TestPart1GetPath3(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	universe := Part1ParseUniverse(scanner)
	Part1ExpandUniverse(universe, 2)
	from := Pos{0, 4}
	to := Pos{6, 1}
	got := Part1GetPath(from, to, universe)
	expected := Path{from, to, 11}
	if got != expected {
		t.Errorf("got = %v; want %v", got, expected)
	}
}

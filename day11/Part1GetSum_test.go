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
	got := len(Part1GetPaths(galaxies))
	expected := 36
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1ExpandUniverse(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	universe := Part1ParseUniverse(scanner)
	got := Part1ExpandUniverse(universe)
	expected := "....#........\n.........#...\n#............\n.............\n.............\n........#....\n.#...........\n............#\n.............\n.............\n.........#...\n#....#......."
	b2 := bytes.NewReader([]byte(expected))
	scanner2 := bufio.NewScanner(b2)
	expected2 := Part1ParseUniverse(scanner2)
	for i := 0; i < len(got); i++ {
		if len(got[i]) != len(expected2[i]) {
			t.Errorf("got = %v; want %v", string(got[i]), string(expected2[i]))
		}
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != expected2[i][j] {
				t.Errorf("got = %v; want %v", string(got[i][j]), string(expected2[i][j]))
			}
		}
	}
}

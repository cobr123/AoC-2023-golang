package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#...."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner)
	expected := 136
	if sum != expected {
		t.Errorf("got = %d; want %d", sum, expected)
	}
}

func TestPart1SlideNorth(t *testing.T) {
	input := "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#...."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	dish := Part1ParseDish(scanner)
	Part1SlideNorth(dish)
	expected := "OOOO.#.O..\nOO..#....#\nOO..O##..O\nO..#.OO...\n........#.\n..#....#.#\n..O..#.O.O\n..O.......\n#....###..\n#....#...."
	expectedReader := bytes.NewReader([]byte(expected))
	expectedScanner := bufio.NewScanner(expectedReader)
	expectedDish := Part1ParseDish(expectedScanner)
	for i := 0; i < len(dish); i++ {
		if string(dish[i]) != string(expectedDish[i]) {
			t.Errorf("line = %v, got = %v; want %v", i, string(dish[i]), string(expectedDish[i]))
		}
	}
}

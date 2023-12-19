package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2GetSum(t *testing.T) {
	input := "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#...."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part2GetSum(scanner)
	expected := 64
	if sum != expected {
		t.Errorf("got = %d; want %d", sum, expected)
	}
}

func TestPart2Slide1Times(t *testing.T) {
	input := "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#...."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	dish := Part1ParseDish(scanner)
	Part2SpinCycle(dish)
	expected := ".....#....\n....#...O#\n...OO##...\n.OO#......\n.....OOO#.\n.O#...O#.#\n....O#....\n......OOOO\n#...O###..\n#..OO#...."
	expectedReader := bytes.NewReader([]byte(expected))
	expectedScanner := bufio.NewScanner(expectedReader)
	expectedDish := Part1ParseDish(expectedScanner)
	for i := 0; i < len(dish); i++ {
		if string(dish[i]) != string(expectedDish[i]) {
			t.Errorf("line = %v, got = %v; want %v", i, string(dish[i]), string(expectedDish[i]))
		}
	}
}

func TestPart2Slide2Times(t *testing.T) {
	input := "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#...."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	dish := Part1ParseDish(scanner)
	Part2SpinCycle(dish)
	Part2SpinCycle(dish)
	expected := ".....#....\n....#...O#\n.....##...\n..O#......\n.....OOO#.\n.O#...O#.#\n....O#...O\n.......OOO\n#..OO###..\n#.OOO#...O"
	expectedReader := bytes.NewReader([]byte(expected))
	expectedScanner := bufio.NewScanner(expectedReader)
	expectedDish := Part1ParseDish(expectedScanner)
	for i := 0; i < len(dish); i++ {
		if string(dish[i]) != string(expectedDish[i]) {
			t.Errorf("line = %v, got = %v; want %v", i, string(dish[i]), string(expectedDish[i]))
		}
	}
}

func TestPart2Slide3Times(t *testing.T) {
	input := "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#...."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	dish := Part1ParseDish(scanner)
	Part2SpinCycle(dish)
	Part2SpinCycle(dish)
	Part2SpinCycle(dish)
	expected := ".....#....\n....#...O#\n.....##...\n..O#......\n.....OOO#.\n.O#...O#.#\n....O#...O\n.......OOO\n#...O###.O\n#.OOO#...O"
	expectedReader := bytes.NewReader([]byte(expected))
	expectedScanner := bufio.NewScanner(expectedReader)
	expectedDish := Part1ParseDish(expectedScanner)
	for i := 0; i < len(dish); i++ {
		if string(dish[i]) != string(expectedDish[i]) {
			t.Errorf("line = %v, got = %v; want %v", i, string(dish[i]), string(expectedDish[i]))
		}
	}
}

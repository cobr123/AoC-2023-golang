package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2GetSum(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part2GetSum(scanner)
	expected := 2286
	if sum != expected {
		t.Errorf("Part2 = %d; want %d", sum, expected)
	}
}

func TestPart2GetSumExample1(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part2GetSum(scanner)
	expected := 48
	if sum != expected {
		t.Errorf("Part2 = %d; want %d", sum, expected)
	}
}

func TestPart2GetSumExample2(t *testing.T) {
	input := "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part2GetSum(scanner)
	expected := 12
	if sum != expected {
		t.Errorf("Part2 = %d; want %d", sum, expected)
	}
}

func TestPart2GetSumExample3(t *testing.T) {
	input := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part2GetSum(scanner)
	expected := 1560
	if sum != expected {
		t.Errorf("Part2 = %d; want %d", sum, expected)
	}
}

func TestPart2GetSumExample4(t *testing.T) {
	input := "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part2GetSum(scanner)
	expected := 630
	if sum != expected {
		t.Errorf("Part2 = %d; want %d", sum, expected)
	}
}

func TestPart2GetSumExample5(t *testing.T) {
	input := "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part2GetSum(scanner)
	expected := 36
	if sum != expected {
		t.Errorf("Part2 = %d; want %d", sum, expected)
	}
}

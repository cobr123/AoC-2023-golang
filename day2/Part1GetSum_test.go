package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	cardsMaxValues := Cards{red: 12, green: 13, blue: 14}
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\nGame 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\nGame 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\nGame 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\nGame 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner, cardsMaxValues)
	expected := 8
	if sum != expected {
		t.Errorf("Part1 = %d; want %d", sum, expected)
	}
}

func TestPart1GetSumExample1(t *testing.T) {
	cardsMaxValues := Cards{red: 12, green: 13, blue: 14}
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner, cardsMaxValues)
	expected := 1
	if sum != expected {
		t.Errorf("Part1 = %d; want %d", sum, expected)
	}
}

func TestPart1GetSumExample2(t *testing.T) {
	cardsMaxValues := Cards{red: 12, green: 13, blue: 14}
	input := "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner, cardsMaxValues)
	expected := 2
	if sum != expected {
		t.Errorf("Part1 = %d; want %d", sum, expected)
	}
}

func TestPart1GetSumExample3(t *testing.T) {
	cardsMaxValues := Cards{red: 12, green: 13, blue: 14}
	input := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner, cardsMaxValues)
	expected := 0
	if sum != expected {
		t.Errorf("Part1 = %d; want %d", sum, expected)
	}
}

func TestPart1GetSumExample4(t *testing.T) {
	cardsMaxValues := Cards{red: 12, green: 13, blue: 14}
	input := "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner, cardsMaxValues)
	expected := 0
	if sum != expected {
		t.Errorf("Part1 = %d; want %d", sum, expected)
	}
}

func TestPart1GetSumExample5(t *testing.T) {
	cardsMaxValues := Cards{red: 12, green: 13, blue: 14}
	input := "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner, cardsMaxValues)
	expected := 5
	if sum != expected {
		t.Errorf("Part1 = %d; want %d", sum, expected)
	}
}

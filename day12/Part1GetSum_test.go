package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := "#.#.### 1,1,3\n.#...#....###. 1,1,3\n.#.###.#.###### 1,3,1,6\n####.#...#... 4,1,1\n#....######..#####. 1,6,5\n.###.##....# 3,2,1"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner)
	expected := 6
	if sum != expected {
		t.Errorf("got = %d; want %d", sum, expected)
	}
}

func TestPart1FindVariants1(t *testing.T) {
	input := "???.### 1,1,3"
	line := Part1ParseLine(input)
	got := Part1FindVariants(line)
	expected := 1
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1FindVariants2(t *testing.T) {
	input := ".??..??...?##. 1,1,3"
	line := Part1ParseLine(input)
	got := Part1FindVariants(line)
	expected := 4
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1FindVariants3(t *testing.T) {
	input := "?#?#?#?#?#?#?#? 1,3,1,6"
	line := Part1ParseLine(input)
	got := Part1FindVariants(line)
	expected := 1
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1FindVariants4(t *testing.T) {
	input := "????.#...#... 4,1,1"
	line := Part1ParseLine(input)
	got := Part1FindVariants(line)
	expected := 1
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1FindVariants5(t *testing.T) {
	input := "????.######..#####. 1,6,5"
	line := Part1ParseLine(input)
	got := Part1FindVariants(line)
	expected := 4
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1FindVariants6(t *testing.T) {
	input := "?###???????? 3,2,1"
	line := Part1ParseLine(input)
	got := Part1FindVariants(line)
	expected := 10
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

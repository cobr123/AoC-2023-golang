package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2GetSum10(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part2GetSum(scanner, 10)
	expected := 1030
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart2GetSum100(t *testing.T) {
	input := "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part2GetSum(scanner, 100)
	expected := 8410
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

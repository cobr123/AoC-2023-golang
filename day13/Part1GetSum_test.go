package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := "#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.\n\n#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part1GetSum(scanner)
	expected := 405
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

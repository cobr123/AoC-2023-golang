package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2GetSum(t *testing.T) {
	input := "#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.\n\n#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part2GetSum(scanner)
	expected := 400
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

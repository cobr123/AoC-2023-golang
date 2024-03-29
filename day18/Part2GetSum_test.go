package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2GetSum(t *testing.T) {
	input := "R 6 (#70c710)\nD 5 (#0dc571)\nL 2 (#5713f0)\nD 2 (#d2c081)\nR 2 (#59c680)\nD 2 (#411b91)\nL 5 (#8ceee2)\nU 2 (#caa173)\nL 1 (#1b58a2)\nU 2 (#caa171)\nR 2 (#7807d2)\nU 3 (#a77fa3)\nL 2 (#015232)\nU 2 (#7a21e3)"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part2GetSum(scanner)
	var expected int64 = 952408144115
	if got != expected {
		t.Errorf("\n got %d\nwant %d", got, expected)
	}
}

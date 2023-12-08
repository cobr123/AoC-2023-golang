package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2GetCnt6(t *testing.T) {
	input := "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part2GetCnt(scanner)
	expected := 6
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

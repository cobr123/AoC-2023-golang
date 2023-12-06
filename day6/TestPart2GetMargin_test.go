package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2GetMargin(t *testing.T) {
	input := "Time:      7  15   30\nDistance:  9  40  200"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part2GetMargin(scanner)
	expected := 71503
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

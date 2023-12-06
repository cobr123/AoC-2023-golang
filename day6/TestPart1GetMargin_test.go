package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetMargin(t *testing.T) {
	input := "Time:      7  15   30\nDistance:  9  40  200"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part1GetMargin(scanner)
	expected := 288
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

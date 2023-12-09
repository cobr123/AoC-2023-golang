package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2GetSum(t *testing.T) {
	input := "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part2GetSum(scanner)
	expected := 2
	if sum != expected {
		t.Errorf("got = %d; want %d", sum, expected)
	}
}

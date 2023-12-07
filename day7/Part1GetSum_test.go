package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner)
	expected := 6440
	if sum != expected {
		t.Errorf("got = %d; want %d", sum, expected)
	}
}

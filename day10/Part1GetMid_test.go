package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetMid(t *testing.T) {
	input := ".....\n.S-7.\n.|.|.\n.L-J.\n....."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part1GetMid(scanner)
	expected := 4
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

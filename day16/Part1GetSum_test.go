package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := ".|...\\....\n|.-.\\.....\n.....|-...\n........|.\n..........\n.........\\\n..../.\\\\..\n.-.-/..|..\n.|....-|.\\\n..//.|...."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part1CountEnergized(scanner)
	expected := 46
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner)
	expected := 1320
	if sum != expected {
		t.Errorf("got = %d; want %d", sum, expected)
	}
}

func TestPart1Hash(t *testing.T) {
	input := "HASH"
	got := Part1Hash(input)
	expected := 52
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

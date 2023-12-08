package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetCnt2(t *testing.T) {
	input := "RL\n\nAAA = (BBB, CCC)\nBBB = (DDD, EEE)\nCCC = (ZZZ, GGG)\nDDD = (DDD, DDD)\nEEE = (EEE, EEE)\nGGG = (GGG, GGG)\nZZZ = (ZZZ, ZZZ)"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part1GetCnt(scanner)
	expected := 2
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1GetCnt6(t *testing.T) {
	input := "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part1GetCnt(scanner)
	expected := 6
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

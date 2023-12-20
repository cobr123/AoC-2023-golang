package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := "2413432311323\n3215453535623\n3255245654254\n3446585845452\n4546657867536\n1438598798454\n4457876987766\n3637877979653\n4654967986887\n4564679986453\n1224686865563\n2546548887735\n4322674655533"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part1FindPath(scanner)
	expected := 102
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

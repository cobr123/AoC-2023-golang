package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner)
	expected := 114
	if sum != expected {
		t.Errorf("got = %d; want %d", sum, expected)
	}
}

func TestPart1FindNext1(t *testing.T) {
	input := "0 3 6 9 12 15"
	numbers := Part1ParseNumbers(input)
	got := Part1FindNext(numbers)
	expected := 18
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1FindNext2(t *testing.T) {
	input := "1 3 6 10 15 21"
	numbers := Part1ParseNumbers(input)
	got := Part1FindNext(numbers)
	expected := 28
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1FindNext3(t *testing.T) {
	input := "10 13 16 21 30 45"
	numbers := Part1ParseNumbers(input)
	got := Part1FindNext(numbers)
	expected := 68
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

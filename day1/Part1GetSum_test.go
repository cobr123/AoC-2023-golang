package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part1GetSum(scanner)
	expected := 142
	if sum != expected {
		t.Errorf("Part1GetSum = %d; want %d", sum, expected)
	}
}

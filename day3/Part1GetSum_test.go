package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetSum(t *testing.T) {
	input := "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	arr := Part1GetLines(scanner)
	sum := Part1GetSum(arr)
	expected := 4361
	if sum != expected {
		t.Errorf("Part1 = %d; want %d", sum, expected)
	}
}

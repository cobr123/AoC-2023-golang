package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2GetSum(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	sum := Part2GetSum(scanner)
	expected := 281
	if sum != expected {
		t.Errorf("Part2GetSum = %d; want %d", sum, expected)
	}
}

func TestParseDigitsExample1(t *testing.T) {
	first, last := Part2ParseDigits("two1nine")
	expectedFirst := 2
	if first != expectedFirst {
		t.Errorf("Part2ParseDigits = %d; want %d", first, expectedFirst)
	}
	expectedLast := 9
	if last != expectedLast {
		t.Errorf("Part2ParseDigits = %d; want %d", last, expectedLast)
	}
}

func TestParseDigitsExample2(t *testing.T) {
	first, last := Part2ParseDigits("eightwothree")
	expectedFirst := 8
	if first != expectedFirst {
		t.Errorf("Part2ParseDigits = %d; want %d", first, expectedFirst)
	}
	expectedLast := 3
	if last != expectedLast {
		t.Errorf("Part2ParseDigits = %d; want %d", last, expectedLast)
	}
}

func TestParseDigitsExample3(t *testing.T) {
	first, last := Part2ParseDigits("abcone2threexyz")
	expectedFirst := 1
	if first != expectedFirst {
		t.Errorf("Part2ParseDigits = %d; want %d", first, expectedFirst)
	}
	expectedLast := 3
	if last != expectedLast {
		t.Errorf("Part2ParseDigits = %d; want %d", last, expectedLast)
	}
}

func TestParseDigitsExample4(t *testing.T) {
	first, last := Part2ParseDigits("xtwone3four")
	expectedFirst := 2
	if first != expectedFirst {
		t.Errorf("Part2ParseDigits = %d; want %d", first, expectedFirst)
	}
	expectedLast := 4
	if last != expectedLast {
		t.Errorf("Part2ParseDigits = %d; want %d", last, expectedLast)
	}
}

func TestParseDigitsExample5(t *testing.T) {
	first, last := Part2ParseDigits("4nineeightseven2")
	expectedFirst := 4
	if first != expectedFirst {
		t.Errorf("Part2ParseDigits = %d; want %d", first, expectedFirst)
	}
	expectedLast := 2
	if last != expectedLast {
		t.Errorf("Part2ParseDigits = %d; want %d", last, expectedLast)
	}
}

func TestParseDigitsExample6(t *testing.T) {
	first, last := Part2ParseDigits("zoneight234")
	expectedFirst := 1
	if first != expectedFirst {
		t.Errorf("Part2ParseDigits = %d; want %d", first, expectedFirst)
	}
	expectedLast := 4
	if last != expectedLast {
		t.Errorf("Part2ParseDigits = %d; want %d", last, expectedLast)
	}
}

func TestParseDigitsExample7(t *testing.T) {
	first, last := Part2ParseDigits("7pqrstsixteen")
	expectedFirst := 7
	if first != expectedFirst {
		t.Errorf("Part2ParseDigits = %d; want %d", first, expectedFirst)
	}
	expectedLast := 6
	if last != expectedLast {
		t.Errorf("Part2ParseDigits = %d; want %d", last, expectedLast)
	}
}

func TestParseDigitsExample8(t *testing.T) {
	first, last := Part2ParseDigits("eighthree")
	expectedFirst := 8
	if first != expectedFirst {
		t.Errorf("Part2ParseDigits = %d; want %d", first, expectedFirst)
	}
	expectedLast := 3
	if last != expectedLast {
		t.Errorf("Part2ParseDigits = %d; want %d", last, expectedLast)
	}
}

func TestParseDigitsExample9(t *testing.T) {
	first, last := Part2ParseDigits("sevenine")
	expectedFirst := 7
	if first != expectedFirst {
		t.Errorf("Part2ParseDigits = %d; want %d", first, expectedFirst)
	}
	expectedLast := 9
	if last != expectedLast {
		t.Errorf("Part2ParseDigits = %d; want %d", last, expectedLast)
	}
}

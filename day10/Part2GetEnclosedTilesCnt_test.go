package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart2GetEnclosedTilesCnt1(t *testing.T) {
	input := "...........\n.S-------7.\n.|F-----7|.\n.||.....||.\n.||.....||.\n.|L-7.F-J|.\n.|..|.|..|.\n.L--J.L--J.\n..........."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part2GetEnclosedTilesCnt(scanner)
	expected := 4
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart2GetEnclosedTilesCnt2(t *testing.T) {
	input := "..........\n.S------7.\n.|F----7|.\n.||....||.\n.||....||.\n.|L-7F-J|.\n.|..||..|.\n.L--JL--J.\n.........."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part2GetEnclosedTilesCnt(scanner)
	expected := 4
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart2GetEnclosedTilesCnt3(t *testing.T) {
	input := ".F----7F7F7F7F-7....\n.|F--7||||||||FJ....\n.||.FJ||||||||L7....\nFJL7L7LJLJ||LJ.L-7..\nL--J.L7...LJS7F-7L7.\n....F-J..F7FJ|L7L7L7\n....L7.F7||L7|.L7L7|\n.....|FJLJ|FJ|F7|.LJ\n....FJL-7.||.||||...\n....L---J.LJ.LJLJ..."
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part2GetEnclosedTilesCnt(scanner)
	expected := 8
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart2GetEnclosedTilesCnt4(t *testing.T) {
	input := "FF7FSF7F7F7F7F7F---7\nL|LJ||||||||||||F--J\nFL-7LJLJ||||||LJL-77\nF--JF--7||LJLJ7F7FJ-\nL---JF-JLJ.||-FJLJJ7\n|F|F-JF---7F7-L7L|7|\n|FFJF7L7F-JF7|JL---7\n7-L-JL7||F7|L7F-7F7|\nL.L7LFJ|||||FJL7||LJ\nL7JLJL-JLJLJL--JLJ.L"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part2GetEnclosedTilesCnt(scanner)
	expected := 10
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

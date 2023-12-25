package main

import (
	"bufio"
	"bytes"
	"testing"
)

func TestPart1GetMult(t *testing.T) {
	input := "jqt: rhn xhk nvd\nrsh: frs pzl lsr\nxhk: hfx\ncmg: qnr nvd lhk bvb\nrhn: xhk bvb hfx\nbvb: xhk hfx\npzl: lsr hfx nvd\nqnr: nvd\nntq: jqt hfx bvb xhk\nnvd: lhk\nlsr: lhk\nrzs: qnr cmg lsr rsh\nfrs: qnr lhk lsr"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	got := Part1GetMult(scanner)
	expected := 54
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

func TestPart1CountPaths(t *testing.T) {
	input := "jqt: rhn xhk nvd\nrsh: frs pzl lsr\nxhk: hfx\ncmg: qnr nvd lhk bvb\nrhn: xhk bvb hfx\nbvb: xhk hfx\npzl: lsr hfx nvd\nqnr: nvd\nntq: jqt hfx bvb xhk\nnvd: lhk\nlsr: lhk\nrzs: qnr cmg lsr rsh\nfrs: qnr lhk lsr"
	b := bytes.NewReader([]byte(input))
	scanner := bufio.NewScanner(b)
	links := Part1GetLinks(scanner)
	link1 := Link{"pzl", "hfx"}
	link2 := Link{"cmg", "bvb"}
	link3 := Link{"jqt", "nvd"}
	got := Part1CountPaths(links, link1, link2, link3)
	expected := 54
	if got != expected {
		t.Errorf("got = %d; want %d", got, expected)
	}
}

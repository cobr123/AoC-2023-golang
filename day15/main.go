package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := Part1GetSum(scanner)
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := Part2GetSum(scanner)
	fmt.Println(sum)
}

func Part1GetSum(scanner *bufio.Scanner) int {
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, ",")
		for _, part := range parts {
			sum += Part1Hash(part)
		}
	}
	return sum
}

func Part1Hash(s string) int {
	result := 0
	for _, ch := range s {
		result += int(ch)
		result *= 17
		result %= 256
	}
	return result
}

type Box struct {
	items []BoxItem
}

type BoxItem struct {
	label       string
	focalLength int
}

func Part2GetSum(scanner *bufio.Scanner) int {
	boxes := Part2GetBoxes(scanner)
	sum := 0
	for i, box := range boxes {
		sum += Part2GetFocusingPower(i, box)
	}
	return sum
}

func Part2GetFocusingPower(idx int, box Box) int {
	focusingPower := 0
	for i, item := range box.items {
		focusingPower += (idx + 1) * (i + 1) * item.focalLength
	}
	return focusingPower
}

func Part2GetBoxes(scanner *bufio.Scanner) [256]Box {
	boxes := [256]Box{}
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, ",")
		for _, part := range parts {
			if strings.Contains(part, "=") {
				labelAndFocalLength := strings.Split(part, "=")
				label := labelAndFocalLength[0]
				focalLength, err := strconv.Atoi(labelAndFocalLength[1])
				if err != nil {
					panic(err)
				}
				labelHash := Part1Hash(label)
				found := false
				for i, item := range boxes[labelHash].items {
					if item.label == label {
						found = true
						boxes[labelHash].items[i].focalLength = focalLength
					}
				}
				if !found {
					boxes[labelHash].items = append(boxes[labelHash].items, BoxItem{label, focalLength})
				}
			} else {
				labelAndNothing := strings.Split(part, "-")
				label := labelAndNothing[0]
				labelHash := Part1Hash(label)
				boxes[labelHash].items = slices.DeleteFunc(boxes[labelHash].items, func(item BoxItem) bool {
					return item.label == label
				})
			}
		}
	}
	return boxes
}

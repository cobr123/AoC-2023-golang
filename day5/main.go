package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

func part1() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := Part1GetMin(scanner)
	fmt.Println(sum)
}

type CategoryMap struct {
	to     string
	ranges []CategoryRanges
}

type CategoryRanges struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func Part1GetMin(scanner *bufio.Scanner) int {
	seeds := []int{}
	categoryMaps := map[string]CategoryMap{}
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "seeds:") {
			nameAndNumbers := strings.Split(s, ":")
			numbers := strings.Split(strings.Trim(nameAndNumbers[1], " "), " ")
			for _, item := range numbers {
				n, err := strconv.Atoi(item)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, n)
			}
			scanner.Scan()
		} else if s != "" {
			catAndNumbers := strings.Split(s, ":")
			catAndMap := strings.Split(catAndNumbers[0], " ")
			catFromTo := strings.Split(catAndMap[0], "-")
			catFrom := catFromTo[0]
			catTo := catFromTo[2]
			categoryRanges := parseCategoryRanges(scanner)
			categoryMaps[catFrom] = CategoryMap{catTo, categoryRanges}
		} else {
			panic(s)
		}

	}
	minLoc := -1
	for _, item := range seeds {
		loc := findLocation(item, "seed", categoryMaps)
		if minLoc < 0 || loc < minLoc {
			minLoc = loc
		}
	}
	return minLoc
}

func findLocation(sourceItem int, category string, categoryMaps map[string]CategoryMap) int {
	for categoryMap, ok := categoryMaps[category]; ok; categoryMap, ok = categoryMaps[category] {
		for _, r := range categoryMap.ranges {
			if sourceItem >= r.sourceRangeStart && sourceItem < r.sourceRangeStart+r.rangeLength {
				delta := sourceItem - r.sourceRangeStart
				sourceItem = r.destinationRangeStart + delta
				break
			}
		}
		category = categoryMap.to
	}
	return sourceItem
}

func parseCategoryRanges(scanner *bufio.Scanner) []CategoryRanges {
	categoryRanges := []CategoryRanges{}
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			return categoryRanges
		}
		numbers := strings.Split(s, " ")
		destinationRangeStart, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		sourceRangeStart, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		rangeLength, err := strconv.Atoi(numbers[2])
		if err != nil {
			panic(err)
		}
		categoryRanges = append(categoryRanges, CategoryRanges{destinationRangeStart, sourceRangeStart, rangeLength})
	}
	return categoryRanges
}

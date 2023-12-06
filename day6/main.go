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
	part2()
}

func part1() {
	f, err := os.Open("./input1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := Part1GetMargin(scanner)
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("./input2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := Part1GetMargin(scanner)
	fmt.Println(sum)
}

func Part1GetMargin(scanner *bufio.Scanner) int {
	times := []int{}
	distances := []int{}
	for scanner.Scan() {
		s := strings.ReplaceAll(scanner.Text(), "  ", " ")
		for strings.Contains(s, "  ") {
			s = strings.ReplaceAll(s, "  ", " ")
		}
		if strings.HasPrefix(s, "Time:") {
			nameAndNumbers := strings.Split(s, ":")
			numbers := strings.Split(strings.Trim(nameAndNumbers[1], " "), " ")
			for _, item := range numbers {
				n, err := strconv.Atoi(item)
				if err != nil {
					panic(err)
				}
				times = append(times, n)
			}
		} else if strings.HasPrefix(s, "Distance:") {
			nameAndNumbers := strings.Split(s, ":")
			numbers := strings.Split(strings.Trim(nameAndNumbers[1], " "), " ")
			for _, item := range numbers {
				n, err := strconv.Atoi(item)
				if err != nil {
					panic(err)
				}
				distances = append(distances, n)
			}
		} else {
			panic(s)
		}

	}
	margin := 1
	for i := 0; i < len(times); i++ {
		margin *= getNumberOfWays(times[i], distances[i])
	}
	return margin
}

func getNumberOfWays(time int, distance int) int {
	cnt := 0

	for hold := 1; hold < time; hold++ {
		if (time-hold)*hold > distance {
			cnt += 1
		}
	}
	return cnt
}

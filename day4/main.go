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
		game := Part1ParseCards(s)
		sum += Part1CompareCards(game)
	}
	return sum
}

func Part2GetSum(scanner *bufio.Scanner) int {
	games := []Game{}
	for scanner.Scan() {
		s := scanner.Text()
		games = append(games, Part1ParseCards(s))
	}
	for i, game := range games {
		winCnt := Part2CompareCards(game)
		cnt := games[i].cnt
		for k := i + 1; k <= i+winCnt; k++ {
			if k >= len(games) {
				break
			}
			games[k].cnt += cnt
		}
	}
	sum := 0
	for _, game := range games {
		sum += game.cnt
	}
	return sum
}

type Game struct {
	winningCards map[int]struct{}
	cardsOnHand  map[int]struct{}
	cnt          int
}

func Part1ParseCards(str string) Game {
	winningCards := make(map[int]struct{})
	cardsOnHand := make(map[int]struct{})

	gameAndCards := strings.Split(strings.ReplaceAll(str, "  ", " "), ":")
	winningAndOnHandCards := strings.Split(strings.Trim(gameAndCards[1], " "), "|")
	cardsL := strings.Split(strings.Trim(winningAndOnHandCards[0], " "), " ")
	cardsR := strings.Split(strings.Trim(winningAndOnHandCards[1], " "), " ")

	for _, item := range cardsL {
		n, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		winningCards[n] = struct{}{}
	}
	for _, item := range cardsR {
		n, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		cardsOnHand[n] = struct{}{}
	}

	return Game{winningCards, cardsOnHand, 1}
}

func Part1CompareCards(game Game) int {
	sum := 0
	for item := range game.winningCards {
		if _, ok := game.cardsOnHand[item]; ok {
			if sum == 0 {
				sum = 1
			} else {
				sum *= 2
			}
		}
	}

	return sum
}

func Part2CompareCards(game Game) int {
	cnt := 0
	for item := range game.winningCards {
		if _, ok := game.cardsOnHand[item]; ok {
			cnt += 1
		}
	}

	return cnt
}

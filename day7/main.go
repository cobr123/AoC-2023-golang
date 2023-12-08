package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type HandType int8

const (
	HighCard     HandType = iota // where all cards' labels are distinct: 23456
	OnePair                      // where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	TwoPair                      // where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	ThreeOfAKind                 // where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	FullHouse                    // where three cards have the same label, and the remaining two cards share a different label: 23332
	FourOfAKind                  // where four cards have the same label and one card has a different label: AA8AA
	FiveOfAKind                  // where all five cards have the same label: AAAAA
)

type CardV1 int8

const (
	c2 CardV1 = iota
	c3
	c4
	c5
	c6
	c7
	c8
	c9
	T
	J
	Q
	K
	A
)

type HandV1 struct {
	cards    []CardV1
	handType HandType
	bid      int
}

func Part1GetSum(scanner *bufio.Scanner) int {
	hands := []HandV1{}
	for scanner.Scan() {
		s := scanner.Text()
		hands = append(hands, Part1ParseHand(s))
	}
	sort.Slice(hands, func(i, j int) bool {
		compare := hands[i].handType - hands[j].handType
		if compare == 0 {
			cardsI := hands[i].cards
			cardsJ := hands[j].cards
			for k := 0; k < len(cardsJ); k++ {
				compare := cardsI[k] - cardsJ[k]
				if compare != 0 {
					return compare < 0
				}
			}
			return false
		} else {
			return compare < 0
		}
	})
	sum := 0
	for rank, hand := range hands {
		sum += (rank + 1) * hand.bid
	}
	return sum
}

func Part1ParseHand(s string) HandV1 {
	cardsAndBid := strings.Split(s, " ")
	cards := []CardV1{}
	for _, ch := range cardsAndBid[0] {
		switch ch {
		case '2':
			cards = append(cards, c2)
		case '3':
			cards = append(cards, c3)
		case '4':
			cards = append(cards, c4)
		case '5':
			cards = append(cards, c5)
		case '6':
			cards = append(cards, c6)
		case '7':
			cards = append(cards, c7)
		case '8':
			cards = append(cards, c8)
		case '9':
			cards = append(cards, c9)
		case 'T':
			cards = append(cards, T)
		case 'J':
			cards = append(cards, J)
		case 'Q':
			cards = append(cards, Q)
		case 'K':
			cards = append(cards, K)
		case 'A':
			cards = append(cards, A)
		default:
			panic("undefined card " + string(ch))
		}
	}
	bid, err := strconv.Atoi(cardsAndBid[1])
	if err != nil {
		panic(err)
	}
	handType := Part1ParseHandType(cards)
	return HandV1{cards, handType, bid}
}

func Part1ParseHandType(cards []CardV1) HandType {
	m := map[CardV1]int{}
	for _, card := range cards {
		m[card]++
	}
	if len(m) == 1 {
		return FiveOfAKind
	} else if len(m) == 5 {
		return HighCard
	} else {
		if len(m) == 2 {
			for _, v := range m {
				if v == 1 || v == 4 {
					return FourOfAKind
				} else if v == 2 || v == 3 {
					return FullHouse
				}
			}
		}
		if len(m) == 3 {
			for _, v := range m {
				if v == 3 {
					return ThreeOfAKind
				} else if v == 2 {
					return TwoPair
				}
			}
		}
		if len(m) == 4 {
			for _, v := range m {
				if v == 2 {
					return OnePair
				}
			}
		}
	}
	panic("hand type not hound")
}

type CardV2 int8

const (
	dJ CardV2 = iota
	d2
	d3
	d4
	d5
	d6
	d7
	d8
	d9
	dT
	dQ
	dK
	dA
)

type HandV2 struct {
	cards    []CardV2
	handType HandType
	bid      int
}

func Part2GetSum(scanner *bufio.Scanner) int {
	hands := []HandV2{}
	for scanner.Scan() {
		s := scanner.Text()
		hands = append(hands, Part2ParseHand(s))
	}
	sort.Slice(hands, func(i, j int) bool {
		compare := hands[i].handType - hands[j].handType
		if compare == 0 {
			cardsI := hands[i].cards
			cardsJ := hands[j].cards
			for k := 0; k < len(cardsJ); k++ {
				compare := cardsI[k] - cardsJ[k]
				if compare != 0 {
					return compare < 0
				}
			}
			return false
		} else {
			return compare < 0
		}
	})
	sum := 0
	for rank, hand := range hands {
		sum += (rank + 1) * hand.bid
	}
	return sum
}

func Part2ParseHand(s string) HandV2 {
	cardsAndBid := strings.Split(s, " ")
	cards := []CardV2{}
	for _, ch := range cardsAndBid[0] {
		switch ch {
		case '2':
			cards = append(cards, d2)
		case '3':
			cards = append(cards, d3)
		case '4':
			cards = append(cards, d4)
		case '5':
			cards = append(cards, d5)
		case '6':
			cards = append(cards, d6)
		case '7':
			cards = append(cards, d7)
		case '8':
			cards = append(cards, d8)
		case '9':
			cards = append(cards, d9)
		case 'T':
			cards = append(cards, dT)
		case 'J':
			cards = append(cards, dJ)
		case 'Q':
			cards = append(cards, dQ)
		case 'K':
			cards = append(cards, dK)
		case 'A':
			cards = append(cards, dA)
		default:
			panic("undefined card " + string(ch))
		}
	}
	bid, err := strconv.Atoi(cardsAndBid[1])
	if err != nil {
		panic(err)
	}
	handType := Part2ParseHandType(cards)
	return HandV2{cards, handType, bid}
}

func Part2ParseHandType(cards []CardV2) HandType {
	m := map[CardV2]int{}
	for _, card := range cards {
		m[card]++
	}
	if len(m) > 1 {
		if cnt, ok := m[dJ]; ok {
			delete(m, dJ)
			left := [][2]int{}
			for k, v := range m {
				left = append(left, [2]int{int(k), v})
			}
			sort.Slice(left, func(i, j int) bool {
				compare := left[i][1] - left[j][1]
				return compare > 0
			})
			m[CardV2(left[0][0])] += cnt
		}
	}
	if len(m) == 1 {
		return FiveOfAKind
	} else if len(m) == 5 {
		return HighCard
	} else {
		if len(m) == 2 {
			for _, v := range m {
				if v == 1 || v == 4 {
					return FourOfAKind
				} else if v == 2 || v == 3 {
					return FullHouse
				}
			}
		}
		if len(m) == 3 {
			for _, v := range m {
				if v == 3 {
					return ThreeOfAKind
				} else if v == 2 {
					return TwoPair
				}
			}
		}
		if len(m) == 4 {
			for _, v := range m {
				if v == 2 {
					return OnePair
				}
			}
		}
	}
	panic("hand type not hound")
}

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

type Card int8

const (
	c2 Card = iota
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

type Hand struct {
	cards    []Card
	handType HandType
	bid      int
}

func Part1GetSum(scanner *bufio.Scanner) int {
	hands := []Hand{}
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

func Part1ParseHand(s string) Hand {
	cardsAndBid := strings.Split(s, " ")
	cards := []Card{}
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
	return Hand{cards, handType, bid}
}

func Part1ParseHandType(cards []Card) HandType {
	m := map[Card]int{}
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

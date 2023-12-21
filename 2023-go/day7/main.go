package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file_name := "test.txt"

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(data), "\n")
	hands := parseInput(input)

	winnings := solve(hands)
	fmt.Println()
	fmt.Println(winnings)
}

func parseNums(s string) []int {
	tmp := strings.Split(s, " ")
	nums := make([]int, 0)

	for _, num := range tmp {
		if num == "" {
			continue
		}
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
}

func parseInput(lines []string) []hand {
	hands := make([]hand, len(lines))

	for i, line := range lines {
		tmp := strings.Split(line, " ")
		cards := tmp[0]
		bid := parseNums(tmp[1])[0]
		hands[i] = hand{
			cards:       cards,
			bid:         bid,
			rankingType: getHandType(getCards(cards)),
		}
	}
	return hands
}

type hand struct {
	cards       string
	bid         int
	rankingType pokerHand
}

type pokerHand int
type pokerCard int

const (
	highCard pokerHand = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func getCards(hand string) []pokerCard {
	cards := make([]pokerCard, len(pokerCardMap))

	for _, c := range hand {
		for i, r := range pokerCardMap {
			if c == r {
				cards[i] += 1
			}
		}
	}

	return cards
}

func getHandType(hand []pokerCard) pokerHand {
	max, max2 := pokerCard(0), pokerCard(0)

	for _, v := range hand[1:] {
		if v >= max {
			max2 = max
			max = v
		} else if v >= max2 {
			max2 = v
		}
	}

	max += hand[0]

	switch max {
	case 5:
		return fiveOfAKind
	case 4:
		return fourOfAKind
	case 3:
		if max2 == 2 {
			return fullHouse
		}
		return threeOfAKind
	case 2:
		if max2 == 2 {
			return twoPair
		}
		return onePair
	default:
		return highCard
	}
}

// var pokerCardMap = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var pokerCardMap = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func solve(hands []hand) int {
	totalWinnings := 0

	slices.SortFunc(hands, sortHands)

	for i, hand := range hands {
		// fmt.Printf("%s, type: %d, bid: %d \n", hand.cards, hand.rankingType, hand.bid)

		totalWinnings += (i + 1) * hand.bid
	}
	return totalWinnings
}

func sortHands(a, b hand) int {
	diff := a.rankingType - b.rankingType
	if diff != 0 {
		return int(diff)
	}
	return compareHighCards(a, b)
}

func compareHighCards(a, b hand) int {
	for i := 0; i < len(a.cards); i++ {
		diff := getCardValue(a.cards[i]) - getCardValue(b.cards[i])
		if diff == 0 {
			continue
		}
		return diff
	}
	return 0
}

func getCardValue(card byte) int {
	for i, r := range pokerCardMap {
		if rune(card) == r {
			return i
		}
	}
	return 0
}

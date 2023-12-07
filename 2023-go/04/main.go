package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	index   int
	winning []int
	nums    []int
}

// const SEPERATOR_POS = 5

const SEPERATOR_POS = 10

func main() {
	file_name := os.Args[1]

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	cards := parseInput(input)

	// sum := solveOne(cards)
	sum := solveTwo(cards)
	fmt.Println(sum)

}

func parseInput(input []string) []Card {
	re := regexp.MustCompile(`\d+`)
	cards := []Card{}

	for i, line := range input {
		card := Card{}
		card.index = i

		tmp := strings.Split(line, ": ")
		numString := tmp[1]

		tmp = re.FindAllString(numString, -1)

		for i, s := range tmp {
			n, err := strconv.Atoi(s)
			if err != nil {
				continue
			}
			if i < SEPERATOR_POS {
				card.winning = append(card.winning, n)
			} else {
				card.nums = append(card.nums, n)
			}
		}
		cards = append(cards, card)
	}
	return cards
}

func solveTwo(cards []Card) int {
	total := 0
	stack := make([]int, 0)

	for _, card := range cards {
		matches := getMatches(card)
		for i := card.index + 1; i <= card.index+matches && i < len(cards); i++ {
			stack = append(stack, i)
		}
		total += 1
	}

	top := 0

	for {
		if len(stack) < 1 {
			return total
		}
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		total += 1

		card := cards[top]
		matches := getMatches(card)
		for i := card.index + 1; i <= card.index+matches && i < len(cards); i++ {
			stack = append(stack, i)
		}
	}
}

func solveOne(cards []Card) int {
	var pointMap = []int{0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	sum := 0

	for _, card := range cards {
		matches := getMatches(card)
		sum += pointMap[matches]
	}

	return sum
}

func getMatches(card Card) int {
	matches := 0
	for _, num := range card.winning {
		if slices.Contains(card.nums, num) {
			matches += 1
		}
	}
	// fmt.Printf("%+v \t%d\n", card, matches)
	return matches
}

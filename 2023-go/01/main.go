package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file_name := os.Args[1]

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}
	// fmt.Print(string(data))

	sum := 0
	for _, line := range strings.Split(string(data), "\n") {
		sum += getNum(line)
	}
	fmt.Println(sum)

}

// getNum parses a Line of input and gets the first and last digit in it.
// It then return the whole number.
func getNum(s string) int {
	re := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)

	// go doesn't support lookahead regex
	// so we need to handle overlapping digit string manually
	s = strings.ReplaceAll(s, "one", "onee")
	s = strings.ReplaceAll(s, "two", "twoo")
	s = strings.ReplaceAll(s, "three", "threee")
	s = strings.ReplaceAll(s, "five", "fivee")
	s = strings.ReplaceAll(s, "seven", "sevenn")
	s = strings.ReplaceAll(s, "eight", "eightt")
	s = strings.ReplaceAll(s, "nine", "ninee")

	matches := re.FindAllString(s, -1)

	value := parseDigit(matches[0])*10 + parseDigit(matches[len(matches)-1])
	fmt.Printf("%s: %d\n", s, value)
	return value
}

// parseDigit takes in a digit 0-9 or digit string 'one' to 'nine'.
// returns int value of the digit
func parseDigit(s string) int {
	m := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	value, err := strconv.Atoi(s)
	if err == nil {
		return value
	}

	for i, v := range m {
		if s == v {
			return i + 1
		}
	}
	return 0
}

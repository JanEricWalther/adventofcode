package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_name := os.Args[1]

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	sum := solveOne(string(data))
	fmt.Println("Part One: ", sum)
	sum = solveTwo(string(data))
	fmt.Println("Part Two: ", sum)
}

func solveOne(data string) int {
	sum := 0
	for _, line := range strings.Split(string(data), ",") {
		start, end := parse(line)
		sum += getInvalidDouble(start, end)
	}
	return sum
}

func solveTwo(data string) int {
	sum := 0
	for _, line := range strings.Split(string(data), ",") {
		start, end := parse(line)
		sum += getInvalid(start, end)
	}
	return sum
}

func parse(s string) (start, end int) {
	if len(s) < 1 {
		return 0, 0
	}

	arr := strings.Split(s, "-")
	if len(arr) != 2 {
		return 0, 0
	}

	start, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, 0
	}
	end, err = strconv.Atoi(arr[1])
	if err != nil {
		return 0, 0
	}
	return start, end
}

func getInvalidDouble(start, end int) int {
	sum := 0

	for i := start; i <= end; i++ {
		s := strconv.Itoa(i)

		s1 := s[:len(s)/2]
		s2 := s[len(s)/2:]

		if s1 == s2 {
			sum += i
		}
	}
	return sum
}

func getInvalid(start, end int) int {
	sum := 0

	for i := start; i <= end; i++ {
		sum += testId(i)
	}
	return sum
}

func testId(id int) int {
	s := strconv.Itoa(id)

outer:
	for length := 1; length <= len(s)/2; length++ {
		if len(s)%length != 0 {
			continue
		}
		curr := 0

		for curr+2*length <= len(s) {
			if s[curr:curr+length] != s[curr+length:curr+2*length] {
				continue outer
			}
			curr += length
		}
		// fmt.Println("Found invalid ID:", id)
		return id
	}
	return 0
}

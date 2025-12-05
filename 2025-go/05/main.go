package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

func main() {
	file_name := os.Args[1]

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}

	sum := solveOne(string(data))
	fmt.Println(sum)
	sum = solveTwo(string(data))
	fmt.Println(sum)

}

func solveOne(data string) int {
	input := strings.Split(data, "\n\n")
	fresh := parseRanges(input[0])

	sum := 0
	for _, line := range strings.Split(input[1], "\n") {
		id := parseId(line)
		sum += checkId(id, fresh)
	}

	return sum

}

func solveTwo(data string) int {
	return 0

}

func parseRanges(s string) [][2]int {
	var arr [][2]int

	for _, line := range strings.Split(s, "\n") {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		arr = append(arr, [2]int{start, end})
	}

	return arr
}

func parseId(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

func checkId(id int, ranges [][2]int) int {
	for _, r := range ranges {
		if id >= r[0] && id <= r[1] {
			return 1
		}
	}
	return 0
}

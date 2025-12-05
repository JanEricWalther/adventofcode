package main

import (
	"fmt"
	"os"
	"slices"
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
	input := strings.Split(data, "\n\n")[0]
	fresh := parseRanges(input)

	return freshIdCount(fresh)
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

// func freshIdCount(ranges [][2]int) int {
// 	m := make(map[int]bool)
// 	for _, r := range ranges {
// 		for i := r[0]; i <= r[1]; i++ {
// 			m[i] = true
// 		}
// 	}
// 	return len(m)
// }

func freshIdCount(ranges [][2]int) int {
	slices.SortFunc(ranges, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	sum := 0
	current := 0
	for _, r := range ranges {
		count := r[1] - r[0] + 1
		sum += count
		if r[0] <= current {
			diff := current - r[0] + 1
			if diff > count {
				diff = count
			}
			sum -= diff
		}
		if r[1] > current {
			current = r[1]
		}

		if DEBUG {
			fmt.Printf("Range: %v, sum: %d, current: %d\n", r, sum, current)
		}
	}
	return sum
}

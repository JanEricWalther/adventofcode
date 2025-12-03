package main

import (
	"fmt"
	"os"
	"strings"
)

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
	sum := 0
	for _, line := range strings.Split(string(data), "\n") {
		arr := parse(line)

		i, v10 := findLargestHighDigit(arr)
		v1 := findLargestLowDigit(arr, i)

		sum += v10*10 + v1

	}
	return sum
}

func solveTwo(data string) int {
	return 0
}

func parse(s string) []int {
	var arr []int
	for _, ch := range s {
		arr = append(arr, int(ch-'0'))
	}
	return arr
}

func findLargestHighDigit(arr []int) (int, int) {
	max := -1
	index := -1

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}

	return index, max
}

func findLargestLowDigit(arr []int, start int) int {
	max := -1

	for i := start + 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

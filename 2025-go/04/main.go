package main

import (
	"fmt"
	"os"
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
	sum := 0
	grid := parse(data)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				// check neighbors
				neighbors := 0
				directions := [][2]int{
					{-1, 0}, {1, 0}, {0, -1}, {0, 1},
					{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
				}
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[i]) {
						if grid[ni][nj] == 1 {
							neighbors++
						}
					}
				}
				if neighbors < 4 {
					sum++
				}

				if DEBUG {
					fmt.Printf("%d", neighbors)
				}
			}
		}
		if DEBUG {
			fmt.Printf("\n")
		}
	}

	return sum
}

func solveTwo(data string) int {
	return 0
}

func parse(s string) [][]int {
	var arr [][]int

	for _, line := range strings.Split(s, "\n") {
		var row []int
		for _, ch := range line {
			if ch == '@' {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		arr = append(arr, row)
	}
	return arr
}

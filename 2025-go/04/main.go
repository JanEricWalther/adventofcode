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
	grid := parse(data)

	return getValid(grid)
}

func solveTwo(data string) int {
	grid := parse(data)

	sum := 0

	for {
		removed := getValid(grid)
		if removed == 0 {
			return sum
		}
		sum += removed
	}
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

func getValid(grid [][]int) int {
	removable := make([][2]int, 0)
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
					removable = append(removable, [2]int{i, j})
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
	sum := len(removable)
	// remove accessible
	for _, pos := range removable {
		grid[pos[0]][pos[1]] = 0
	}

	return sum
}

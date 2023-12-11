package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file_name := "06/test.txt"
	file_name := "06/input.txt"

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")
	races := parseInput2(input[0], input[1])

	waysToWin := solve(races)
	fmt.Println(waysToWin)
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

func _parseInput(timeString, distanceString string) []race {
	tmp := strings.Split(timeString, ": ")[1]
	tmp = strings.TrimSpace(tmp)
	times := parseNums(tmp)
	tmp = strings.Split(distanceString, ": ")[1]
	tmp = strings.TrimSpace(tmp)
	distances := parseNums(tmp)
	races := make([]race, len(times))

	for i := range times {
		race := race{
			time:     times[i],
			distance: distances[i],
		}
		races[i] = race
	}
	return races
}

func parseInput2(timeString, distanceString string) []race {
	tmp := strings.Split(timeString, ": ")[1]
	tmp = strings.ReplaceAll(tmp, " ", "")
	times := parseNums(tmp)
	tmp = strings.Split(distanceString, ": ")[1]
	tmp = strings.ReplaceAll(tmp, " ", "")
	distances := parseNums(tmp)

	return []race{
		{
			time:     times[0],
			distance: distances[0],
		},
	}

}

type race struct {
	time     int
	distance int
}

func solve(races []race) int {
	waysToWin := 1

	for _, race := range races {
		ways := 0
		speed := 0
		for sec := 0; sec < race.time; sec += 1 {
			dist := speed * (race.time - sec)
			if dist > race.distance {
				ways += 1
			}
			speed += 1
		}
		waysToWin *= ways
	}

	return waysToWin
}

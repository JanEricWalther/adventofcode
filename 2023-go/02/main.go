package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	MAX_RED   = 12
	MAX_GREEN = 13
	MAX_BLUE  = 14
)

func main() {
	file_name := os.Args[1]

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}
	sum := 0

	for _, line := range strings.Split(string(data), "\n") {
		sum += getPower(line)
	}
	fmt.Println(sum)
}

type state struct {
	minRed   int
	minGreen int
	minBlue  int
}

func getPower(s string) int {
	game := strings.Split(s, ": ")[1]
	re := regexp.MustCompile(`\d+ (blue|red|green)`)
	draws := strings.Split(game, ";")

	state := state{
		minRed:   0,
		minGreen: 0,
		minBlue:  0,
	}

	for _, draw := range draws {
		colors := re.FindAllString(draw, -1)
		for _, colorDraw := range colors {

			tmp := strings.Split(colorDraw, " ")
			num, _ := strconv.Atoi(tmp[0])
			color := tmp[1]
			checkMinOfColor(color, num, &state)
		}
	}
	return state.minRed * state.minGreen * state.minBlue
}

func checkMinOfColor(color string, num int, state *state) {
	switch color {
	case "red":
		if num > state.minRed {
			state.minRed = num
		}
	case "green":
		if num > state.minGreen {
			state.minGreen = num
		}
	case "blue":
		if num > state.minBlue {
			state.minBlue = num
		}
	}
}

func checkGame(s string) int {
	re := regexp.MustCompile(`Game \d+`)
	id := strings.Split(re.FindString(s), " ")[1]
	game := strings.Split(s, ": ")[1]

	re = regexp.MustCompile(`\d+ (blue|red|green)`)
	draws := strings.Split(game, ";")

	isValidGame := true
	for _, draw := range draws {
		balls := re.FindAllString(draw, -1)
		for _, ball := range balls {
			tmp := strings.Split(ball, " ")
			num, _ := strconv.Atoi(tmp[0])
			color := tmp[1]
			isValidGame = isValidGame && checkColor(num, color)
		}
	}
	if !isValidGame {
		return 0
	}

	fmt.Println(id)

	value, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return value
}

func checkColor(num int, color string) bool {
	// fmt.Printf("%d %s\n", num, color)
	switch color {
	case "red":
		if num > MAX_RED {
			return false
		}
	case "green":
		if num > MAX_GREEN {
			return false
		}
	case "blue":
		if num > MAX_BLUE {
			return false
		}
	}
	return true
}

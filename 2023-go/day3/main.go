package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file_name := "input.txt"

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}
	input := strings.Split(string(data), "\n")

	nums := getValidNumbers(input)
	gears := getGearNumbers(input)

	fmt.Println(getSum(nums))
	fmt.Println(getSum(gears))
}

func getValidNumbers(m []string) []int {
	nums := []int{}

	current := 0
	firstPos := -1
	lastPos := -1

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if isDigit(m[y][x]) {
				current = current*10 + getDigit(m[y][x])

				if firstPos < 0 {
					firstPos = x
				}
				lastPos = x
				if x != len(m[y])-1 {
					continue
				}
			}
			if firstPos < 0 {
				continue
			}

			if checkValid(m, x, y, firstPos, lastPos) {
				nums = append(nums, current)
			}

			current = 0
			firstPos = -1
		}
	}
	return nums
}

func checkValid(m []string, x, y, firstPos, lastPos int) bool {
	// left
	if firstPos > 0 && m[y][firstPos-1] != '.' {
		return true
	}
	// right
	if lastPos < len(m[y])-1 && m[y][lastPos+1] != '.' {
		return true
	}
	// up
	if y > 0 {
		for i := firstPos; i <= lastPos; i++ {
			if m[y-1][i] != '.' {
				return true
			}
		}
	}
	// down
	if y < len(m)-1 {
		for i := firstPos; i <= lastPos; i++ {
			if m[y+1][i] != '.' {
				return true
			}
		}
	}
	// up left
	if y > 0 && firstPos > 0 && m[y-1][firstPos-1] != '.' {
		return true
	}
	// up right
	if y > 0 && lastPos < len(m[y])-1 && m[y-1][lastPos+1] != '.' {
		return true
	}
	// down left
	if y < len(m)-1 && firstPos > 0 && m[y+1][firstPos-1] != '.' {
		return true
	}
	// down right
	if y < len(m)-1 && lastPos < len(m[y])-1 && m[y+1][lastPos+1] != '.' {
		return true
	}
	return false
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func getDigit(b byte) int {
	return int(b - '0')
}

func getSum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return sum
}

type number struct {
	value int
	xPos  int
	yPos  int
	gearX int
	gearY int
}

func getGearNumbers(m []string) []int {
	numbers := []number{}

	current := 0
	firstPos := -1
	lastPos := -1

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if isDigit(m[y][x]) {
				current = current*10 + getDigit(m[y][x])

				if firstPos < 0 {
					firstPos = x
				}
				lastPos = x
				if x != len(m[y])-1 {
					continue
				}
			}
			if firstPos < 0 {
				continue
			}

			if n := checkValidGearNumber(m, x, y, firstPos, lastPos); n != nil {
				(*n).value = current
				numbers = append(numbers, *n)
			}

			current = 0
			firstPos = -1
		}
	}

	return getValidGears(numbers)
}

func checkValidGearNumber(m []string, x, y, firstPos, lastPos int) *number {
	// left
	if firstPos > 0 && m[y][firstPos-1] == '*' {
		return &number{
			value: 0,
			xPos:  firstPos,
			yPos:  y,
			gearX: firstPos - 1,
			gearY: y,
		}
	}
	// right
	if lastPos < len(m[y])-1 && m[y][lastPos+1] != '.' {
		return &number{
			value: 0,
			xPos:  firstPos,
			yPos:  y,
			gearX: lastPos + 1,
			gearY: y,
		}
	}
	// up
	if y > 0 {
		for i := firstPos; i <= lastPos; i++ {
			if m[y-1][i] != '.' {
				return &number{
					value: 0,
					xPos:  firstPos,
					yPos:  y,
					gearX: i,
					gearY: y - 1,
				}
			}
		}
	}
	// down
	if y < len(m)-1 {
		for i := firstPos; i <= lastPos; i++ {
			if m[y+1][i] != '.' {
				return &number{
					value: 0,
					xPos:  firstPos,
					yPos:  y,
					gearX: i,
					gearY: y + 1,
				}
			}
		}
	}
	// up left
	if y > 0 && firstPos > 0 && m[y-1][firstPos-1] != '.' {
		return &number{
			value: 0,
			xPos:  firstPos,
			yPos:  y,
			gearX: firstPos - 1,
			gearY: y - 1,
		}
	}
	// up right
	if y > 0 && lastPos < len(m[y])-1 && m[y-1][lastPos+1] != '.' {
		return &number{
			value: 0,
			xPos:  firstPos,
			yPos:  y,
			gearX: lastPos + 1,
			gearY: y - 1,
		}
	}
	// down left
	if y < len(m)-1 && firstPos > 0 && m[y+1][firstPos-1] != '.' {
		return &number{
			value: 0,
			xPos:  firstPos,
			yPos:  y,
			gearX: firstPos - 1,
			gearY: y + 1,
		}
	}
	// down right
	if y < len(m)-1 && lastPos < len(m[y])-1 && m[y+1][lastPos+1] != '.' {
		return &number{
			value: 0,
			xPos:  firstPos,
			yPos:  y,
			gearX: lastPos + 1,
			gearY: y + 1,
		}
	}
	return nil
}

func getValidGears(n []number) []int {
	gears := []int{}

	for i := 0; i < len(n); i++ {
		for j := i + 1; j < len(n); j++ {
			if i == j {
				continue
			}
			if n[i].gearX != n[j].gearX || n[i].gearY != n[j].gearY {
				continue
			}
			gears = append(gears, n[i].value*n[j].value)
		}
	}

	return gears
}

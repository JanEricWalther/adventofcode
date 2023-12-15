package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lens struct {
	label    string
	focalLen int
}

func main() {
	// file_name := "15/test.txt"
	file_name := "15/input.txt"

	data, err := os.ReadFile(file_name)

	if err != nil {
		panic(err)
	}
	// fmt.Print(string(data))

	input := strings.Split(string(data), ",")

	// solveOne(input)
	solveTwo(input)
}

func parseNums(line string, sep string) []int {
	nums := []int{}

	for _, n := range strings.Split(line, sep) {
		d, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		nums = append(nums, d)
	}
	return nums
}

func solveOne(lines []string) {
	sum := 0
	for _, line := range lines {
		sum += hash(line)
	}

	fmt.Println(sum)
}

func solveTwo(lines []string) {
	boxes := make([][]lens, 256)
	for i := range boxes {
		boxes[i] = make([]lens, 0)
	}

	for _, line := range lines {
		if strings.Contains(line, "-") {
			label := line[:len(line)-1]
			i := hash(label)

			j := 0
			for ; j < len(boxes[i]); j++ {
				if boxes[i][j].label == label {
					boxes[i] = append(boxes[i][:j], boxes[i][j+1:]...)
					break
				}
			}
		} else {
			tmp := strings.Split(line, "=")
			label := tmp[0]
			focalLen, err := strconv.Atoi(tmp[1])
			if err != nil {
				panic(err)
			}
			i := hash(label)
			found := false
			for j := range boxes[i] {
				if boxes[i][j].label == label {
					boxes[i][j].focalLen = focalLen
					found = true
					break
				}
			}
			if !found {
				boxes[i] = append(boxes[i], lens{label: label, focalLen: focalLen})
			}
		}
		// fmt.Printf("\nAfter: %s\n", line)
		// prettyPrint(boxes)
	}

	fmt.Println(getFocusPower(boxes))
}

func getFocusPower(boxes [][]lens) int {
	focusPower := 0
	for n, box := range boxes {
		for sl, lens := range box {
			focusPower += (n + 1) * (sl + 1) * lens.focalLen
		}
	}
	return focusPower
}

func hash(s string) int {
	currVal := 0

	for _, c := range s {
		currVal += int(c)
		currVal *= 17
		currVal %= 256
	}
	return currVal
}

func prettyPrint(boxes [][]lens) {
	for i, box := range boxes {
		if len(box) == 0 {
			continue
		}
		fmt.Printf("Box %d: %v\n", i, box)
	}
}

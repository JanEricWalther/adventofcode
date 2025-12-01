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

	countClicks := false
	if len(os.Args) == 3 && (os.Args[2] == "-c" || os.Args[2] == "--clicks") {
		countClicks = true
	}

	dial := Dial{pos: 50, zeroCount: 0, allClicks: countClicks, debug: false}

	for _, line := range strings.Split(string(data), "\n") {
		dial.exec(line)
	}
	fmt.Println(dial.zeroCount)

}

type Dial struct {
	pos       int
	zeroCount int
	allClicks bool
	debug     bool
}

func (d *Dial) exec(s string) {
	val := parseRotation(s)
	if d.allClicks {
		zeros := abs(val) / 100
		d.zeroCount += zeros
	}
	val %= 100

	d.pos += val
	if d.pos == 0 || d.pos == 100 {
		d.pos = 0
		d.zeroCount++
	}

	if d.pos < 0 {
		if d.allClicks && d.pos-val != 0 {
			d.zeroCount++
		}
		d.pos += 100
	}
	if d.pos > 100 {
		if d.allClicks && d.pos-val != 0 {
			d.zeroCount++
		}
		d.pos -= 100
	}

	if d.debug {
		fmt.Println("After", s, "pos:", d.pos, "zeros:", d.zeroCount)
	}
}

func parseRotation(s string) int {
	if len(s) < 1 {
		return 0
	}

	dir := 1
	if s[0] == 'L' {
		dir = -1
	}

	value, err := strconv.Atoi(s[1:])
	if err == nil {
		return dir * value
	}
	return 0
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

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
	// fmt.Print(string(data))

	dial := Dial{pos: 50, zeroCount: 0}

	for _, line := range strings.Split(string(data), "\n") {
		dial.exec(line)
	}
	fmt.Println(dial.zeroCount)

}

type Dial struct {
	pos       int
	zeroCount int
}

func (d *Dial) exec(s string) {
	val := parseRotation(s)
	val %= 100

	d.pos += val
	if d.pos < 0 {
		d.pos += 100
	}
	if d.pos >= 100 {
		d.pos -= 100
	}

	if d.pos == 0 {
		d.zeroCount++
	}
	fmt.Println(s, "->", d.pos)
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

package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	data := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	expected := 13
	result := solveOne(data)
	if result != expected {
		t.Errorf("Part One: expected %d, got %d", expected, result)
	}
}

func TestPart2(t *testing.T) {
	data := `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

	expected := 43
	result := solveTwo(data)
	if result != expected {
		t.Errorf("Part Two: expected %d, got %d", expected, result)
	}
}

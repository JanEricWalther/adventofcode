package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	data := `987654321111111
811111111111119
234234234234278
818181911112111`

	expected := 357
	result := solveOne(data)
	if result != expected {
		t.Errorf("Part One: expected %d, got %d", expected, result)
	}
}

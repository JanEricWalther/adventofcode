package main

import "testing"

func TestPart1(t *testing.T) {
	data := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

	expected := 3
	result := solveOne(data)
	if result != expected {
		t.Errorf("Part One: expected %d, got %d", expected, result)
	}
}

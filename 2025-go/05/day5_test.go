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

func TestPart2(t *testing.T) {

	tests := []struct {
		data     string
		expected int
	}{
		{
			data: `3-5
10-14
16-20
12-18

1
`,
			expected: 14,
		},
		{
			data: `1-10
1-2
1-2

9
`,
			expected: 10,
		},
		{
			data: `63012323001376-68526035085810
354523654291459-355141743321025


4
`,
			expected: 5513712084435 + 618089029567,
		},
	}

	for _, test := range tests {
		result := solveTwo(test.data)
		if result != test.expected {
			t.Errorf("Part Two: expected %d, got %d", test.expected, result)
		}
	}
}

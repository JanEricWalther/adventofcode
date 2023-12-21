package main

import "testing"

func TestGetValidNumbers(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected []int
	}{
		{
			desc: "testing getValidNumbers",
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: []int{467, 35, 633, 617, 592, 755, 664, 598},
		},
		{
			desc: "getValidNumbers edge case hehe",
			input: []string{
				"467..114..",
				"...*......",
				"..35...633",
				"..........",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: []int{467, 35, 617, 592, 755, 664, 598},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			actual := getValidNumbers(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
			for i := range actual {
				if actual[i] != tc.expected[i] {
					t.Errorf("expected %v, got %v", tc.expected, actual)
				}
			}
		})
	}
}

func TestGetSum(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []int
		expected int
	}{
		{
			desc:     "",
			input:    []int{467, 35, 633, 617, 592, 755, 664, 598},
			expected: 4361,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			actual := getSum(tc.input)
			if actual != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestGetGearNumbers(t *testing.T) {
	tcs := []struct {
		desc     string
		input    []string
		expected []int
	}{
		{
			desc: "testing getGearNumbers",
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: []int{16345, 451490},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			actual := getGearNumbers(tc.input)
			if len(actual) != len(tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
			for i := range actual {
				if actual[i] != tc.expected[i] {
					t.Errorf("expected %v, got %v", tc.expected, actual)
				}
			}
		})
	}
}

package main

import "testing"

func TestGetHandType(t *testing.T) {
	tsc := []struct {
		desc     string
		input    string
		expected pokerHand
	}{
		{
			desc:     "Five of a kind",
			input:    "AAAAA",
			expected: fiveOfAKind,
		},
		{
			desc:     "Four of a kind",
			input:    "AAAA2",
			expected: fourOfAKind,
		},
		{
			desc:     "Full house",
			input:    "AA222",
			expected: fullHouse,
		},
		{
			desc:     "Three of a kind",
			input:    "AAA23",
			expected: threeOfAKind,
		},
		{
			desc:     "Two pair",
			input:    "AA223",
			expected: twoPair,
		},
		{
			desc:     "One pair",
			input:    "AA234",
			expected: onePair,
		},
		{
			desc:     "High card",
			input:    "A2345",
			expected: highCard,
		},
	}

	for _, tc := range tsc {
		t.Run(tc.desc, func(t *testing.T) {
			cards := getCards(tc.input)
			actual := getHandType(cards)
			if actual != tc.expected {
				t.Errorf("hand: %s\n expected %v, got %v", tc.input, tc.expected, actual)
			}
		})
	}
}

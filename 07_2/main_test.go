package main

import (
	"reflect"
	"testing"
)

func TestParseSetJoker(t *testing.T) {
	testcases := []struct{
		input string
		expected Set
	}{
		{
			"JJJJJ",
			Set{
				FiveOfAKind,
				"JJJJJ",
			},
		},
		// four jokers
		{
			"JJJJK",
			Set{
				FiveOfAKind,
				"JJJJK",
			},
		},

		// 3 jokers
		{
			"JJJKK",
			Set{
				FiveOfAKind,
				"JJJKK",
			},
		},
		{
			"JJJT9",
			Set{
				FourOfAKind,
				"JJJT9",
			},
		},

		// 2 jokers
		{
			"JJQQQ",
			Set{
				FiveOfAKind,
				"JJQQQ",
			},
		},
		{
			"JJAA8",
			Set{
				FourOfAKind,
				"JJAA8",
			},
		},
		{
			"JJ765",
			Set{
				ThreeOfAKind,
				"JJ765",
			},
		},

		// 1 joker
		{
			"J2222",
			Set{
				FiveOfAKind,
				"J2222",
			},
		},
		{
			"J3334",
			Set{
				FourOfAKind,
				"J3334",
			},
		},
		{
			"J4455",
			Set{
				FullHouse,
				"J4455",
			},
		},
		{
			"J6678",
			Set{
				ThreeOfAKind,
				"J6678",
			},
		},
		{
			"J9TQK",
			Set{
				OnePair,
				"J9TQK",
			},
		},

		// no joker
		{
			"KA234",
			Set{
				HighCard,
				"KA234",
			},
		},
	}

	for _, tc := range testcases {
		result := ParseSet(tc.input)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("Expected %v, Got %v", tc.expected,  result)
		}
	}
}

func TestParseHand(t *testing.T) {
	testcases := []struct{
		input string
		expected Hand
	}{
		{
			"QQQJA 1234",
			Hand {
				set: Set{
					FourOfAKind,
					"QQQJA",
				},
				score: 1234,
			},
		},
	}

	for _, tc := range testcases {
		result := ParseHand(tc.input)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("Expected %v, Got %v", tc.expected,  result)
		}
	}
}

func TestParseSetJoker2(t *testing.T) {
	testcases := []struct{
		input string
		expected Set
	}{
		{
			"J4455",
			Set{
				FullHouse,
				"J4455",
			},
		},
	}

	for _, tc := range testcases {
		result := ParseSet(tc.input)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Fatalf("Expected %v, Got %v", tc.expected,  result)
		}
	}
}

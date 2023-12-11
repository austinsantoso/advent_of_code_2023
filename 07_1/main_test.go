package main

import (
	"reflect"
	"testing"
)

func TestParseSet(t *testing.T) {
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
		{
			"JJJJK",
			Set{
				FourOfAKind,
				"JJJJK",
			},
		},
		{
			"88899",
			Set{
				FullHouse,
				"88899",
			},
		},
		{
			"QQQJA",
			Set{
				ThreeOfAKind,
				"QQQJA",
			},
		},
		{
			"77665",
			Set{
				TwoPair,
				"77665",
			},
		},
		{
			"44329",
			Set{
				OnePair,
				"44329",
			},
		},
		{
			"AK964",
			Set{
				HighCard,
				"AK964",
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
					ThreeOfAKind,
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

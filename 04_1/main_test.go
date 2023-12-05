package main

import (
	"reflect"
	"testing"
)

func TestParseCardFromString(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	result, ok := parseCardFromString(input)

	if !ok {
		t.Fatalf("Failed to parse string")
	}

	expected := Card{
		id: 1,
		winningNumbers: []int{41, 48, 83, 86, 17},
		haveNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("not the same. Expected: %v, got %v", expected, result)
	}
}

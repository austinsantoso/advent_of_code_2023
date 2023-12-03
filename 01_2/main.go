package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println("Hello World")
	var out int

	for {
		if s, e := readLine(); e == nil {
			temp := 10*findFirstDigit(s) + findLastDigit(s)
			// fmt.Printf("Temp value: %d\n", temp)
			out += temp
		} else {
			break
		}
		// fmt.Printf("Processing %s\n", s)
	}

	fmt.Println(out)
}

func readLine() (string, error) {
	var a string
	_, err := fmt.Scanf("%s", &a)

	return a, err
}

func findFirstDigit(s string) int {
	outValue := -1
	outIndex := len(s)
	for k, v := range digits {
		if i := strings.Index(s, k); i >= 0 {
			if i < outIndex {
				outIndex = i
				outValue = v
			}
		}
	}
	return outValue
}

func findLastDigit(s string) int {
	outValue := -1
	outIndex := -1
	for k, v := range digits {
		if i := strings.LastIndex(s, k); i >= 0 {
			if i > outIndex {
				outIndex = i
				outValue = v
			}
		}
	}
	return outValue
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"0":     0,
}

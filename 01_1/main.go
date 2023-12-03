package main

import "errors"
import "fmt"

func main() {
	// fmt.Println("Hello World")
	var out int

	for {
		var temp int

		if s, e := readLine(); e == nil {
			if a, err := findFirstDigit(s); err == nil {
				temp = a * 10
			} else {
				break
			}
			if b, err := findLastDigit(s); err == nil {
				temp += b
			} else {
				break
			}
			out += temp
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

func findFirstDigit(s string) (int, error) {
	for i := range s {
		if isDigit(s[i]) {
			// fmt.Printf("found first digit: %c\n", s[i])
			return byteToDigit(s[i]), nil
		}
	}

	return 0, errors.New("No digit found")
}

func findLastDigit(s string) (int, error) {
	l := len(s)
	for i := range s {
		c := s[l-1-i]
		if isDigit(c) {
			// fmt.Printf("found last digit: %c\n", c)
			return byteToDigit(c), nil
		}
	}

	return 0, errors.New("No digit found")
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func byteToDigit(b byte) int {
	return int(b - '0')
}

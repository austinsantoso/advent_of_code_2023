package main

import (
	"bufio"
	"os"
	"testing"
)

func TestSolve(t *testing.T) {

	tests := []struct {
		grid [][]byte
		expected int
	}{
		{grid: [][]byte{[]byte{'.'}}, expected: 0},
	}

	for i, tc := range tests {
		var g Grid = Grid{tc.grid, len(tc.grid[0]), len(tc.grid)}
		got := solve(&g)
		if tc.expected != got {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.expected, got)
		}
	}

}

func GridFromFile(filepath string) *Grid {
	if file, err := os.Open(filepath); err == nil {
		return makeGrid(bufio.NewScanner(file))
	}
	return nil
}

func TestSolve_path(t *testing.T) {

	tests := []struct {
		filepath string
		expected int
	}{
		{filepath: "./in/1.in", expected: 4361},
		{filepath: "./in/2.in", expected: 0},
		{filepath: "./in/3.in", expected: 0},
		{filepath: "./in/4.in", expected: 1},
		{filepath: "./in/5.in", expected: 1234},
		{filepath: "./in/6.in", expected: 805},
		{filepath: "./in/7.in", expected: 102},
		{filepath: "./in/8.in", expected: 925},
		{filepath: "./in/9.in", expected: 0},
		{filepath: "./in/10.in", expected: 156},
		{filepath: "./in/11.in", expected: 2},
		{filepath: "./in/13.in", expected: 4},
	}

	for i, tc := range tests {

		var g = GridFromFile(tc.filepath)
		if g == nil {
			t.Fatalf("test %d: Failed to load file", i+1)
		}

		got := solve(g)
		if tc.expected != got {
			t.Fatalf("test %d: expected: %v, got: %v", i+1, tc.expected, got)
		}
	}

}

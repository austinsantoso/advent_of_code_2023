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
		{filepath: "./in/1.in", expected: 467835},
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

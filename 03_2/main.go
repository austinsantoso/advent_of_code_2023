package main

import "bufio"
import "fmt"
import "os"

func main() {
	grid := GridFromStdin()

	fmt.Println(solve(grid))
}

func GridFromStdin() *Grid {
	sc := bufio.NewScanner(os.Stdin)
	return makeGrid(sc)
}

func makeGrid(sc *bufio.Scanner) *Grid {
	var out [][]byte
	for sc.Scan() {
		scanned := sc.Bytes()
		var temp = make([]byte, len(scanned))
		copy(temp, scanned)
		out = append(out, temp)
	}

	return &Grid{
		out,
		len(out[0]),
		len(out),
	}
}

func solve(grid *Grid) int {
	sum := 0

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if grid.getByte(x, y) == '*' {
				sum += grid.getGearRatio(x, y)
			}
		}
	}

	return sum
}

type Grid struct {
	Grid   [][]byte
	Width  int
	Height int
}

func (g *Grid) String() string {

	var out = make([]byte, 0)

	for y := 0; y < g.Height; y++ {
		if y != 0 {
			out = append(out, '\n')
		}
		for x := 0; x < g.Width; x++ {
			if x != 0 {
				out = append(out, ' ')
			}
			out = append(out, g.Grid[y][x])
		}
	}

	return string(out)
}

func (g *Grid) getByte(x int, y int) byte {
	return g.Grid[y][x]
}

func (g *Grid) isDigit(x int, y int) bool {
	return g.getByte(x, y) >= '0' && g.getByte(x, y) <= '9'
}

func (g *Grid) isSymbol(x int, y int) bool {
	return !g.isDigit(x, y) && g.getByte(x, y) != '.'
}

func (g *Grid) safeIsSymbol(x int, y int) bool {
	return g.inGrid(x, y) && g.isSymbol(x, y)
}

func (g *Grid) inGrid(x int, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

type delta struct {
	x int
	y int
}

var deltas []*delta = []*delta{
	&delta{-1, -1},
	&delta{0, -1},
	&delta{1, -1},

	&delta{-1, 0},
	&delta{1, 0},

	&delta{-1, 1},
	&delta{0, 1},
	&delta{1, 1},
}

func (g *Grid) surroungedBySymbol(x int, y int) bool {
	for _, d := range deltas {
		if g.safeIsSymbol(x+d.x, y+d.y) {
			return true
		}
	}
	return false
}

type numberEntry struct {
	value int
	start int
	end int
}

func (g *Grid) getDigitsInRow(y int) []numberEntry {
	if y < 0 || y >= g.Height {
		return nil
	}
	//

	out := make([]numberEntry, 0)

	currentNumber:= 0
	startIndex := -1
	for x := 0; x < g.Width; x++ {
		if g.isDigit(x, y) {
			currentNumber = currentNumber*10 + int(g.getByte(x, y)-'0')
			if startIndex < 0 {
				startIndex = x
			}
			continue
		}
		// if not a digit
		//check if previous number can be added
		if startIndex > -1 {
			out = append(out, numberEntry{currentNumber, startIndex, x})
			startIndex = -1
		}
		currentNumber = 0
	}

	if currentNumber > 0 {
		out = append(out, numberEntry{currentNumber, startIndex, g.Width})
	}

	return out
}

func (g *Grid) getGearRatio(x int, y int) int {
	nums := make([]int, 0)
	// row above
	topRow := g.getDigitsInRow(y-1)
	for _, e := range topRow {
		if e.start <= x-1 && x-1 <= e.end-1 {
			nums = append(nums, e.value)
		} else if e.start <= x && x <= e.end-1 {
			nums = append(nums, e.value)
		} else if e.start <= x+1 && x+1 <= e.end-1 {
			nums = append(nums, e.value)
		}
	}

	// this row
	curRow := g.getDigitsInRow(y)
	for _, e := range curRow {
		if e.start <= x-1 && x-1 <= e.end-1 {
			nums = append(nums, e.value)
		} else if e.start <= x && x <= e.end-1 {
			nums = append(nums, e.value)
		} else if e.start <= x+1 && x+1 <= e.end-1 {
			nums = append(nums, e.value)
		}
	}

	// bottom row
	bottomRow := g.getDigitsInRow(y+1)
	for _, e := range bottomRow {
		if e.start <= x-1 && x-1 <= e.end-1 {
			nums = append(nums, e.value)
		} else if e.start <= x && x <= e.end-1 {
			nums = append(nums, e.value)
		} else if e.start <= x+1 && x+1 <= e.end-1 {
			nums = append(nums, e.value)
		}
	}

	if len(nums) == 2 {
		return nums[0] * nums[1]
	}
	return 0
}

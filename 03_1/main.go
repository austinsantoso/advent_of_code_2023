package main

import "bufio"
import "fmt"
import "os"

func main() {
	grid := GridFromStdin()
	// from top left to top right
	//fmt.Println(grid)

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

	fmt.Println(grid)

	temp := 0
	for y := 0; y < grid.Height; y++ {
		currentNumber := 0
		hasSymbol := false
		for x := 0; x < grid.Width; x++ {
			temp++
			//fmt.Println(currentNumber)
			//fmt.Println(hasSymbol)
			//fmt.Println()
			if grid.isDigit(x, y) {
				currentNumber = currentNumber*10 + int(grid.getByte(x, y)-'0')
				if !hasSymbol {
					// check surrounding
					hasSymbol = grid.surroungedBySymbol(x, y)
				}
				continue
			}

			// if not a digit

			//check if previous number can be added
			if hasSymbol {
				sum += currentNumber
				fmt.Println()
				fmt.Println("Adding:")
				fmt.Println(currentNumber)
				fmt.Printf("sum: %d\n", sum)
			}
			hasSymbol = false
			currentNumber = 0
		}

		// in case previous row last number not added
		if hasSymbol {
			sum += currentNumber
			fmt.Println()
			fmt.Println("Adding:")
			fmt.Println(currentNumber)
			fmt.Printf("sum: %d\n", sum)
		}
		hasSymbol = false
		currentNumber = 0
	}
	fmt.Println(temp)

	return sum
}

type Grid struct {
	Grid   [][]byte
	Width  int
	Height int
}

func (g *Grid) String() string {

	var out []byte = make([]byte, 0)

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

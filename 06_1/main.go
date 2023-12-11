package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/austinsantoso/advent_of_code_2023/util"
)

func main() {
	races := ParseRaceFromStdIn()
	out := 1
	for _, r := range races {
		out = out * r.WaysToWinRace()
	}

	fmt.Println(out)
}

type Race struct {
	Time int
	Distance int
}

func (r Race) WaysToWinRace() int {
	out := 0
	for i := 0; i <= r.Time; i++ {
		if r.canWinRace(i) {
			out++
		}
	}
	return out
}

func (r Race) canWinRace(t int) bool {
	speed := t
	time := r.Time - t
	return speed * time > r.Distance
}


func ParseRaceFromStdIn() []Race {
	sc := bufio.NewScanner(os.Stdin)

	out := make([]Race, 0)

	// scan times
	sc.Scan()
	times := sc.Text()
	c := util.NewParsingContext(times)
	if newContext, ok := util.ParseKeyword(c, "Time:"); !ok {
		panic("Failed to parse times")
	} else {
		c = newContext
		for !c.IsEnd() {
			if time, parseIntContext, parseIntOk := util.ParseInteger(c); !parseIntOk {
				panic("Failed to parse time")
			} else {
				out = append(out, Race{Time: time})
				c = parseIntContext
			}
		}
	}

	sc.Scan()
	distances := sc.Text()
	c = util.NewParsingContext(distances)
	if newContext, ok := util.ParseKeyword(c, "Distance:"); !ok {
		panic("Failed to parse Distances")
	} else {
		c = newContext
		i := 0
		for !c.IsEnd() {
			if distance, parseIntContext, parseIntOk := util.ParseInteger(c); !parseIntOk {
				panic("Failed to parse distance")
			} else {
				out[i].Distance = distance
				i++
				c = parseIntContext
			}
		}
	}

	return out
}


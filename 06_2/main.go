package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "github.com/austinsantoso/advent_of_code_2023/util"
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
	time := strings.TrimPrefix(times, "Time: ")
	time = strings.Trim(time, " ")
	time = strings.ReplaceAll(time, " ", "")
	timeInt, _ := strconv.Atoi(time)
	fmt.Println(time)

	// scan times
	sc.Scan()
	distances:= sc.Text()
	dist := strings.TrimPrefix(distances, "Distance: ")
	dist = strings.Trim(dist, " ")
	dist = strings.ReplaceAll(dist, " ", "")
	distInt, _ := strconv.Atoi(dist)
	fmt.Println(dist)

	out = append(out, Race{timeInt, distInt})

	return out
}


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	w, _ := WorldFromStdin()

	out := -1
	numSeeds := len(w.Seeds)
	for i := 0; i < numSeeds; i = i + 2 {
		for s := w.Seeds[i]; s < w.Seeds[i] + w.Seeds[i + 1]; s++ {
			if out < 0 {
				out = SeedToLocation(w, s)
			} else {
				out = min(out, SeedToLocation(w, s))
			}
		}
	}
	fmt.Println(out)
}

func SeedToLocation(w World, seed int) int {
	v := seed
	//fmt.Println()
	//fmt.Printf("Seed: %d\n", v)
	v = getValueOrSelf(*w.SeedToSoil, v)
	//fmt.Printf("Soil: %d\n", v)
	v = getValueOrSelf(*w.SoilToFertilizer, v)
	//fmt.Printf("FertilizerToWater: %d\n", v)
	v = getValueOrSelf(*w.FertilizerToWater, v)
	//fmt.Printf("Water: %d\n", v)
	v = getValueOrSelf(*w.WaterToLight, v)
	//fmt.Printf("Light: %d\n", v)
	v = getValueOrSelf(*w.LightToTemperature, v)
	//fmt.Printf("Temperature: %d\n", v)
	v = getValueOrSelf(*w.TemperatureToHumidity, v)
	//fmt.Printf("Humidity: %d\n", v)
	v = getValueOrSelf(*w.HumidityToLocation, v)
	//fmt.Printf("Location: %d\n", v)

	return v
}

type World struct {
	Seeds []int
	SeedToSoil *[]SourceDestinationRange
	SoilToFertilizer *[]SourceDestinationRange
	FertilizerToWater *[]SourceDestinationRange
	WaterToLight *[]SourceDestinationRange
	LightToTemperature *[]SourceDestinationRange
	TemperatureToHumidity *[]SourceDestinationRange
	HumidityToLocation *[]SourceDestinationRange
}

type SourceDestinationRange struct {
	DestinationStart int
	SourceStart int
	Range int
}

func (s SourceDestinationRange) isInRange(v int) bool {
	return s.SourceStart <= v && v < s.SourceStart + s.Range
}

func (s SourceDestinationRange) getValue(v int) int{
	if !s.isInRange(v) {
		return v
	}
	return s.DestinationStart + (v - s.SourceStart)
}

func NewWorld() World {
	seedToSoil:= make([]SourceDestinationRange, 0)
	soilToFertilizer:= make([]SourceDestinationRange, 0)
	fertilizerToWater:= make([]SourceDestinationRange, 0)
	waterToLight:= make([]SourceDestinationRange, 0)
	lightToTemperature:= make([]SourceDestinationRange, 0)
	temperatureToHumidity:= make([]SourceDestinationRange, 0)
	humidityToLocation:= make([]SourceDestinationRange, 0)

	return World{
		Seeds: make([]int, 0),
		SeedToSoil: &seedToSoil,
		SoilToFertilizer: &soilToFertilizer,
		FertilizerToWater: &fertilizerToWater,
		WaterToLight: &waterToLight,
		LightToTemperature: &lightToTemperature,
		TemperatureToHumidity: &temperatureToHumidity,
		HumidityToLocation: &humidityToLocation,
	}
}

func WorldFromStdin() (World, bool) {
	sc := bufio.NewScanner(os.Stdin)
	worldString := make([]string, 0)
	for sc.Scan() {
		worldString = append(worldString, sc.Text())
	}

	return parseWorld(worldString)
}

func getValueOrSelf(m []SourceDestinationRange, k int) int {
	for _, r := range m {
		if r.isInRange(k) {
			return r.getValue(k)
		}
	}
	return k
}


func parseWorld(worldString []string) (World, bool) {
	index := 0
	ok := true
	w := NewWorld()
	index, ok = ParseSeeds(worldString, index, &w)
	if !ok {
		panic("Failed to parse seeds")
	}
	index++ // skip new line
	fmt.Println("ParseSeeds")

	index, ok = ParseMap(worldString, index, &w.SeedToSoil)
	if !ok {
		panic("Failed to parse SeedToSoil")
	}
	index++
	fmt.Println("SeedToSoil")

	index, ok = ParseMap(worldString, index, &w.SoilToFertilizer)
	if !ok {
		panic("Failed to parse SoilToFertilizer")
	}
	index++
	fmt.Println("SoilToFertilizer")

	index, ok = ParseMap(worldString, index, &w.FertilizerToWater)
	if !ok {
		panic("Failed to parse FertilizerToWater")
	}
	index++
	fmt.Println("FertilizerToWater")

	index, ok = ParseMap(worldString, index, &w.WaterToLight)
	if !ok {
		panic("Failed to parse WaterToLight")
	}
	index++
	fmt.Println("WaterToLight")

	index, ok = ParseMap(worldString, index, &w.LightToTemperature)
	if !ok {
		panic("Failed to parse LightToTemperature")
	}
	index++
	fmt.Println("LightToTemperature")

	index, ok = ParseMap(worldString, index, &w.TemperatureToHumidity)
	if !ok {
		panic("Failed to parse TemperatureToHumidity")
	}
	index++
	fmt.Println("TemperatureToHumidity")

	index, ok = ParseMap(worldString, index, &w.HumidityToLocation)
	if !ok {
		panic("Failed to parse HumidityToLocation")
	}
	index++
	fmt.Println("HumidityToLocation")

	fmt.Println(w)
	return w, true
}

func ParseMap(worldString []string, index int, m **[]SourceDestinationRange) (int, bool) {
	newIndex := index
	newIndex++ // skip the description

	// parse 3 numbers
	ranges := make([]SourceDestinationRange, 0)
	for newIndex < len(worldString) && len(worldString[newIndex]) > 0 {
		values := strings.Split(worldString[newIndex], " ")
		if len(values) > 3 {
			panic("more than 3 values")
		}
		firstValue := StringToInt(values[0])
		secondValue := StringToInt(values[1])
		thirdValue := StringToInt(values[2])
		ranges = append(ranges, SourceDestinationRange{
			DestinationStart: firstValue,
			SourceStart: secondValue,
			Range: thirdValue,
		})
		*m = &ranges
		
		newIndex++
	}

	return newIndex, true
}

func ParseSeeds(worldString []string, index int, w *World) (int, bool) {
	newIndex := index
	c := NewParsingContext(worldString[newIndex])
	if newContext, ok := parseKeyword(c, "seeds:"); ok {
		c = newContext
	}
	for !c.isEnd() {
		if seed, newContext, ok := parseInteger(c); ok {
			c = newContext
			w.Seeds = append(w.Seeds, seed)
		}
	}

	newIndex++
	return newIndex, true
}

func parseKeyword(c *parsingContext, keyword string) (*parsingContext, bool) {
	newContext := c.clone()
	newContext.skipSpace()
	for i := range keyword {
		curChar := keyword[i]

		if newContext.isEnd() {
			return nil, false
		}

		if curChar != newContext.currentByte() {
			return nil, false
		}

		newContext.incrementIndex()
	}

	return newContext, true
}

func parseInteger(c *parsingContext) (int, *parsingContext, bool) {
	newContext := c.clone()
	newContext.skipSpace()
	if !newContext.isNumber() {
		return -1, nil, false
	}

	temp := 0
	for newContext.isNumber() {
		curNum := int(newContext.currentByte() - '0')
		newContext.incrementIndex()
		temp = temp * 10 + curNum
	}

	return temp, newContext, true
}


func NewParsingContext(s string) *parsingContext {
	return &parsingContext{
		s,
		len(s),
		0,
	}
}
type parsingContext struct {
	Input string
	Length int
	Index int
}

func (p *parsingContext) clone() *parsingContext {
	return &parsingContext{p.Input, p.Length, p.Index}
}

func (p *parsingContext) currentByte() byte {
	return p.Input[p.Index]
}

func (p *parsingContext) isEnd() bool {
	return p.Index >= p.Length
}

func (p *parsingContext) incrementIndex() {
	p.Index = p.Index+1
}

func (p *parsingContext) skipSpace() {
	for !p.isEnd() && p.isSpace() {
		p.incrementIndex()
	}
}

func (p *parsingContext) isSpace() bool {
	return p.currentByte() == ' '
}

func (p *parsingContext) isNumber() bool {
	return !p.isEnd() && p.currentByte() >= '0' && p.currentByte() <= '9'
}

func StringToInt(s string) int {
	temp := 0
	for i := range s {
		curNum := int(s[i] - '0')
		temp = temp * 10 + curNum
	}
	return temp
}

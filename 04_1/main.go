package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	cards := CardsFromStdin()
	out := 0
	for _, c := range cards {
		// fmt.Printf("Game %d has points %d\n", i, c.Points())
		out += c.Points()
	}

	fmt.Println(out)
}

func CardsFromStdin() []Card {
	sc := bufio.NewScanner(os.Stdin)
	out := make([]Card, 0)
	for sc.Scan() {
		scanned := sc.Text()
		if card, ok := parseCardFromString(scanned); ok {
			out = append(out, card)
		}
	}

	return out
}

type Card struct {
	id int
	winningNumbers []int
	haveNumbers []int
}

func (c Card) Points() int {
	out := 0
	for _, n := range c.haveNumbers {
		for _, w := range c.winningNumbers {
			if w == n {
				if out > 0 {
					out = out * 2
				} else {
					out = 1
				}
			}
		}
	}
	return out
}

func parseCardFromString(s string) (Card, bool) {
	c := parsingContext{s, len(s), 0}
	return parseCard(&c)
}

func parseCard(c *parsingContext) (Card, bool) {
	context := c.clone()

	out := Card{-1, make([]int, 0), make([]int, 0)}
	// skip Card keyword
	if newContext, ok := parseKeyword(context, "Card"); ok {
		context = newContext
	}
	// get ID
	if id, newContext, ok := parseInteger(context); ok {
		context = newContext
		out.id = id
	}

	if newContext, ok := parseKeyword(context, ":"); ok {
		context = newContext
	}
	// parse winning numbers
	for {
		if num, newContext, ok := parseInteger(context); ok {
			context = newContext
			out.winningNumbers = append(out.winningNumbers, num)
		}

		if newContext, ok := parseKeyword(context, "|"); ok {
			context = newContext
				break
		}
	}
	// parse owned numbers
	for {
		if num, newContext, ok := parseInteger(context); ok {
			context = newContext
			out.haveNumbers= append(out.haveNumbers, num)
		}

		// if ended escape
		if context.isEnd() {
			break
		}
	}

	return out, true
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

// Hello world
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

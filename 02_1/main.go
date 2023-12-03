package main

import "fmt"
import "bufio"
import "os"

func main() {
	fmt.Println("Hello World")

	in := bufio.NewReader(os.Stdin)

	out := 0

	for {
		if line, err := in.ReadString('\n'); err == nil {
			if len(line) < 2 {
				break // last newline or something
			}

			g := parseGame(line[:len(line)-1]) // to remove the ending newline
			if verifyGame(g) {
				out += g.id
			}
		} else {
			break
		}
	}

	fmt.Println(out)
}

func verifyGame(g Game) bool {
	maxSet := Bag{12, 13, 14}

	for _, s := range g.sets {
		if s.Red > maxSet.Red {
			return false
		}

		if s.Green > maxSet.Green {
			return false
		}

		if s.Blue > maxSet.Blue {
			return false
		}

	}

	return true
}

func readLine() (string, error) {
	var a string
	_, err := fmt.Scanf("%s", &a)

	return a, err
}

func parseGame(s string) Game {
	out := Game{}

	context := &parsingContext{s, len(s), 0}
	// parse game keyword
	if newContext, ok := parseKeyword(context, "Game"); ok {
		context = newContext
	}

	// parse game ID
	if id, newContext, ok := parseInteger(context); ok {
		out.id = id
		context = newContext
	}

	// skip the colon :
	if newContext, ok := parseKeyword(context, ":"); ok {
		context = newContext
	}

	// parse each game
	for {
		if b, newContext, ok := parseSet(context); ok {
			out.sets = append(out.sets, b)
			context = newContext
		}

		if context.isEnd() {
			break
		}
		if newContext, ok := parseKeyword(context, ";"); ok {
			context = newContext
		}
	}

	return out
}

func parseSet(c *parsingContext) (Bag, *parsingContext, bool) {
	context := c.clone()
	context.skipSpace()

	out := Bag{}
	// parse number
	for {
		num := 0
		if n, newContext, ok := parseInteger(context); ok {
			num = n
			context = newContext
		}

		// parse is it blue or red or green
		if newContext, ok := parseKeyword(context, "red"); ok {
			out.Red = num
			context = newContext
		} else if newContext, ok := parseKeyword(context, "green"); ok {
			out.Green = num
			context = newContext
		} else if newContext, ok := parseKeyword(context, "blue"); ok {
			out.Blue = num
			context = newContext
		}

		if newContext, ok := parseKeyword(context, ","); ok {
			context = newContext
			continue
		}

		if context.isEnd() {
			break
		}

		if newContext, ok := parseKeyword(context, ";"); ok {
			context = newContext
			break
		}
	}

	return out, context, true
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
	return p.currentByte() >= '0' && p.currentByte() <= '9'
}

type Game struct {
	id int
	sets []Bag
}


type Bag struct {
	Red int
	Green int
	Blue int
}


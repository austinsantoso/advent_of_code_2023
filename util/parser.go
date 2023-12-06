package util

func ParseKeyword(c *ParsingContext, keyword string) (*ParsingContext, bool) {
	newContext := c.Clone()
	newContext.SkipSpace()
	for i := range keyword {
		curChar := keyword[i]

		if newContext.IsEnd() {
			return nil, false
		}

		if curChar != newContext.CurrentByte() {
			return nil, false
		}

		newContext.IncrementIndex()
	}

	return newContext, true
}

func ParseInteger(c *ParsingContext) (int, *ParsingContext, bool) {
	newContext := c.Clone()
	newContext.SkipSpace()
	if !newContext.IsNumber() {
		return -1, nil, false
	}

	temp := 0
	for newContext.IsNumber() {
		curNum := int(newContext.CurrentByte() - '0')
		newContext.IncrementIndex()
		temp = temp * 10 + curNum
	}

	return temp, newContext, true
}

// Hello world
type ParsingContext struct {
	Input string
	Length int
	Index int
}

func NewParsingContext(s string) *ParsingContext {
	return &ParsingContext {
		Input: s,
		Length: len(s),
		Index: 0,
	}
}

func (p *ParsingContext) Clone() *ParsingContext {
	return &ParsingContext{p.Input, p.Length, p.Index}
}

func (p *ParsingContext) CurrentByte() byte {
	return p.Input[p.Index]
}

func (p *ParsingContext) IsEnd() bool {
	return p.Index >= p.Length
}

func (p *ParsingContext) IncrementIndex() {
	p.Index = p.Index+1
}

func (p *ParsingContext) SkipSpace() {
	for !p.IsEnd() && p.IsSpace() {
		p.IncrementIndex()
	}
}

func (p *ParsingContext) IsSpace() bool {
	return p.CurrentByte() == ' '
}

func (p *ParsingContext) IsNumber() bool {
	return !p.IsEnd() && p.CurrentByte() >= '0' && p.CurrentByte() <= '9'
}

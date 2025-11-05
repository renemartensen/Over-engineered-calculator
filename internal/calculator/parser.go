package calculator

import (
	"fmt"
	"strconv"
)

// AST

type Exp interface{}

type Literal struct {
	Value float64
}

type BinaryOp struct {
	Left     Exp
	Operator string
	Right    Exp
}

type UnaryOp struct {
	Operator string
	Operand  Exp
}

func (l *Literal) String() string {
	return fmt.Sprintf("%g", l.Value)
}

func (b *BinaryOp) String() string {
	return fmt.Sprintf("(%s %s %s)", b.Left, b.Operator, b.Right)
}

func (u *UnaryOp) String() string {
	return fmt.Sprintf("(%s%s)", u.Operator, u.Operand)
}

// _______________________________________________________________________
// Parser

type Parser struct {
	tokens []string
	pos    int
}

// look at the current token without consuming it
func (p *Parser) peek() string {
	if p.pos >= len(p.tokens) {
		return ""
	}
	return p.tokens[p.pos]
}

// consume and return the current token
func (p *Parser) next() string {
	token := p.peek()
	p.pos++
	return token
}

func (p *Parser) Parse() Exp {
	return p.parseAddSub()
}

func (p *Parser) parseAddSub() Exp {
	left := p.parseMulDiv()
	for {
		op := p.peek()
		if op == "+" || op == "-" {
			p.next()
			right := p.parseMulDiv()
			left = &BinaryOp{Left: left, Operator: op, Right: right}
		} else {
			break
		}
	}
	return left
}

func (p *Parser) parseMulDiv() Exp {
	left := p.parsePrimary()
	for {
		op := p.peek()
		if op == "*" || op == "/" {
			p.next()
			right := p.parsePrimary()
			left = &BinaryOp{Left: left, Operator: op, Right: right}
		} else {
			break
		}
	}
	return left
}

func (p *Parser) parsePrimary() Exp {
	token := p.peek()

	if token == "+" || token == "-" {
		p.next()
		return &UnaryOp{
			Operator: token,
			Operand:  p.parsePrimary(),
		}
	}

	if token == "(" {
		p.next()
		expr := p.parseAddSub()
		if p.peek() != ")" {
			panic("expected )")
		}
		p.next()
		return expr
	}

	p.next()
	val, err := strconv.ParseFloat(token, 64)
	if err != nil {
		panic(fmt.Sprintf("expected number, got %q", token))
	}
	return &Literal{Value: val}
}

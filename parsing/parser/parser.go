package parser

import (
	"fmt"
	"github.com/karlmoad/go_util_lib/generics/queue"
	"github.com/karlmoad/go_util_lib/parsing/dialect"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
	"strings"
)

type Parser struct {
	lex        *lexer.Lexer
	reg        *Registry
	pos        int
	depth      int
	dialect    dialect.Dialect
	eventQueue queue.Queue[ConditionHandler]
}

func NewParser(source string, dialect dialect.Dialect) *Parser {
	p := &Parser{
		pos:        0,
		reg:        newRegistry(),
		lex:        lexer.Tokenize(source, dialect),
		depth:      0,
		dialect:    dialect,
		eventQueue: queue.NewLIFOQueue[ConditionHandler](),
	}

	p.dialect.RegisterParsingHandlers(p.reg)

	return p
}

func (p *Parser) currentToken() lexer.Token {
	return p.lex.Tokens[p.pos]
}

func (p *Parser) advance() lexer.Token {
	token := p.currentToken()
	p.pos++
	return token
}

func (p *Parser) hasMoreTokens() bool {
	return p.pos < len(p.lex.Tokens)-1 && p.currentToken().Kind != lexer.EOF
}

func (p *Parser) peek(n int) lexer.Token {
	tPos := p.pos + n
	if tPos < 0 {
		tPos = 0
	} else {
		if tPos > len(p.lex.Tokens)-1 {
			tPos = len(p.lex.Tokens) - 1
		}
	}
	return p.lex.Tokens[tPos]
}

func (p *Parser) peekTokenStream(plusMinus int) []lexer.Token {
	sPos := p.pos - plusMinus
	ePos := p.pos + plusMinus + 1

	if sPos < 0 {
		sPos = 0
	}

	if ePos > len(p.lex.Tokens) {
		return p.lex.Tokens[sPos:]
	} else {
		return p.lex.Tokens[sPos:ePos]
	}
}

func (p *Parser) next() lexer.Token {
	return p.peek(1)
}

func (p *Parser) prev() lexer.Token {
	return p.peek(-1)
}

func (p *Parser) expect(expectedKind ...lexer.TokenKind) {
	if !p.currentToken().IsKindOf(expectedKind...) {
		kinds := make([]string, 0)
		for _, kind := range expectedKind {
			kinds = append(kinds, p.lex.TokenKindString(kind))
		}

		stream := p.errorContext()
		err := fmt.Sprintf("Expected %s but recieved %s:(%s) instead (index:%d)\n stream: %s\n",
			strings.Join(kinds, "|"),
			p.lex.TokenKindString(p.currentToken().Kind),
			p.currentToken().Value,
			p.pos,
			stream)

		panic(err)
	}
}

func (p *Parser) errorContext() string {
	stream := p.peekTokenStream(4)
	var buffer strings.Builder
	for _, token := range stream {
		buffer.WriteString(fmt.Sprintf("%s:[%s],", p.lex.TokenKindString(token.Kind), token.Value))
	}
	return fmt.Sprintf("Current: %s:(%s)[%d], Stream: %s\n", p.currentToken().Value, p.lex.TokenKindString(p.currentToken().Kind), p.pos, buffer.String())
}

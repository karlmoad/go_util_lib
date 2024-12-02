package parser

import (
	"fmt"
	"github.com/karlmoad/go_util_lib/generics/queue"
	"github.com/karlmoad/go_util_lib/parsing/ast"
	"github.com/karlmoad/go_util_lib/parsing/errors"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
	"log/slog"
	"strings"
)

type ParseCallback func(p *Parser) bool

type Parser struct {
	lex           *lexer.Lexer
	reg           *Registry
	pos           int
	depth         int
	callbackQueue queue.Queue[ParseCallback]
	logger        *slog.Logger
}

func NewParser(reg *Registry, lexer *lexer.Lexer) *Parser {
	p := &Parser{
		pos:           0,
		reg:           reg,
		lex:           lexer,
		depth:         0,
		callbackQueue: queue.NewLIFOQueue[ParseCallback](),
		logger:        slog.Default(),
	}
	return p
}

func (p *Parser) getRegistry() *Registry {
	return p.reg
}

func (p *Parser) Pos() int {
	return p.pos
}

func (p *Parser) GetLogger() *slog.Logger {
	return p.logger
}

func (p *Parser) CurrentToken() lexer.Token {
	return p.currentToken()
}

func (p *Parser) currentToken() lexer.Token {
	return p.lex.Tokens[p.pos]
}

func (p *Parser) Advance() lexer.Token {
	return p.advance()
}

func (p *Parser) advance() lexer.Token {
	token := p.currentToken()
	p.pos++
	return token
}

func (p *Parser) HasMoreTokens() bool {
	return p.hasMoreTokens()
}

func (p *Parser) hasMoreTokens() bool {
	return p.pos < len(p.lex.Tokens)-1 && p.currentToken().Kind != lexer.EOF
}

func (p *Parser) peek(n int) (lexer.Token, bool) {
	tPos := p.pos + n
	valid := true

	if tPos >= len(p.lex.Tokens) {
		tPos = len(p.lex.Tokens) - 1
		valid = false
	}

	if tPos < 0 {
		tPos = 0
		valid = false
	}

	return p.lex.Tokens[tPos], valid
}

func (p *Parser) PeekNext() (lexer.Token, bool) {
	return p.next()
}

func (p *Parser) next() (lexer.Token, bool) {
	return p.peek(1)
}

func (p *Parser) prev() (lexer.Token, bool) {
	return p.peek(-1)
}

func (p *Parser) Expect(kind ...lexer.TokenKind) error {
	return p.expect(kind...)
}

func (p *Parser) expect(expectedKind ...lexer.TokenKind) error {
	if !p.currentToken().IsKindOf(expectedKind...) {
		kinds := make([]string, 0)
		for _, kind := range expectedKind {
			kinds = append(kinds, p.lex.TokenKindString(kind))
		}
		err := fmt.Sprintf("Expected %s but recieved %s:(%s)",
			strings.Join(kinds, " or "),
			p.lex.TokenKindString(p.currentToken().Kind),
			p.currentToken().Value)

		return errors.NewUnexpectedTokenError(err)
	}
	return nil
}

func (p *Parser) errorContext() string {
	return p.lex.GetContext(p.pos-10, p.pos+10)
}

func (p *Parser) Parse(source string) ([]ast.Element, error) {
	if err := p.lex.Tokenize(source); err != nil {
		return nil, err
	}

	elem := make([]ast.Element, 0)

	for p.hasMoreTokens() {
		if p.currentToken().Kind == lexer.EOF {
			break
		}

		el, err := p.ProcessNextToken()
		if err != nil {
			return nil, err
		} else {
			if el != nil {
				elem = append(elem, el)
			} else {
				break
			}
		}
	}
	return elem, nil
}

func (p *Parser) ProcessNextToken() (ast.Element, error) {
	if p.evalCallbacks() {
		return nil, nil
	}

	p.depth++
	defer func() { p.depth-- }()
	handler := p.reg.evaluateConditions(p)
	if obj, valid, err := handler(p); valid {
		return obj, nil
	} else {
		if err != nil {
			return nil, errors.NewHandlerError(err.Error(), p.pos)
		} else {
			return nil, errors.NewHandlerError("parsing handler fault, invalid object returned", p.pos)
		}
	}
}

func (p *Parser) evalCallbacks() bool {
	if p.callbackQueue.Depth() > 0 {
		if funq, valid := p.callbackQueue.Current(); valid {
			if ret := funq(p); ret {
				return true
			}
		}
	}
	return p.reg.evaluateCallbacks(p)
}

func (p *Parser) PushCallback(cb ParseCallback) {
	p.callbackQueue.Enqueue(cb)
}

func (p *Parser) CurrentCallback() (ParseCallback, bool) {
	return p.callbackQueue.Current()
}

func (p *Parser) DequeueCurrentCallback() {
	p.callbackQueue.Dequeue()
}

func (p *Parser) Depth() int {
	return p.depth
}

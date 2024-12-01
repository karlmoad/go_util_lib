package parser

import (
	"fmt"
	"github.com/karlmoad/go_util_lib/generics/queue"
	"github.com/karlmoad/go_util_lib/parsing/ast"
	"github.com/karlmoad/go_util_lib/parsing/errors"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
	"log/slog"
	"math"
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

func (p *Parser) peek(n int) lexer.Token {
	tPos := p.pos + int(math.Abs(math.Inf(n)))

	if tPos >= len(p.lex.Tokens) {
		tPos = len(p.lex.Tokens) - 1
	}

	return p.lex.Tokens[tPos]
}

func (p *Parser) PeekNext() lexer.Token {
	return p.next()
}

func (p *Parser) next() lexer.Token {
	return p.peek(1)
}

func (p *Parser) prev() lexer.Token {
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

		stream := p.errorContext()
		err := fmt.Sprintf("Expected %s but recieved %s:(%s) instead (index:%d)\n stream: %s\n",
			strings.Join(kinds, "|"),
			p.lex.TokenKindString(p.currentToken().Kind),
			p.currentToken().Value,
			p.pos,
			stream)

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

		el, err := p.ParseNext()
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

func (p *Parser) ParseNext() (ast.Element, error) {
	//check if escape conditions are met, if so return nil
	if p.evalCallbacks() { // <-- eval callbacks
		return nil, nil // signaled callback encountered , return null, no error
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
	var ret bool
	//iterate exemption queue until empty or current == false
	if p.callbackQueue.Depth() > 0 {
		for {
			if funq, valid := p.callbackQueue.Current(); valid {
				if ret = funq(p); ret {
					p.callbackQueue.Dequeue()
				} else {
					break
				}
			} else {
				break
			}
		}
	} else {
		return true
	}
	return ret && p.callbackQueue.Depth() == 0
}

func (p *Parser) PushCallback(cb ParseCallback) {
	p.callbackQueue.Enqueue(cb)
}

func (p *Parser) ResetCallbacks() {
	p.callbackQueue.Clear()
}

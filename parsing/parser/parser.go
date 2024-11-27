package parser

import (
	"fmt"
	"github.com/karlmoad/go_util_lib/generics/queue"
	"github.com/karlmoad/go_util_lib/parsing"
	"github.com/karlmoad/go_util_lib/parsing/ast"
	"github.com/karlmoad/go_util_lib/parsing/dialect"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
	"math"
	"strings"
)

type ParseCallback func(p *Parser) bool

type Parser struct {
	lex           *lexer.Lexer
	reg           *Registry
	pos           int
	depth         int
	dialect       dialect.Dialect
	callbackQueue queue.Queue[ParseCallback]
}

func NewParser(source string, dialect dialect.Dialect) *Parser {
	p := &Parser{
		pos:           0,
		reg:           newRegistry(),
		lex:           lexer.NewLexer(source, dialect),
		depth:         0,
		dialect:       dialect,
		callbackQueue: queue.NewLIFOQueue[ParseCallback](),
	}
	p.dialect.RegisterParser(p.reg)
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
	tPos := p.pos + int(math.Abs(math.Inf(n)))

	if tPos >= len(p.lex.Tokens) {
		tPos = len(p.lex.Tokens) - 1
	}

	return p.lex.Tokens[tPos]
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
	return p.lex.GetContext(p.pos-10, p.pos+10)
}

func (p *Parser) Parse() ([]ast.ObjType, error) {
	if err := p.lex.Tokenize(); err != nil {
		return nil, err
	}

	objects := make([]ast.ObjType, 0)

	for p.hasMoreTokens() {
		if p.currentToken().Kind == lexer.EOF {
			break
		}

		obj, err := parseObject(p)
		if err != nil {
			return nil, err
		} else {
			if obj != nil {
				objects = append(objects, obj)
			} else {
				break
			}
		}
	}
	return objects, nil
}

func parseObject(p *Parser) (ast.ObjType, error) {
	//check if escape conditions are met, if so return nil
	if checkEscapeConditions(p) { // <-- eval callbacks
		return nil, nil // signal escape encountered by returning null, no error
	}

	p.depth++
	defer func() { p.depth-- }()
	handler := p.reg.evaluateConditions(p)
	if obj, valid := handler(p); valid {
		return obj, nil
	} else {
		// propagate error
		return nil, parsing.NewHandlerError("parsing handler fault, invalid object returned", p.pos)
	}
}

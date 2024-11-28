package parser

import (
	"github.com/karlmoad/go_util_lib/parsing/ast"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
)

type ParsingHandler func(p *Parser) (ast.ObjType, bool)
type Condition func(p *Parser) bool

func TokenTypeCondition(kind lexer.TokenKind) Condition {
	return func(p *Parser) bool {
		return p.currentToken().Kind == kind
	}
}

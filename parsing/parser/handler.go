package parser

import (
	"github.com/karlmoad/go_util_lib/parsing/ast"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
)

type ParsingHandler func(p *Parser) (ast.ObjType, bool)
type ConditionHandler func(p *Parser) bool

func TokenTypeConditionHandler(kind lexer.TokenKind) ConditionHandler {
	return func(p *Parser) bool {
		return p.currentToken().Kind == kind
	}
}

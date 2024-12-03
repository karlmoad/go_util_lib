package grammar

import (
	"github.com/karlmoad/go_util_lib/parsing/ast"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
)

const (
	RULE_ELEM ast.ElementKind = iota
	IDENTIFIER_ELEM
	STRING_ELEM
	BODY_ELEM
	SET_ELEM
	UNK_ELEM
	ALT_ELEM
)

type RuleExpr struct {
	Identifier lexer.Token
	Body       ast.Element
}

func (g RuleExpr) Expr() {}
func (g RuleExpr) Elem() ast.ElementMeta {
	return ast.InitElementMeta(func() ast.ElementKind { return RULE_ELEM })
}

type StringOrIdentifierStmt struct {
	Value     string
	TokenType lexer.TokenKind
}

func (s StringOrIdentifierStmt) Stmt() {}
func (s StringOrIdentifierStmt) Elem() ast.ElementMeta {
	return ast.InitElementMeta(func() ast.ElementKind {
		switch s.TokenType {
		case IDENTIFIER:
			return IDENTIFIER_ELEM
		default:
			return STRING_ELEM

		}
	})
}

type BodyStmt struct {
	Elements []ast.Element
}

func (b BodyStmt) Expr() {}
func (b BodyStmt) Elem() ast.ElementMeta {
	return ast.InitElementMeta(func() ast.ElementKind { return BODY_ELEM })
}

type SetExpr struct {
	BodyStmt   BodyStmt
	IsOptional bool
	IsGrouped  bool
	IsRepeated bool
	Qualifier  lexer.Token
}

func (g SetExpr) Expr() {}
func (g SetExpr) Elem() ast.ElementMeta {
	return ast.InitElementMeta(func() ast.ElementKind { return SET_ELEM })
}

type UnknownStmt struct {
	Token lexer.Token
	Pos   int
}

func (g UnknownStmt) Expr() {}
func (g UnknownStmt) Elem() ast.ElementMeta {
	return ast.InitElementMeta(func() ast.ElementKind { return UNK_ELEM })
}

type AlternativeExpr struct {
	Alternate ast.Element
}

func (a AlternativeExpr) Expr() {}
func (a AlternativeExpr) Elem() ast.ElementMeta {
	return ast.InitElementMeta(func() ast.ElementKind { return ALT_ELEM })
}

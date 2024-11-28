package grammar

import (
	"github.com/karlmoad/go_util_lib/parsing/ast"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
)

type RuleExpr struct {
	Identifier lexer.Token
	Body       ast.Expression
}

func (g RuleExpr) Expr() {}
func (g RuleExpr) Obj()  {}

type StringOrIdentifierExpr struct {
	Value     string
	TokenType lexer.TokenKind
}

func (s StringOrIdentifierExpr) Expr() {}
func (s StringOrIdentifierExpr) Obj()  {}

type BodyExpr struct {
	Elements []ast.Expression
}

func (b BodyExpr) Expr() {}
func (b BodyExpr) Obj()  {}

type SetExpr struct {
	BodyExpr
	IsOptional bool
	IsGrouped  bool
	IsRepeated bool
	Qualifier  lexer.Token
}

func (g SetExpr) Expr() {}
func (g SetExpr) Obj()  {}

type UnknownExpr struct {
	Token lexer.Token
	Pos   int
}

func (g UnknownExpr) Expr() {}
func (g UnknownExpr) Obj()  {}

type AlternativeExpr struct {
	Alternate ast.Expression
}

func (a AlternativeExpr) Expr() {}
func (a AlternativeExpr) Obj()  {}

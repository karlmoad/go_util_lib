package grammar

import (
	"github.com/karlmoad/go_util_lib/parsing/ast"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
)

type RuleExpr struct {
	Identifier lexer.Token
	Body       ast.Element
}

func (g RuleExpr) Expr() {}
func (g RuleExpr) Elem() string {
	return "Rule Expression"
}

type StringOrIdentifierStmt struct {
	Value     string
	TokenType lexer.TokenKind
}

func (s StringOrIdentifierStmt) Stmt() {}
func (s StringOrIdentifierStmt) Elem() string {
	return "String Or Identifier Statement"
}

type BodyStmt struct {
	Elements []ast.Element
}

func (b BodyStmt) Expr() {}
func (b BodyStmt) Elem() string {
	return "Body Statement"
}

type SetExpr struct {
	BodyStmt   BodyStmt
	IsOptional bool
	IsGrouped  bool
	IsRepeated bool
	Qualifier  lexer.Token
}

func (g SetExpr) Expr() {}
func (g SetExpr) Elem() string {
	return "Set Expression"
}

type UnknownStmt struct {
	Token lexer.Token
	Pos   int
}

func (g UnknownStmt) Expr() {}
func (g UnknownStmt) Elem() string {
	return "Unknown Statement"
}

type AlternativeExpr struct {
	Alternate ast.Element
}

func (a AlternativeExpr) Expr() {}
func (a AlternativeExpr) Elem() string {
	return "Alternative Expression"
}

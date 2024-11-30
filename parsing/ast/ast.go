package ast

type Expression interface {
	Expr()
}

type Statement interface {
	Stmt()
}

type Element interface {
	Elem() string
}

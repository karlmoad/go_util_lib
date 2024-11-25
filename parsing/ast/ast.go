package ast

type Expression interface {
	Expr()
}

type Statement interface {
	Stmt()
}

type ObjType interface {
	Obj()
}

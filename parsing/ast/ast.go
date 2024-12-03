package ast

type ElementKind int

const (
	UNKNOWN ElementKind = -1
)

type ElementKindFunction func() ElementKind

type ElementMeta struct {
	kind ElementKindFunction
}

func (e ElementMeta) Kind() ElementKind {
	return e.kind()
}

func InitElementMeta(kind ElementKindFunction) ElementMeta {
	return ElementMeta{kind: kind}
}

type Expression interface {
	Expr()
}

type Statement interface {
	Stmt()
}

type Element interface {
	Elem() ElementMeta
}

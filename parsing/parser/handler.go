package parser

import "github.com/karlmoad/go_util_lib/parsing/ast"

type ParsingHandler func(p *Parser) (ast.ObjType, bool)
type ConditionHandler func(p *Parser) bool

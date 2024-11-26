package dialect

import (
	"github.com/karlmoad/go_util_lib/parsing/lexer"
	"github.com/karlmoad/go_util_lib/parsing/parser"
)

type Dialect interface {
	RegisterLexer(reg *lexer.Registry)
	RegisterParser(reg *parser.Registry)
}

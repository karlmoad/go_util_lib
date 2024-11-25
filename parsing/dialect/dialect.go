package dialect

import (
	"github.com/karlmoad/go_util_lib/parsing/lexer"
	"github.com/karlmoad/go_util_lib/parsing/parser"
)

type Dialect interface {
	RegisterTokenKinds(reg *lexer.Registry)
	RegisterTokenizationHandlers(reg *lexer.Registry)
	RegisterParsingHandlers(reg *parser.Registry)
}

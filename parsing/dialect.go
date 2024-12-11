package parsing

import (
	"github.com/karlmoad/go_util_lib/parsing/lexer"
	"github.com/karlmoad/go_util_lib/parsing/parser"
)

type Dialect interface {
	RegisterLexer(reg *lexer.Registry)
	RegisterParser(reg *parser.Registry)
}

func NewParserForDialect(dialect Dialect) (*parser.Parser, error) {
	parReg := parser.NewParsingRegistry()
	lexReg := lexer.NewLexerRegistry()

	dialect.RegisterLexer(lexReg)
	dialect.RegisterParser(parReg)

	lex := lexer.NewLexer(lexReg)
	par := parser.NewParser(parReg, lex)

	return par, nil
}

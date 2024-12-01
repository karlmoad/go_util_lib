package parsing

import (
	"github.com/karlmoad/go_util_lib/parsing/dialects/grammar"
	"github.com/karlmoad/go_util_lib/parsing/errors"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
	"github.com/karlmoad/go_util_lib/parsing/parser"
)

type Dialect interface {
	RegisterLexer(reg *lexer.Registry)
	RegisterParser(reg *parser.Registry)
}

type DialectKind int

const (
	NONE DialectKind = iota
	GRAMMAR
)

func NewParserForDialect(dialect DialectKind) (*parser.Parser, error) {
	switch dialect {
	case GRAMMAR:
		return initDialect(grammar.NewGrammarDialect())
	default:
		return nil, errors.NewInvalidValueError("invalid dialect selection")
	}

}

func initDialect(dialect Dialect) (*parser.Parser, error) {
	parReg := parser.NewParsingRegistry()
	lexReg := lexer.NewLexerRegistry()

	dialect.RegisterLexer(lexReg)
	dialect.RegisterParser(parReg)

	lex := lexer.NewLexer(lexReg)
	par := parser.NewParser(parReg, lex)

	return par, nil
}

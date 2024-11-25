package lexer

import (
	"fmt"
	"github.com/karlmoad/go_util_lib/parsing/dialect"
)

type Lexer struct {
	Tokens  []Token
	source  string
	pos     int
	length  int
	dialect dialect.Dialect
	reg     *Registry
}

func Tokenize(source string, dialect dialect.Dialect) *Lexer {
	lexer := initLexer(source, dialect)
	for !lexer.isEOF() {
		matched := lexer.reg.EvaluateTokenizationHandlers(lexer)
		if !matched {
			var chunk string
			end := lexer.pos + 25
			if end >= lexer.length {
				chunk = lexer.source[lexer.pos:]
			} else {
				chunk = lexer.source[lexer.pos:end]
			}

			panic(fmt.Sprintf("Tokenizer::Error -> unexpected token near [%d]: %s", lexer.pos, chunk))
		}
	}

	lexer.push(NewToken(EOF, "EOF"))
	return lexer
}

func initLexer(source string, dialect dialect.Dialect) *Lexer {
	t := &Lexer{
		Tokens:  make([]Token, 0),
		source:  source,
		pos:     0,
		length:  len(source),
		dialect: dialect,
		reg:     newLexerRegistry(),
	}
	t.dialect.RegisterTokenKinds(t.reg)
	t.dialect.RegisterTokenizationHandlers(t.reg)
	return t
}

func (l *Lexer) advance(n int) {
	l.pos += n
}

func (l *Lexer) push(token Token) {
	l.Tokens = append(l.Tokens, token)
}

func (l *Lexer) at() byte {
	return l.source[l.pos]
}

func (l *Lexer) remainder() string {
	return l.source[l.pos:]
}

func (l *Lexer) isEOF() bool {
	return l.pos >= l.length
}

func (l *Lexer) TokenKindString(kind TokenKind) string {
	return l.reg.TokenKindToString(kind)
}

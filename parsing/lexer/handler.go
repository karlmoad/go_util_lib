package lexer

import (
	"github.com/karlmoad/go_util_lib/common/regex"
)

type TokenizationHandler func(lex *Lexer) (*Token, bool)

func RegexPatternHandler(pattern *regex.Pattern, kind TokenKind) TokenizationHandler {
	return func(lex *Lexer) (*Token, bool) {
		if match, valid := pattern.MatchSourceStart(lex.remainder()); valid {
			lex.advance(len(match))
			token := NewToken(kind, match)
			return &token, true
		} else {
			return nil, false
		}
	}
}

func RegexHandler(pattern *regex.Pattern, handler TokenizationHandler) TokenizationHandler {
	return func(lex *Lexer) (*Token, bool) {
		if _, valid := pattern.MatchSourceStart(lex.remainder()); valid {
			return handler(lex)
		} else {
			return nil, false
		}
	}
}

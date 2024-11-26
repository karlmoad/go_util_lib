package lexer

import "regexp"

type TokenizationHandler func(lex *Lexer) bool

func RegexPatternHandler(pattern *regexp.Regexp, kind TokenKind, value string) TokenizationHandler {
	return func(lex *Lexer) bool {
		match := pattern.FindStringIndex(lex.remainder())
		sourceVal := lex.remainder()[match[0]:match[1]]
		if match != nil && match[0] == 0 {
			lex.advance(len(sourceVal))
			lex.push(NewToken(kind, value))
			return true
		} else {
			return false
		}
	}
}

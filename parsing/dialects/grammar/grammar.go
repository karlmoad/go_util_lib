package grammar

import (
	"github.com/karlmoad/go_util_lib/common/regex"
	"github.com/karlmoad/go_util_lib/common/state"
	"github.com/karlmoad/go_util_lib/generics/queue"
	"github.com/karlmoad/go_util_lib/parsing/ast"
	"github.com/karlmoad/go_util_lib/parsing/lexer"
	"github.com/karlmoad/go_util_lib/parsing/parser"
)

const (
	NULL lexer.TokenKind = iota
	EOF
	UNKNOWN
	NUMBER
	STRING
	IDENTIFIER
	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_PAREN
	CLOSE_PAREN
	OPEN_BRACE
	CLOSE_BRACE
	COMMENT
	OPEN_COMMENT
	CLOSE_COMMENT
	OPERATOR
	OR
	AND
	DOT
	SEMICOLON
	COLON
	QUESTION
	COMMA
	STAR
	PLUS
	MINUS
	BLANK_LINE
	ELLIPSIS
	SLASH
	PERCENT
)

var (
	whitespacePattern            = regex.NewPattern(`^\s+`)
	dashCommentPattern           = regex.NewPattern(`^--\s*.*(\n|\r|\r\n)`)
	slashCommentPattern          = regex.NewPattern(`^//.*(\n|\r|\r\n)`) //, commentHandler},
	multilineCommentOpenPattern  = regex.NewPattern(`^/\*`)              //, multilineCommentHandler},
	multilineCommentClosePattern = regex.NewPattern(`^\*/`)
	doubleQuotedStringPattern    = regex.NewPattern(`^"[^"]*"`)                    //, stringHandler},
	singleQuotedStringPattern    = regex.NewPattern(`^'[^']*'`)                    //, stringHandler},
	numberPattern                = regex.NewPattern(`^[0-9]+(\.[0-9]+)?`)          //, numberHandler},
	symbolPattern                = regex.NewPattern(`^<?[a-zA-Z_][a-zA-Z0-9_]*>?`) //, symbolHandler},
	operatorPattern              = regex.NewPattern(`^(:*=)`)                      //, operatorHandler},
	ellipsisPattern              = regex.NewPattern(`^\.{3}`)                      //, ellipsisHandler},
	openBracketPattern           = regex.NewPattern(`^\[`)                         //, defaultHandler(OPEN_BRACKET, "[")},
	closeBracketPattern          = regex.NewPattern(`^\]`)                         //, defaultHandler(CLOSE_BRACKET, "]")},
	openParenPattern             = regex.NewPattern(`^\(`)                         //, defaultHandler(OPEN_PAREN, "(")},
	closeParenPattern            = regex.NewPattern(`^\)`)                         //, defaultHandler(CLOSE_PAREN, ")")},
	openBracePattern             = regex.NewPattern(`^\{`)                         //, defaultHandler(OPEN_BRACE, "{")},
	closeBracePattern            = regex.NewPattern(`^\}`)                         //, defaultHandler(CLOSE_BRACE, "}")},
	orPattern                    = regex.NewPattern(`^\|`)                         //, defaultHandler(OR, "|")},
	dotPattern                   = regex.NewPattern(`^\.`)                         //, defaultHandler(DOT, ".")},
	semicolonPattern             = regex.NewPattern(`^;`)                          //, defaultHandler(SEMICOLON, ";")},
	colonPattern                 = regex.NewPattern(`^:`)                          //, defaultHandler(COLON, ":")},
	questionPattern              = regex.NewPattern(`^\?`)                         //, defaultHandler(QUESTION, "?")},
	commaPattern                 = regex.NewPattern(`^,`)                          //, defaultHandler(COMMA, ",")},
	starPattern                  = regex.NewPattern(`^\*`)                         //, defaultHandler(STAR, "*")},
	plusPattern                  = regex.NewPattern(`^\+`)                         //, defaultHandler(PLUS, "+")},
	minusPattern                 = regex.NewPattern(`^-`)
	slashPattern                 = regex.NewPattern(`^/`)
	percentPattern               = regex.NewPattern(`^%`)
)

var tokenKindMap = map[lexer.TokenKind]string{
	NULL:          "INVALID NULL",
	EOF:           "EOF",
	UNKNOWN:       "UNKNOWN",
	NUMBER:        "NUMBER",
	STRING:        "STRING",
	IDENTIFIER:    "IDENTIFIER",
	OPEN_BRACKET:  "OPEN_BRACKET",
	CLOSE_BRACKET: "CLOSE_BRACKET",
	OPEN_PAREN:    "OPEN_PAREN",
	CLOSE_PAREN:   "CLOSE_PAREN",
	OPEN_BRACE:    "OPEN_BRACE",
	CLOSE_BRACE:   "CLOSE_BRACE",
	DOT:           "DOT",
	SEMICOLON:     "SEMICOLON",
	COLON:         "COLON",
	QUESTION:      "QUESTION",
	COMMA:         "COMMA",
	OR:            "OR",
	AND:           "AND",
	COMMENT:       "COMMENT",
	OPEN_COMMENT:  "OPEN_COMMENT",
	CLOSE_COMMENT: "CLOSE_COMMENT",
	OPERATOR:      "OPERATOR",
	STAR:          "STAR",
	PLUS:          "PLUS",
	MINUS:         "MINUS",
	BLANK_LINE:    "BLANKLINE",
	ELLIPSIS:      "ELLIPSIS",
	SLASH:         "SLASH",
	PERCENT:       "PERCENT",
}

type Grammar struct {
	state       state.Depth
	markerQueue queue.Queue[lexer.TokenKind]
}

func NewGrammarDialect() *Grammar {
	return &Grammar{markerQueue: queue.NewLIFOQueue[lexer.TokenKind]()}
}

func (g *Grammar) RegisterLexer(reg *lexer.Registry) {
	for k, n := range tokenKindMap {
		reg.RegisterTokenKind(k, n)
	}

	reg.RegisterTokenizationHandler(lexer.RegexHandler(whitespacePattern, g.SkipHandler(whitespacePattern)))
	reg.RegisterTokenizationHandler(lexer.RegexHandler(dashCommentPattern, g.SkipHandler(dashCommentPattern)))
	reg.RegisterTokenizationHandler(lexer.RegexHandler(slashCommentPattern, g.SkipHandler(slashCommentPattern)))
	reg.RegisterTokenizationHandler(lexer.RegexHandler(multilineCommentOpenPattern, g.MultilineCommentHandler))
	//reg.RegisterTokenizationHandler(lexer.RegexHandler(multilineCommentClosePattern, g.MultilineCommentHandler))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(doubleQuotedStringPattern, STRING))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(singleQuotedStringPattern, STRING))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(numberPattern, NUMBER))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(symbolPattern, IDENTIFIER))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(operatorPattern, OPERATOR))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(ellipsisPattern, ELLIPSIS))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(openBracketPattern, OPEN_BRACKET))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(closeBracketPattern, CLOSE_BRACKET))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(openParenPattern, OPEN_PAREN))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(closeParenPattern, CLOSE_PAREN))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(openBracePattern, OPEN_BRACE))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(closeBracePattern, CLOSE_BRACE))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(orPattern, OR))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(dotPattern, DOT))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(semicolonPattern, SEMICOLON))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(colonPattern, COLON))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(questionPattern, QUESTION))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(commaPattern, COMMA))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(starPattern, STAR))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(plusPattern, PLUS))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(minusPattern, MINUS))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(slashPattern, SLASH))
	reg.RegisterTokenizationHandler(lexer.RegexPatternHandler(percentPattern, PERCENT))

}

func (g *Grammar) RegisterParser(reg *parser.Registry) {
	reg.RegisterDefaultHandler(g.UnknownStatementHandler)
	reg.RegisterHandler(g.IsNewExpr, g.NewExpressionHandler)
	reg.RegisterHandler(parser.TokenKindCondition(IDENTIFIER), g.StringOrIdentifierHandler)
	reg.RegisterHandler(parser.TokenKindCondition(STRING), g.StringOrIdentifierHandler)
	reg.RegisterHandler(parser.TokenKindCondition(OR), g.AlternativeHandler)
	reg.RegisterHandler(parser.TokenKindCondition(OPEN_PAREN), g.GroupedExprHandler(OPEN_PAREN, CLOSE_PAREN, true, false, false))
	reg.RegisterHandler(parser.TokenKindCondition(OPEN_BRACKET), g.GroupedExprHandler(OPEN_BRACKET, CLOSE_BRACKET, false, true, false))
	reg.RegisterHandler(parser.TokenKindCondition(OPEN_BRACE), g.GroupedExprHandler(OPEN_BRACE, CLOSE_BRACE, false, false, true))
	reg.RegisterFixedCallback(g.NewExpressionCallback)
	reg.RegisterFixedCallback(g.EOFCallback)
}

//<editor-fold desc="lexicographical handlers and callbacks">

func (g *Grammar) SkipHandler(pattern *regex.Pattern) lexer.TokenizationHandler {
	return func(lex *lexer.Lexer) (*lexer.Token, bool) {
		if match, valid := pattern.MatchSourceStart(lex.Remainder()); valid {
			lex.Advance(len(match))
			return nil, true
		} else {
			return nil, false
		}
	}
}

func (g *Grammar) CommentStateCallback(lex *lexer.Lexer) bool {
	if match, valid := multilineCommentClosePattern.MatchSourceStart(lex.Remainder()); valid {
		lex.Advance(len(match))
		g.state.Decrease()
	} else {
		if match, valid := multilineCommentOpenPattern.MatchSourceStart(lex.Remainder()); valid {
			lex.Advance(len(match))
			g.state.Increase()
		} else {
			lex.Advance(1)
		}
	}
	return g.state.CurrentState() == false
}

func (g *Grammar) MultilineCommentHandler(lex *lexer.Lexer) (*lexer.Token, bool) {
	if match, valid := multilineCommentOpenPattern.MatchSourceStart(lex.Remainder()); valid {
		lex.Advance(len(match))
		lex.PushCallback(g.CommentStateCallback)
		g.state.Increase()
		return nil, true
	}
	return nil, false
}

//</editor-fold>

//<editor-fold desc="Parsing Handlers and callbacks">

func (g *Grammar) EOFCallback(p *parser.Parser) bool {
	return p.CurrentToken().Kind == EOF
}

func (g *Grammar) IsNewExpr(p *parser.Parser) bool {
	if p.CurrentToken().Kind == IDENTIFIER {
		if t, valid := p.PeekNext(); valid {
			if t.Kind == OPERATOR {
				return true
			}
		}
	}
	return false
}

func (g *Grammar) NewExpressionCallback(p *parser.Parser) bool {
	if g.IsNewExpr(p) && p.Depth() > 0 {
		p.DequeueCurrentCallback()
		return true
	}
	return false
}

func (g *Grammar) NewExpressionHandler(p *parser.Parser) (ast.Element, bool, error) {
	if g.IsNewExpr(p) {
		//p.PushCallback(g.NewExpressionCallback)

		if err := p.Expect(IDENTIFIER); err != nil {
			return nil, false, err
		}
		rule := RuleExpr{}
		rule.Identifier = p.Advance()
		if err := p.Expect(OPERATOR); err != nil {
			return nil, false, err
		}
		p.Advance()
		elem := make([]ast.Element, 0)
		for {
			if p.HasMoreTokens() {
				if e, err := p.ProcessNextToken(); e != nil && err == nil {
					elem = append(elem, e)
				} else {
					if err != nil {
						return nil, false, err
					} else {
						break
					}
				}
			} else {
				break
			}
		}
		rule.Body = BodyStmt{Elements: elem}
		return rule, true, nil
	}
	return nil, false, nil
}

func (g *Grammar) AlternativeHandler(p *parser.Parser) (ast.Element, bool, error) {
	if err := p.Expect(OR); err != nil {
		return nil, false, err
	}
	p.Advance()
	elem := make([]ast.Element, 0)
	for {
		if el, err := p.ProcessNextToken(); el != nil && err == nil {
			elem = append(elem, el)
		} else {
			if err != nil {
				return nil, false, err
			}
			break
		}
	}
	body := BodyStmt{Elements: elem}
	return AlternativeExpr{Alternate: body}, true, nil
}

func (g *Grammar) StringOrIdentifierHandler(p *parser.Parser) (ast.Element, bool, error) {
	if err := p.Expect(STRING, IDENTIFIER); err != nil {
		return nil, false, err
	}
	elmt := StringOrIdentifierStmt{Value: p.CurrentToken().Value, TokenType: p.CurrentToken().Kind}
	p.Advance()
	return elmt, true, nil
}

func (g *Grammar) UnknownStatementHandler(p *parser.Parser) (ast.Element, bool, error) {
	return UnknownStmt{Token: p.Advance(), Pos: p.Pos()}, true, nil
}

func (g *Grammar) GroupedExprHandler(open lexer.TokenKind, close lexer.TokenKind, group, optional, repeat bool) parser.ParsingHandler {
	return func(p *parser.Parser) (ast.Element, bool, error) {
		checkClose := func(p *parser.Parser) bool {
			if close == p.CurrentToken().Kind {
				return true
			}
			return false
		}

		if err := p.Expect(open); err != nil {
			return nil, false, err
		}

		p.PushCallback(checkClose)
		p.Advance()
		elem := make([]ast.Element, 0)
		for {
			ex, err := p.ProcessNextToken()
			if ex != nil {
				elem = append(elem, ex)
			} else {
				if err != nil {
					return nil, false, err
				} else {
					break
				}
			}
		}

		if err := p.Expect(close); err != nil {
			return nil, false, err
		}
		if cb, valid := p.CurrentCallback(); valid {
			if cb(p) && p.CurrentToken().Kind == close {
				p.DequeueCurrentCallback()
			}
		}
		p.Advance()
		body := BodyStmt{Elements: elem}
		exprSet := SetExpr{BodyStmt: body, IsGrouped: group, IsOptional: optional, IsRepeated: repeat}
		return exprSet, true, nil
	}
}

//</editor-fold>

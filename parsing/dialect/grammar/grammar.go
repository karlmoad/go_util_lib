package grammar

import (
	"fmt"
	"github.com/karlmoad/go_util_lib/common/regex"
	"github.com/karlmoad/go_util_lib/common/state"
	"github.com/karlmoad/go_util_lib/generics/queue"
	"github.com/karlmoad/go_util_lib/parsing/ast"
	"github.com/karlmoad/go_util_lib/parsing/dialect"
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
	state      state.Depth
	eventQueue queue.Queue[lexer.TokenKind]
}

func NewGrammarDialect() dialect.Dialect {
	return &Grammar{}
}

func (g *Grammar) RegisterLexer(reg *lexer.Registry) {
	for k, n := range tokenKindMap {
		reg.RegisterTokenKind(k, n)
	}

	reg.RegisterTokenizationHandler(lexer.RegexHandler(whitespacePattern, g.SkipHandler(whitespacePattern)))
	reg.RegisterTokenizationHandler(lexer.RegexHandler(dashCommentPattern, g.SkipHandler(dashCommentPattern)))
	reg.RegisterTokenizationHandler(lexer.RegexHandler(slashCommentPattern, g.SkipHandler(slashCommentPattern)))
	reg.RegisterTokenizationHandler(lexer.RegexHandler(multilineCommentOpenPattern, g.MultilineCommentHandler))
	reg.RegisterTokenizationHandler(lexer.RegexHandler(multilineCommentClosePattern, g.MultilineCommentHandler))
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

func (g *Grammar) RegisterParser(reg *parser.Registry) {}

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

func (g *Grammar) CommentExemptionCallback(lex *lexer.Lexer) bool {
	return !g.state.CurrentState()
}

func (g *Grammar) MultilineCommentHandler(lex *lexer.Lexer) (*lexer.Token, bool) {
	if match, valid := multilineCommentOpenPattern.MatchSourceStart(lex.Remainder()); valid {
		lex.Advance(len(match))
		if !g.state.CurrentState() {
			lex.PushCallback(g.CommentExemptionCallback)
		}
		g.state.Increase()
		return nil, true
	}

	if match, valid := multilineCommentClosePattern.MatchSourceStart(lex.Remainder()); valid {
		lex.Advance(len(match))
		g.state.Decrease()
		return nil, true
	}
	return nil, false
}

//</editor-fold>

//<editor-fold desc="Parsing Handlers and callbacks">

func (g *Grammar) isNewExpr(p *parser.Parser) bool {
	return p.CurrentToken().Kind == IDENTIFIER && p.PeekNext().Kind == OPERATOR
}

func (g *Grammar) NewExpressionCallback(p *parser.Parser) bool {
	return g.isNewExpr(p)
}

func (g *Grammar) GroupingStatementCallback(p *parser.Parser) bool {

}

func (g *Grammar) NewExpressionHandler(p *parser.Parser) (ast.ObjType, bool) {
	if g.isNewExpr(p) {
		p.PushCallback(g.NewExpressionCallback)

		if err := p.Expect(IDENTIFIER); err != nil {
			logger := p.GetLogger().With("handler", "ebnf grammar (NewExpressionHandler)")
			logger.Error(err.Error())
			return nil, false
		}
		rule := ast.GrammarRuleExpr{}
		rule.Identifier = p.advance()
		p.expect(lexer.OPERATOR)
		p.advance()
		expr := make([]ast.Expr, 0)
		for {
			if p.HasMoreTokens() {
				if e := parseExpr(p); e != nil {
					expr = append(expr, e)
				} else {
					break
				}
			} else {
				break
			}
		}
		rule.Body = ast.BodyExpr{Elements: expr}
		return rule, true
	}
	return nil, false
}

func (g *Grammar) AlternativeHandler(p *Parser) (ast.Expr, bool) {
	p.expect(lexer.OR)
	p.advance()
	expr := make([]ast.Expr, 0)
	for {
		if ex := parseExpr(p); ex != nil {
			expr = append(expr, ex)
		} else {
			break
		}
	}
	body := ast.BodyExpr{Elements: expr}
	return ast.AlternativeExpr{Alternate: body}, true
}

func (g *Grammar) stringOrIdentifierHandler(p *Parser) (ast.Expr, bool) {
	p.expect(lexer.STRING, lexer.IDENTIFIER)
	expr := ast.StringOrIdentifierExpr{Value: p.currentToken().Value, TokenType: p.currentToken().Kind}
	p.advance()
	return expr, true
}

func (g *Grammar) unknownTypeExprHandler(p *Parser) (ast.Expr, bool) {
	return ast.UnknownExpr{Token: p.advance(), Pos: p.pos}, true
}

func (g *Grammar) groupedExprHandler(open lexer.TokenKind, close lexer.TokenKind, group, optional, repeat bool) ExprHandler {
	return func(p *Parser) (ast.Expr, bool) {
		checkClose := func(p *Parser) bool {
			//See if the current token is a registered grouping indicator
			if alt, ok := p.reg.groupingTokens[p.currentToken().Kind]; ok {
				if open == alt {
					return true
				} else {
					panic(fmt.Sprintf("grouping termination mismatch expected %s got %s : [%d]", lexer.TokenKindToString(close), lexer.TokenKindToString(p.currentToken().Kind), p.pos))
				}
			}
			return false
		}

		p.expect(open)
		p.eventQueue.Enqueue(checkClose)
		p.advance()
		expr := make([]ast.Expr, 0)
		for {
			if checkClose(p) {
				break
			}
			if ex := parseExpr(p); ex != nil {
				expr = append(expr, ex)
			} else {
				break
			}
		}
		p.expect(close)
		p.advance()
		body := ast.BodyExpr{Elements: expr}
		exprSet := ast.SetExpr{BodyExpr: body, IsGrouped: group, IsOptional: optional, IsRepeated: repeat}
		return exprSet, true
	}
}

func (g *Grammar) ruleExprHandler(p *Parser) (ast.Expr, bool) {

}

//</editor-fold>

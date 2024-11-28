package lexer

import (
	"fmt"
	"github.com/karlmoad/go_util_lib/generics/queue"
	"github.com/karlmoad/go_util_lib/parsing"
	"github.com/karlmoad/go_util_lib/parsing/dialect"
	"strings"
)

type LexCallback func(lex *Lexer) bool

type Lexer struct {
	Tokens        []Token
	source        string
	pos           int
	length        int
	dialect       dialect.Dialect
	callbackQueue queue.Queue[LexCallback]
	reg           *Registry
}

func NewLexer(source string, dialect dialect.Dialect) *Lexer {
	t := &Lexer{
		Tokens:        make([]Token, 0),
		source:        source,
		pos:           0,
		length:        len(source),
		dialect:       dialect,
		reg:           newLexerRegistry(),
		callbackQueue: queue.NewLIFOQueue[LexCallback](),
	}
	t.dialect.RegisterLexer(t.reg)
	return t
}

func (l *Lexer) processCallbacks() bool {
	var ret bool

	//iterate exemption queue until empty or current == false
	if l.callbackQueue.Depth() > 0 {
		for {
			if funq, valid := l.callbackQueue.Current(); valid {
				if ret = funq(l); ret {
					l.callbackQueue.Dequeue()
				} else {
					break
				}
			} else {
				break
			}
		}
	} else {
		return true
	}
	return ret && l.callbackQueue.Depth() == 0
}

func (l *Lexer) Tokenize() error {
	for !l.isEOF() {
		token, matched := l.reg.EvaluateTokenizationHandlers(l)
		if !matched {
			return parsing.NewHandlerError(fmt.Sprintf("Tokenizer::Error -> unexpected token near [%d]", l.pos), l.pos)
		}

		//review callbacks for any exemptions
		//only push the token into the list if clear
		if l.processCallbacks() {
			l.push(*token)
		}
	}

	l.push(NewToken(EOF, "EOF"))
	return nil
}

func (l *Lexer) Advance(n int) {
	l.advance(n)
}

func (l *Lexer) advance(n int) {
	l.pos += n
}

func (l *Lexer) PushToken(token Token) {
	l.push(token)
}

func (l *Lexer) push(token Token) {
	l.Tokens = append(l.Tokens, token)
}

func (l *Lexer) at() byte {
	return l.source[l.pos]
}

func (l *Lexer) Remainder() string {
	return l.remainder()
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

func (l *Lexer) GetTokenSegment(start int, end int) map[int]Token {
	if start < 0 {
		start = 0
	}

	var segment []Token

	if end < start || end >= len(l.Tokens) {
		segment = l.Tokens[start:]
	} else {
		segment = l.Tokens[start:end]
	}

	rez := make(map[int]Token)
	itr := start
	for _, token := range segment {
		rez[itr] = token
		itr++
	}

	return rez
}

func (l *Lexer) GetContext(start int, end int) string {
	seg := l.GetTokenSegment(start, end)
	buffer := make([]string, 0)
	for position, token := range seg {
		buffer = append(buffer, fmt.Sprintf("[#(%d) %s]:%s", position, l.TokenKindString(token.Kind), token.Value))
	}
	return strings.Join(buffer, ", ")
}

func (l *Lexer) PushCallback(cb LexCallback) {
	l.callbackQueue.Enqueue(cb)
}

func (l *Lexer) ResetCallbacks() {
	l.callbackQueue.Clear()
}

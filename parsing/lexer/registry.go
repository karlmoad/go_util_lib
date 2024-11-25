package lexer

import "strings"

type Registry struct {
	tokenKinds           map[TokenKind]string
	tokenizationHandlers []TokenizationHandler
}

func newLexerRegistry() *Registry {
	reg := &Registry{tokenKinds: make(map[TokenKind]string), tokenizationHandlers: make([]TokenizationHandler, 0)}
	// add baseline values
	reg.RegisterTokenKind(EOF, "EOF")
	reg.RegisterTokenKind(UNKNOWN, "UNKNOWN")

	return reg
}

func (reg *Registry) RegisterTokenKind(kind TokenKind, name string) {
	reg.tokenKinds[kind] = strings.ToUpper(name)
}

func (reg *Registry) TokenKindToString(kind TokenKind) string {
	if val, ok := reg.tokenKinds[kind]; ok {
		return val
	} else {
		return reg.tokenKinds[UNKNOWN]
	}
}

func (reg *Registry) StringToTokenKind(s string) TokenKind {
	for k, v := range reg.tokenKinds {
		if strings.ToLower(strings.ToUpper(s)) == v {
			return k
		}
	}
	return UNKNOWN
}

func (reg *Registry) RegisterTokenizationHandler(handler TokenizationHandler) {
	reg.tokenizationHandlers = append(reg.tokenizationHandlers, handler)
}

func (reg *Registry) EvaluateTokenizationHandlers(lexer *Lexer) bool {
	for _, handler := range reg.tokenizationHandlers {
		if activated := handler(lexer); activated {
			return true
		}
	}
	return false
}

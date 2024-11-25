package lexer

type TokenKind int

const (
	EOF     TokenKind = -1001
	UNKNOWN TokenKind = -1002
)

type Token struct {
	Kind  TokenKind
	Value string
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}

func (t Token) IsKindOf(kinds ...TokenKind) bool {
	for _, kind := range kinds {
		if t.Kind == kind {
			return true
		}
	}
	return false
}

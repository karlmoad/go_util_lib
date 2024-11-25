package lexer

type InvalidTokenKindError struct{}

func (e *InvalidTokenKindError) Error() string {
	return "Invalid token kind"
}

package parsing

import "fmt"

type HandlerError struct {
	Message  string
	Position int
}

func (e *HandlerError) Error() string {
	return fmt.Sprintf("%s : [position (%d)]", e.Message, e.Position)
}

func NewHandlerError(message string, position int) *HandlerError {
	return &HandlerError{Message: message, Position: position}
}

type InvalidValueError struct {
	Message string
}

func (e *InvalidValueError) Error() string {
	return e.Message
}

func NewInvalidValueError(message string) *InvalidValueError {
	return &InvalidValueError{Message: message}
}

type EventHandler func(context interface{}) bool

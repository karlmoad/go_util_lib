package parser

import (
	"sync"
)

type Registry struct {
	conditions     []Condition
	handlers       []ParsingHandler
	defaultHandler ParsingHandler
	//escapeConditions []Condition  TODO Remove
	mut sync.Mutex
}

func NewParsingRegistry() *Registry {
	return &Registry{conditions: make([]Condition, 0),
		handlers: make([]ParsingHandler, 0)}
	//escapeConditions: make([]Condition, 0)} TODO Remove
}

func (r *Registry) RegisterHandler(condition Condition, handler ParsingHandler) {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.conditions = append(r.conditions, condition)
	r.handlers = append(r.handlers, handler)
}

// TODO Remove (Relocated and replaced)
//func (r *Registry) RegisterGlobalEscapeHandlers(condition Condition) {
//	r.mut.Lock()
//	defer r.mut.Unlock()
//	r.escapeConditions = append(r.escapeConditions, condition)
//}

func (r *Registry) RegisterDefaultHandler(handler ParsingHandler) {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.defaultHandler = handler
}

func (r *Registry) evaluateConditions(p *Parser) ParsingHandler {
	r.mut.Lock()
	defer r.mut.Unlock()
	for i, condition := range r.conditions {
		if assert := condition(p); assert {
			return r.handlers[i]
		}
	}
	return r.defaultHandler
}

//TODO REMOVE  (Relocated and replaced logic)
//func (r *Registry) evaluateEscapeConditions(p *Parser) bool {
//	r.mut.Lock()
//	defer r.mut.Unlock()
//	for _, condition := range r.escapeConditions {
//		if assert := condition(p); assert {
//			return true
//		}
//	}
//	return false
//}

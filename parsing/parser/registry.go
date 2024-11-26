package parser

import (
	"sync"
)

type Registry struct {
	conditions       []ConditionHandler
	handlers         []ParsingHandler
	defaultHandler   ParsingHandler
	escapeConditions []ConditionHandler
	mut              sync.Mutex
}

func newRegistry() *Registry {
	return &Registry{conditions: make([]ConditionHandler, 0),
		handlers:         make([]ParsingHandler, 0),
		escapeConditions: make([]ConditionHandler, 0)}
}

func (r *Registry) RegisterHandler(condition ConditionHandler, handler ParsingHandler) {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.conditions = append(r.conditions, condition)
	r.handlers = append(r.handlers, handler)
}

func (r *Registry) RegisterGlobalEscapeHandlers(condition ConditionHandler) {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.escapeConditions = append(r.escapeConditions, condition)
}

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

func (r *Registry) evaluateEscapeConditions(p *Parser) bool {
	r.mut.Lock()
	defer r.mut.Unlock()
	for _, condition := range r.escapeConditions {
		if assert := condition(p); assert {
			return true
		}
	}
	return false
}

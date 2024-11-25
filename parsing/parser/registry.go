package parser

import (
	"sync"
)

type Registry struct {
	conditions     []ConditionHandler
	handlers       []ParsingHandler
	defaultHandler ParsingHandler
	mut            sync.Mutex
}

func newRegistry() *Registry {
	return &Registry{conditions: make([]ConditionHandler, 0), handlers: make([]ParsingHandler, 0)}
}

func (r *Registry) RegisterHandler(condition ConditionHandler, handler ParsingHandler) {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.conditions = append(r.conditions, condition)
	r.handlers = append(r.handlers, handler)
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

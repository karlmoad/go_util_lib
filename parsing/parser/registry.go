package parser

import (
	"sync"
)

type Registry struct {
	conditions     []Condition
	handlers       []ParsingHandler
	defaultHandler ParsingHandler
	mut            sync.Mutex
	callbacks      []ParseCallback
}

func NewParsingRegistry() *Registry {
	return &Registry{conditions: make([]Condition, 0),
		handlers:  make([]ParsingHandler, 0),
		callbacks: make([]ParseCallback, 0)}
}

func (r *Registry) RegisterHandler(condition Condition, handler ParsingHandler) {
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

func (r *Registry) RegisterFixedCallback(callback ParseCallback) {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.callbacks = append(r.callbacks, callback)
}

func (r *Registry) evaluateCallbacks(p *Parser) bool {
	for _, cb := range r.callbacks {
		if cb(p) {
			return true
		}
	}
	return false
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

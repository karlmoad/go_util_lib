package state

import "sync"

type Depth struct {
	state bool
	depth int
	lock  sync.Mutex
}

func (s *Depth) CurrentState() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.state
}

func (s *Depth) CurrentDepth() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.depth
}

func (s *Depth) Increase() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.depth++
	s.evalCurrentState()
}

func (s *Depth) Decrease() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.depth > 0 {
		s.depth--
	}
	s.evalCurrentState()
}

func (s *Depth) evalCurrentState() {
	if s.depth > 0 {
		s.state = true
	}
	if s.depth == 0 {
		s.state = false
	}
}

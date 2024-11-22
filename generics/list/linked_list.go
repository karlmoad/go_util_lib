package list

import (
	"sync"
)

type node[T any] struct {
	item T
	next *node[T]
	prev *node[T]
}

// LinkedList Double Linked List implementation that allows for new values
// to be added to head of list, while allowing values to be removed
// from either head or tail of list
type LinkedList[T any] struct {
	head *node[T]
	tail *node[T]
	mutx sync.Mutex
	size int
}

func (h *LinkedList[T]) PushHead(item T) {
	h.mutx.Lock()
	defer h.mutx.Unlock()
	if h.head == nil {
		h.head = &node[T]{item, nil, nil}
		h.tail = h.head
	} else {
		node := &node[T]{item, nil, nil}
		node.next = h.head
		h.head.prev = node
		h.head = node
	}
	h.size++
}

func (h *LinkedList[T]) PopHead() (T, bool) {
	if h.head == nil {
		var zero T
		return zero, false
	} else {
		h.mutx.Lock()
		defer h.mutx.Unlock()
		node := h.head
		h.head = node.next
		h.eval()
		h.size--
		return node.item, true
	}
}

func (h *LinkedList[T]) eval() {
	if h.head != nil {
		h.head.prev = nil
	} else {
		h.tail = nil
	}

	if h.tail != nil {
		h.tail.next = nil
	} else {
		h.head = nil
	}
}

func (h *LinkedList[T]) PopTail() (T, bool) {
	if h.tail == nil {
		var zero T
		return zero, false
	} else {
		h.mutx.Lock()
		defer h.mutx.Unlock()
		node := h.tail
		h.tail = node.prev
		h.eval()
		h.size--
		return node.item, true
	}
}

func (h *LinkedList[T]) PeekHead() (T, bool) {
	if h.head == nil {
		return *new(T), false
	} else {
		return h.head.item, true
	}
}

func (h *LinkedList[T]) PeekTail() (T, bool) {
	if h.tail == nil {
		return *new(T), false
	} else {
		return h.tail.item, true
	}
}

func (h *LinkedList[T]) Size() int {
	return h.size
}

func (h *LinkedList[T]) Clear() {
	h.mutx.Lock()
	defer h.mutx.Unlock()
	h.head = nil
	h.tail = nil
	h.size = 0
}

func NewLInkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{head: nil, tail: nil, size: 0}
}

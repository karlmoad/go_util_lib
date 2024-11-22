package queue

import "github.com/karlmoad/go_util_lib/generics/list"

// LIFOQueue generic last in first out queue implemented with
// a linked list backend
type LIFOQueue[T any] struct {
	queue *list.LinkedList[T]
}

// NewLIFOQueue Create new instance of Last In First Out queue
func NewLIFOQueue[T any]() *LIFOQueue[T] {
	return &LIFOQueue[T]{queue: list.NewLInkedList[T]()}
}

// Enqueue Add new item to the queue
func (q *LIFOQueue[T]) Enqueue(item T) {
	q.queue.PushHead(item)
}

// Dequeue Remove item from queue
func (q *LIFOQueue[T]) Dequeue() (T, bool) {
	return q.queue.PopHead()
}

// Depth Get the size (number of items) contained within queue
func (q *LIFOQueue[T]) Depth() int {
	if q.queue == nil {
		return -1
	} else {
		return q.queue.Size()
	}
}

// Current Retrieve but do not remove the next item in the queue
func (q *LIFOQueue[T]) Current() (T, bool) {
	return q.queue.PeekHead()
}

// Clear Remove all contents of queue
func (q *LIFOQueue[T]) Clear() {
	q.queue.Clear()
}

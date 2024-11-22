package queue

import "github.com/karlmoad/go_util_lib/generics/list"

// FIFOQueue generic last in first out queue implemented with
// a linked list backend
type FIFOQueue[T any] struct {
	queue *list.LinkedList[T]
}

// NewFIFOQueue Create new instance of Last In First Out queue
func NewFIFOQueue[T any]() *FIFOQueue[T] {
	return &FIFOQueue[T]{queue: list.NewLInkedList[T]()}
}

// Enqueue Add new item to the queue
func (q *FIFOQueue[T]) Enqueue(item T) {
	q.queue.PushHead(item)
}

// Dequeue Remove item from queue
func (q *FIFOQueue[T]) Dequeue() (T, bool) {
	return q.queue.PopTail()
}

// Depth Get the size (number of items) contained within queue
func (q *FIFOQueue[T]) Depth() int {
	if q.queue == nil {
		return -1
	} else {
		return q.queue.Size()
	}
}

// Current Retrieve but do not remove the next item in the queue
func (q *FIFOQueue[T]) Current() (T, bool) {
	return q.queue.PeekTail()
}

// Clear Remove all contents of queue
func (q *FIFOQueue[T]) Clear() {
	q.queue.Clear()
}

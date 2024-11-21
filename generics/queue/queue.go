package queue

// Queue generic interface providing polymorphism
// to queue based types in this lib
type Queue[T any] interface {
	Enqueue(item T)
	Dequeue() (*T, bool)
	Current() (T, bool)
	Depth() int
	Clear()
}

package result

type Result[T any] struct {
	value   T
	nothing bool
}

func (r *Result[T]) Value() T {
	return r.value
}

func (r *Result[T]) Nothing() bool {
	return r.nothing
}

func (r *Result[T]) Set(value T) {
	r.value = value
	r.nothing = false
}

func NewResult[T any]() *Result[T] {
	return &Result[T]{nothing: true}
}

func NewResultWithValue[T any](value T) *Result[T] {
	return &Result[T]{value: value, nothing: false}
}

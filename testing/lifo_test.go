package testing

import (
	"testing"
)

func TestLIFOQueue_Enqueue(t *testing.T) {
	q := initAndFillQueue_int(true)
	if q.Depth() != 5 {
		t.Errorf("q.Depth() = %d, want %d", q.Depth(), 5)
	}
}

func TestLIFOQueue_Peek(t *testing.T) {
	q := initAndFillQueue_int(true)

	if val, ok := q.Current(); ok {
		if val != 5 {
			t.Errorf("q.Current() = %d, want %d", val, 1)
		}
	}
}

func TestLIFOQueue_Clear(t *testing.T) {
	q := initAndFillQueue_int(true)
	if q.Depth() != 5 {
		t.Errorf("q.Depth() = %d, want %d", q.Depth(), 0)
	}
	q.Clear()
	if q.Depth() != 0 {
		t.Errorf("q.Depth() = %d, want %d", q.Depth(), 0)
	}
}

func TestLIFOQueue_Dequeue(t *testing.T) {
	q := initAndFillQueue_int(true)

	for i := 5; i >= 1; i-- {
		if val, ok := q.Dequeue(); ok {
			if val != i {
				t.Errorf("q.Dequeue() = %d, want %d", val, i)
			}
		} else {
			t.Errorf("Dequeue() produced and unexpected value")
		}
	}

	if q.Depth() != 0 {
		t.Errorf("q.Depth() = %d, want %d", q.Depth(), 0)
	}

	// force one more dequeue affirm produces an invalid value
	if _, ok := q.Dequeue(); ok {
		t.Errorf("q.Dequeue() shoudl have produced an invalid value")
	}
}

func TestLIFOQueue_EnqueueDequeueComplex(t *testing.T) {
	q := initAndFillQueue_complex(true)
	if q.Depth() != 5 {
		t.Errorf("q.Depth() = %d, want %d", q.Depth(), 6)
	}

	var value int = 2

	for i := 5; i >= 1; i-- {
		if funq, ok := q.Dequeue(); ok {
			ret := funq(value)
			if ret != value*i {
				t.Errorf("func(%d) = %d, want %d", value, ret, i*value)
			}
		} else {
			t.Errorf("Dequeue() produced an unexpected invalid value")
		}
	}
}

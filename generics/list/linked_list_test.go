package list

import "testing"

func TestLinkedList_Push(t *testing.T) {
	ll := createAndFillIntList()
	if ll.head.item != 10 {
		t.Errorf("head item vale should be 10, is %d", ll.head.item)
	}
}

func TestLinkedList_Depth(t *testing.T) {
	// put 10 values in list check size
	ll := createAndFillIntList()
	if ll.Size() != 10 {
		t.Errorf("ll.Size() = %d, want 10", ll.Size())
	}
}

func TestLinkedList_PopHead(t *testing.T) {
	ll := createAndFillIntList()
	if v, ok := ll.PopHead(); ok {
		if *v != 10 {
			t.Errorf("PopHead() = %d, want 10", v)
		}
		if ll.head.item != 9 {
			t.Errorf("head item vale should be 9, is %d", ll.head.item)
		}
	} else {
		t.Errorf("ll.PopHead() produced invalid value")
	}
}

func TestLinkedList_PopTail(t *testing.T) {
	ll := createAndFillIntList()
	if v, ok := ll.PopTail(); ok {
		if *v != 1 {
			t.Errorf("PopTail() = %d, want 1", v)
		}
		if ll.tail.item != 2 {
			t.Errorf("head item vale should be 2, is %d", ll.tail.item)
		}
	} else {
		t.Errorf("ll.PopTail() produced invalid value")
	}
}

func TestNewLInkedList_Fill_RemAll_Head(t *testing.T) {
	ll := createAndFillIntList()
	var start int = 10
	for i := 0; i < 11; i++ {
		if v, ok := ll.PopHead(); ok {
			if *v != start {
				t.Errorf("PopHead() = %d, want %d", *v, start)
			}
			start--

			if i == 8 {
				if ll.head != ll.tail {
					t.Errorf("Last Item in list, head and tail should be equal again")
				}
			}

			if i == 9 {
				if ll.head != nil {
					t.Errorf("head is not nil, want nil")
				}
				if ll.tail != nil {
					t.Errorf("tail is not nil, want nil")
				}
			}
		} else {
			if i < 10 {
				t.Errorf("ll.PopHead() produced unexpected invalid value")
			}
		}
	}
}

func createAndFillIntList() *LinkedList[int] {
	ll := NewLInkedList[int]()
	for i := 0; i < 10; i++ {
		ll.PushHead(i + 1)
	}
	return ll
}

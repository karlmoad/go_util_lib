package testing

import "github.com/karlmoad/go_util_lib/generics/queue"

type testhandler func(i int) int

func initAndFillQueue_complex(lifo bool) queue.Queue[testhandler] {
	var que queue.Queue[testhandler]

	if lifo {
		que = queue.NewLIFOQueue[testhandler]()
	} else {
		que = queue.NewFIFOQueue[testhandler]()
	}

	for i := 1; i <= 5; i++ {
		que.Enqueue(func(x int) int { return x * i })
	}

	return que
}

func initAndFillQueue_int(lifo bool) queue.Queue[int] {
	var que queue.Queue[int]

	if lifo {
		que = queue.NewLIFOQueue[int]()
	} else {
		que = queue.NewFIFOQueue[int]()
	}

	for i := 1; i <= 5; i++ {
		que.Enqueue(i)
	}
	return que
}

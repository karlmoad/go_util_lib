package queue

type testhandler func(i int) int

func initAndFillQueue_complex(lifo bool) Queue[testhandler] {
	var queue Queue[testhandler]

	if lifo {
		queue = NewLIFOQueue[testhandler]()
	} else {
		queue = NewFIFOQueue[testhandler]()
	}

	for i := 1; i <= 5; i++ {
		queue.Enqueue(func(x int) int { return x * i })
	}

	return queue
}

func initAndFillQueue_int(lifo bool) Queue[int] {
	var queue Queue[int]

	if lifo {
		queue = NewLIFOQueue[int]()
	} else {
		queue = NewFIFOQueue[int]()
	}

	for i := 1; i <= 5; i++ {
		queue.Enqueue(i)
	}
	return queue
}

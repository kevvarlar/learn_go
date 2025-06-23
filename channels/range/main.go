package main

func concurrentFib(n int) []int {
	fibCh := make(chan int)
	fib := []int{}
	go fibonacci(n, fibCh)
	for num := range fibCh {
		fib = append(fib, num)
	}
	return fib
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

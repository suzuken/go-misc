package main

// sample codes for making deadlock
func main() {
	f()

	closed()

	gen(2, 3, 4)
}

// fatal error: all goroutines are asleep - deadlock!
func f() {
	ch := make(chan struct{})
	<-ch
}

// panic: send on closed channel
func closed() {
	ch := make(chan struct{})
	close(ch)
	ch <- struct{}{}
}

// forget close channel
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		// to be close channel here.
		// close(out)
	}()
	return out
}

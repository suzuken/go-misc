package main

// sample codes for making deadlock
func main() {
	f()

	ch := make(chan int)
	sq(ch)
}

func f() {
	ch := make(chan struct{})
	<-ch
}

// forget close channel
func sq(nums <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range nums {
			out <- n * n
		}
	}()
	return out
}

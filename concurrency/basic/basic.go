package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	nonBufferedSelect()
	buffered()
	for n := range sq(sq(gen(1, 2, 3))) {
		fmt.Println(n)
	}

	// fanout
	in := gen(2, 3)
	c1, c2 := sq(in), sq(in)
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}

	many()
}

func nonBufferedSelect() {
	ch := make(chan struct{})
	select {
	case <-ch:
		fmt.Println("finish")
	default:
		fmt.Println("default")
	}
}

func buffered() {
	ch := make(chan struct{}, 10)
	select {
	case <-ch:
		fmt.Println("finish")
	default:
		fmt.Println("default")
	}
}

// from https://blog.golang.org/pipelines
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func many() {
	task := make(chan string)
	quit := make(chan bool)
	workerquit := make(chan bool)

	go func() {
	loop:
		for {
			select {
			case <-quit:
				workerquit <- true
				break loop
			case job := <-task:
				fmt.Printf("make %s\n", job)
			}
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			task <- fmt.Sprintf("task %d", i+1)
			time.Sleep(1 * time.Second)
		}
		quit <- true
	}()

	<-workerquit
}

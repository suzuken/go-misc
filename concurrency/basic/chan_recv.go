package main

import (
	"fmt"
	"time"
)

func gen(ch chan<- string) {
	go func() {
		for {
			ch <- "🍜"
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			ch <- "🚀"
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		for {
			ch <- "🐰"
			time.Sleep(1 * time.Second)
		}
	}()
}

func main() {
	ch := make(chan string)
	gen(ch)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		}
	}
}

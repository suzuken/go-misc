package main

import (
	"fmt"
	"time"
)

func gen() chan string {
	ch := make(chan string)
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
	return ch
}

func main() {
	for v := range gen() {
		fmt.Println(v)
	}
}

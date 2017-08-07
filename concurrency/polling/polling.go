package main

import (
	"fmt"
	"time"
)

// from https://mattn.kaoriya.net/software/lang/go/20160706165757.htm
func main() {
	// buffered channel
	q := make(chan struct{}, 2)

	// sender
	go func() {
		time.Sleep(3 * time.Second)
		q <- struct{}{}
	}()

	// receiver
	for {
		if len(q) > 0 {
			break
		}
		time.Sleep(1 * time.Second)
		fmt.Println("something")
	}
}

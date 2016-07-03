package main

import (
	"fmt"
	"sync"
)

func main() {
	var mux sync.Mutex
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		mux.Lock()
		m["1"] = "a" // First conflicting access.
		mux.Unlock()
		c <- true
	}()
	mux.Lock()
	m["2"] = "b" // Second conflicting access.
	mux.Unlock()
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}
}

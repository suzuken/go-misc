package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, q <-chan string) {
	defer wg.Done()
	for {
		url, ok := <-q
		if !ok {
			// channel closed
			return
		}
		fmt.Printf("fetch %s ...\n", url)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	q := make(chan string, 5)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(&wg, q)
	}

	urls := []string{
		"https://example.com/1",
		"https://example.net/2",
		"https://example.org/3",
		"https://example.co.jp/4",
		"https://example.jp/5",
	}
	for _, url := range urls {
		q <- url
	}
	close(q)

	wg.Wait()
}

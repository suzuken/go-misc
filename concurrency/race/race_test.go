package main

import "testing"

func TestHelloWorld(t *testing.T) {
	m := map[int]bool{}

	go func() {
		m[1] = true
	}()

	go func() {
		_ = m[1]
	}()
}

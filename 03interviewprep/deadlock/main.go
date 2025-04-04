package main

import "sync"

var mu sync.Mutex

func main() {
	mu.Lock()
	mu.Lock()
	mu.Unlock()
	mu.Unlock()
}

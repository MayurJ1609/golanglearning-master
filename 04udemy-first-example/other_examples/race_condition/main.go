package main

import (
	"fmt"
	"sync"
)

var c = 0
var mu sync.Mutex
var wg sync.WaitGroup

func inc() {
	defer wg.Done()
	mu.Lock()
	c++
	mu.Unlock()
}

func dec() {
	defer wg.Done()
	mu.Lock()
	c--
	mu.Unlock()
}

func main() {

	wg.Add(2000)

	for i := 0; i < 1000; i++ {
		go inc()
		go dec()
	}

	wg.Wait()
	fmt.Println(c)
}

package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
	wg      sync.WaitGroup
)

func increment() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func main() {
	numGoroutines := 5

	// Start multiple goroutines to increment the counter
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go increment()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter)
}

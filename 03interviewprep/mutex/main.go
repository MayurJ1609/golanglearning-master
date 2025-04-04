package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	rwMutex sync.RWMutex
	wg      sync.WaitGroup
)

func readCounter(id int) {
	defer wg.Done()
	rwMutex.RLock()
	fmt.Printf("Goroutine %d reads counter value: %d\n", id, counter)
	rwMutex.RUnlock()
}

func incrementCounter(id int) {
	defer wg.Done()
	rwMutex.Lock()
	counter++
	fmt.Printf("Goroutine %d increments counter to: %d\n", id, counter)
	rwMutex.Unlock()
}

func main() {
	numReaders := 5
	numWriters := 3

	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go readCounter(i)
	}

	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go incrementCounter(i)
	}

	// time.Sleep(time.Second) // Sleep to let goroutines finish
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}

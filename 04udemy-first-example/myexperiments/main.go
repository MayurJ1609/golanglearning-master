package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, mutex *sync.Mutex) {
	defer wg.Done()
	mutex.Lock()
	msg = s
	mutex.Unlock()
}

func main() {
	var mutex sync.Mutex
	msg := "Hello, world"
	wg.Add(2)
	go updateMessage("Hello, Universe!", &mutex)
	go updateMessage("Hello, cosmos!", &mutex)
	wg.Wait()
	fmt.Println(msg)
}

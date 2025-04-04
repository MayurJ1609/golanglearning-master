package main

import (
	"fmt"
	"learning/golang-advance-learning/pkg/goroutines"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")

	fmt.Println("Goroutines")

	/*
		* Channels
			* Channels are used to communicate between goroutines
			* Consider channels are FIFO queue
			* Test
	*/

	goroutines.ChannelsExample()
	// Waitgroup example
	// goroutines.WaitGroupExample()

	// Channels Example
	// goroutines.ChannelsExample()
}

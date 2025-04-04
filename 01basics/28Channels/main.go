package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myChannel <- 5
	// fmt.Println(<-myChannel)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChannel

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		fmt.Println(<-myChannel)
		// fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myChannel <- 5
		myChannel <- 6
		close(myChannel)
		wg.Done()
	}(myChannel, wg)
	wg.Wait()
}

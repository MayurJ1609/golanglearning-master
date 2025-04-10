package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		fmt.Println("Got", i, "from channel")

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)

	// for i := 0; i <= 10; i++ {
	// 	ch <- i
	// }
	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("Sending", i, "to channel")
		ch <- i
		fmt.Println("Sent", i, "to channel")
	}

	fmt.Println("Done!")
	close(ch)
}

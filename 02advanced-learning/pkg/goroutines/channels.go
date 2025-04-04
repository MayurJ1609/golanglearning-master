package goroutines

import (
	"fmt"
)

func ChannelsExample() {

	c := make(chan string)
	go channelsExampleCount("Hello", c)
	// Type 1
	// for {
	// 	msg, open := <-c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	// Type 2
	for msg := range c {
		fmt.Println(msg)
	}

	// Take control on go scheduler explicitly (gosched())
	// runtime.Gosched()

}

func channelsExampleCount(s string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- s
	}

	// Close the channel - Mandatory
	// It will signal that no more values will be sent to the receiver
	close(c)
}

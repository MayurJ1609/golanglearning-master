package goroutines

import (
	"fmt"
	"sync"
)

func WaitGroupExample() {
	var wg sync.WaitGroup
	// n := 3
	wg.Add(1)

	// for i := 0; i < n; i++ {
	// 	go func() {
	// 		waitGroupExampleCount("World", i)
	// 		wg.Done()
	// 	}()
	// }
	go func() {
		waitGroupExampleCount("Hello")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Done!")
}

func waitGroupExampleCount(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, s)
	}
}

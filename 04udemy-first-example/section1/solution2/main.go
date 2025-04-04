package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	// Solution 2 : using sync.WaitGroup (not best, but easiest)
	var wg sync.WaitGroup

	words := []string{"alpha", "beta", "delta", "gamma", "pi", "zeta", "eta", "theta", "epsilon"}
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	wg.Wait()

	// Here wg will already be 0. So we need to add 1 to avoid it going -ve in next line
	wg.Add(1)
	printSomething(("This is the second thing to be printed"), &wg)

}

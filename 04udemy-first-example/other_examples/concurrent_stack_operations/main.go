package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	data []int
	mu   sync.Mutex
}

func (s *Stack) Push(x int) {
	s.mu.Lock()
	s.data = append(s.data, x)
	s.mu.Unlock()
}

func (s *Stack) Pop() (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.data) == 0 {
		return 0, false
	}
	x := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return x, true
}

func main() {
	stack := &Stack{}

	var wg sync.WaitGroup
	wg.Add(2)

	pushDone := make(chan struct{})
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			stack.Push(i)
			fmt.Println("Pushed", i)
		}
		close(pushDone)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			<-pushDone
			// fmt.Println("In POP", i)
			val, ok := stack.Pop()
			if ok {
				fmt.Println("Popped", val)
			}
		}
	}()

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

// The Dining Philosophers problem (5 Philosophers, 5 Forks)
// Solution: https://en.wikipedia.org/wiki/Dining_philosophers_problem

// Philosopher is a struct which stores a name and two forks
type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

// List of philosophers
var philosophers = []Philosopher{
	{
		name:      "Plato",
		leftFork:  4,
		rightFork: 0,
	},
	{
		name:      "Socrates",
		leftFork:  0,
		rightFork: 1,
	},
	{
		name:      "Aristotle",
		leftFork:  1,
		rightFork: 2,
	},
	{
		name:      "Pascal",
		leftFork:  2,
		rightFork: 3,
	},
	{
		name:      "Locke",
		leftFork:  3,
		rightFork: 4,
	},
}

// Define some variables
var hunger = 3 // howmany times the person eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

// *** added this
var orderMutex sync.Mutex  // a mutex for the slice orderFinished; part of challenge!
var orderFinished []string // the order in which philosophers finish dining and leave; part of challenge!

func main() {
	// print out a welcome message
	fmt.Println("Dining philosophers problem")
	fmt.Println("----------------------------")
	fmt.Println("The table is empty")

	// start the meal
	dine()

	// print out a message
	fmt.Println("The table is empty")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// Fire of a goroutine for current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s seated at the table\n", philosopher.name)
	seated.Done()

	seated.Wait()

	// eat three times
	for i := hunger; i > 0; i-- {
		// get lock on both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s puts down the forks.\n", philosopher.name)
	}
	fmt.Println(philosopher.name, "is satisfied.")
	fmt.Println(philosopher.name, "left the table")

	// *** added this
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}

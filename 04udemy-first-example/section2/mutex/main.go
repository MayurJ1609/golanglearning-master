package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

// Program to return how much money a person makes in 52 weeks
func main() {
	// variable to store for bank balance
	var bankBalance int
	var balanceMutex sync.Mutex

	// print out starting values
	fmt.Printf("Initial Account Balance: $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Partime Job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balanceMutex.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balanceMutex.Unlock()
				fmt.Printf("On week %d. you earned $%d.00 from %s.\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	//print out final balance
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()
}

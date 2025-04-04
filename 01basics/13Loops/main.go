package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang ")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	// For loop - Type 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// For loop - Type 2
	for i := range days {
		fmt.Println(days[i])
	}

	// For loop - Type 3
	// incase index is not used, use _ instead of index
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n ", index, day)
	}

	rougueValue := 1

	// For loop - Type 4
	for rougueValue < 10 {

		if rougueValue == 9 {
			goto mayur
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

mayur:
	fmt.Println("Jumping to goto mayur")
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proResult)

	retMulInt, retMulStr := returnMultiple()
	fmt.Println("Return multiple Int: ", retMulInt)
	fmt.Println("Return multiple Str: ", retMulStr)
}

// Type1
func greeter() {
	fmt.Println("Namaste from golang")
}

// Type2: Returns 1 output. Ex int value will be returned
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// Type 3
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func returnMultiple() (int, string) {
	return 23287, "Return multiple"
}

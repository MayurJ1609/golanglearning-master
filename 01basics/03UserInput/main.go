package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the rating for our Pizza:")

	// Comma Ok Syntax || Error Ok Syntax

	input, _ := reader.ReadString(('\n'))
	fmt.Println("Thanks for Rating: ", input)
	fmt.Printf("Type of this Rating is: %T", input)
}

package pointers

import "fmt"

func PointerEx() {
	var age int
	age = 42

	fmt.Println(age)

	// * is used for 2 things - 1. Declare a pointer variable, 2.
	var pointer *int
	pointer = &age

	fmt.Println("address: ", pointer)
	fmt.Println("value: ", *pointer)

	age = 32
	fmt.Println("value: ", *pointer)

	*pointer = 22
	fmt.Println("value age: ", age)
}

package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const Kevin = 273.15

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func ModifySlice(s []int) {
	s[0] = 100
}

func main() {
	var c Celsius = 37.0
	f := CToF(c)
	fmt.Println(f) // Output: 98.6

	// We cannot take address fo constant as it does not reside like a pointer
	// const testPointer int = 10
	// i := &testPointer - This line will throw syntax error

	var i int = 10
	fmt.Println(&i)

	j := 10
	fmt.Println(j)

	s := []int{1, 2, 3, 4, 5}
	fmt.Println("capacity: ", cap(s))
	fmt.Println("length: ", len(s))

	s = append(s, 6)
	fmt.Println("capacity: ", cap(s))
	fmt.Println("length: ", len(s))

	fmt.Println("Before func call: ", s)
	ModifySlice(s)
	fmt.Println("After func call: ", s)

}

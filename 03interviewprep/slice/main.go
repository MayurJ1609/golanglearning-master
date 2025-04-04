package main

import "fmt"

func main() {
	mySlice := make([]int, 5)

	mySlice = append(mySlice, 1)
	fmt.Println(cap(mySlice))
}

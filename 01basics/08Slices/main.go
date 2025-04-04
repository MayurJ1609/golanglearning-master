package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	// Create Slice - Type 1
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	// Add Value to slice using append() method
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) // Tomato, Peach, Mango, Banana
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // Peach, Mango
	fmt.Println(fruitList)

	// Create Slice - Type 2
	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	// Add Value to slice using append() method
	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	// Sort slice
	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// How to remove value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	// Remove Value in slice using append() method
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}

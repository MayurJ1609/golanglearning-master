package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")
	// var numberOne int = 2
	// var numberTwo float64 = 4

	// fmt.Println("The sum is: ", numberOne+int(numberTwo))

	// Random number - math/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// Random number - crypto / rand

	myRandomNom, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNom)

}

package main

import (
	"fmt"
	"learning/golang-advance-learning/pkg/closures"
)

func main() {

	useGiftCard1 := closures.ActivateGiftCard()
	useGiftCard2 := closures.ActivateGiftCard()

	fmt.Println(useGiftCard1(10))
	fmt.Println(useGiftCard2(5))

}

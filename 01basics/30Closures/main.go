package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func intSeqWithParam(a int) func() int {
	i := 0
	return func() int {
		i++
		i = i + a
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println("With parameters")

	newIntWithParam := intSeqWithParam(1)
	fmt.Println(newIntWithParam())
	fmt.Println(newIntWithParam())
}

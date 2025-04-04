package main

import (
	"flag"
	"log"
	"strconv"
	"time"
)

func calculateSum(outputCh chan int, a, b int) {
	result := 0
	result = a + b
	outputCh <- result
	return
}

func main() {
	var (
		frequency int
	)

	outputCh := make(chan int, 10)
	flag.Parse()
	frequency, _ = strconv.Atoi(flag.Args()[0])

	returnedCount := 0

	log.Println("Frequency selected", frequency)
	startTime := time.Now()

	go func() {
		for currIdx := 0; currIdx < frequency; currIdx++ {
			go calculateSum(outputCh, 4, currIdx)
		}
	}()

	for {
		<-outputCh
		returnedCount += 1
		if returnedCount == frequency {
			break
		}
	}
	timeTaken := time.Since(startTime)
	log.Println("Total Queries processed:", returnedCount, "in", timeTaken.Milliseconds(), "ms")

}

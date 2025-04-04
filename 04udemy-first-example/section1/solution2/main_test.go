package main

import (
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"testing"
)

// ***** RUN this command for testing: go test section1/solution2/main_test.go section1/solution2/main.go

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("epsilon", &wg)

	wg.Wait()
	w.Close()

	result, _ := ioutil.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon") {
		t.Errorf("Expected %s but got %s", "epsilon", output)
	}
}

package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// go test section2/mutex/main_test.go section2/mutex/main.go
func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$34320.00") {
		t.Errorf("Expected %s but got %s", "$34320.00", output)
	}
}

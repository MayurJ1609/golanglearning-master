package main_1

import "testing"

func Test_updateMessage(t *testing.T) {
	msg := "Hello, world!"

	wg.Add(2)
	go updateMessage("Test1")
	go updateMessage("Goodbye, cruel world")
	wg.Wait()

	if msg != "Goodbye, cruel world" {
		t.Errorf("Expected %s but got %s", "Goodbye, cruel world", msg)
	}
}

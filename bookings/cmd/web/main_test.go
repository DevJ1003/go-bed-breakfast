package main

import "testing"

// testing function for our main function

func TestRun(t *testing.T) {
	_, err := run()
	if err != nil {
		t.Error("Failed run()")
	}
}

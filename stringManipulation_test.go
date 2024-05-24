package main

import "testing"

func TestAtoi(t *testing.T) {
	testStr := "123"
	expected := 123
	got := atoi(testStr)

	if got != expected {
		t.Errorf("Got: %d", got)
		t.Errorf("Expected: %d", expected)
		t.Errorf("TestAtoi Failed!")
	}
}

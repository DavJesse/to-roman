package main

import "testing"

// Test atoi function
func TestAtoi(t *testing.T) {
	testStr := "123"
	expected := 123
	got, err := atoi(testStr)

	//Incase of a fail, print the 'got' and 'expected' variables
	if err != "" && got != expected {
		t.Errorf("Got: %d", got)
		t.Errorf("Expected: %d", expected)
		t.Errorf("TestAtoi Failed!")
	}
}

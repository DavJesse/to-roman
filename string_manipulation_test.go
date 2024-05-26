package main

import "testing"

func TestAtoi(t *testing.T) {
	testStr := "123"
	expected := 123
	got, err := atoi(testStr)

	if err != "" && got != expected {
		t.Errorf("Got: %d", got)
		t.Errorf("Expected: %d", expected)
		t.Errorf("TestAtoi Failed!")
	}
}

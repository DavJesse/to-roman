package main

import "testing"

func TestRomanNumeros(t *testing.T) {
	testInt := 999
	got, _ := romanNumeros(testInt)
	expected := "CMXCIX"

	if got != expected {
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestRomanNumeros Failed!")
	}
}

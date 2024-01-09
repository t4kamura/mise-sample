package main

import "testing"

func TestCalculate(t *testing.T) {
	if calculate(2, 4) != 6 {
		t.Error("Expected 2 + 4 to equal 6")
	}
}

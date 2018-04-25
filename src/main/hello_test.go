package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestSum2(t *testing.T) {
	total := Sum(38, 4)
	if total != 42 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

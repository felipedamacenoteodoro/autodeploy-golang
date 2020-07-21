package main

import "testing"

func TestSum(t *testing.T) {
	got := sum(10, 10)
	if got != 20 {
		t.Errorf("Sum return the value = %d; want 20", got)
	}
}

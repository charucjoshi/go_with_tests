package main

import "testing"

func TestAdder(t *testing.T){
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d, got %d", expected, sum)
	}
}

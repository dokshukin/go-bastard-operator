package main

import "testing"

func TestRandom(t *testing.T) {
	var v int
	v = random()
	if v < 0 {
		t.Error("Expected positive integer, got ", v)
	}

	if v >= len(excuses) {
		t.Error("array range exceeded: ", v)
	}
}

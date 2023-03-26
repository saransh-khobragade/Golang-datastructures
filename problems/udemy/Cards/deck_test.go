package main

import "testing"

func TestGetDeck(t *testing.T) {
	d := getDeck()

	if len(d) != 9 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}
}

package search

import (
	"testing"
)


func TestBinarySearchFindsExistingElement(t *testing.T) {
	ints := []int16{1, 10, 100, 1000}
	target := 100

	findResult := BinarySearch(ints, int16(target))
	if findResult == nil {
		t.Error("Binary search didn't find an element")
	}

	if *findResult != 2 {
		t.Error("Binary search didn't find the correct element")
	}
}
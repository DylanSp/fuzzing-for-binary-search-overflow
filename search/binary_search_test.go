package search

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestBinarySearchFindsExistingElement(t *testing.T) {
	ints := []uint8{1, 10, 100}
	target := 10

	findResult := BinarySearch(ints, uint8(target))
	if findResult == nil {
		t.Error("Binary search didn't find an element")
	}

	if *findResult != 1 {
		t.Error("Binary search didn't find the correct element")
	}
}

func TestBinarySearchDoesNotFindMissingElement(t *testing.T) {
	ints := []uint8{1, 10, 100}
	target := 5

	findResult := BinarySearch(ints, uint8(target))
	if findResult != nil {
		t.Error("Binary search found a nonexistent element")
	}
}

func FuzzBinarySearch(f *testing.F) {
	f.Fuzz(func(t *testing.T, sliceLength uint8, target uint8) {
		ints := createSortedFilledSlice(sliceLength)

		findResult := BinarySearch(ints, target)

		sliceActuallyContainsTarget := slices.Contains(ints, target)
		if sliceActuallyContainsTarget && findResult == nil {
			t.Error("Binary search didn't find existing element")
		} else if !sliceActuallyContainsTarget && findResult != nil {
			t.Error("Binary search erroneously found nonexisting element")
		}

		if findResult != nil && (ints[*findResult] != target) {
			t.Error("Binary search found incorrect element")
		}
	})
}

func createSortedFilledSlice(sliceLength uint8) []uint8 {
	arr := make([]uint8, sliceLength)

	for i := range arr {
		arr[i] = uint8(i)
	}

	return arr
}

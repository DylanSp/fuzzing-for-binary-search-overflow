package search

import (
	"sort"
)

// returns the index of searchTarget in sortedInts, or nil if it's not present
func BinarySearch(sortedInts []int16, searchTarget int16) *int {
	if !sort.SliceIsSorted(sortedInts, func(i, j int) bool {
		return sortedInts[i] < sortedInts[j]
	}) {
		panic("Tried to binary search a non-sorted array")
	}

	lowerBoundIndex := 0
	higherBoundIndex := len(sortedInts) - 1

	for lowerBoundIndex <= higherBoundIndex {
		midIndex := (lowerBoundIndex + higherBoundIndex) / 2
		midVal := sortedInts[midIndex]

		if midVal < searchTarget {
			lowerBoundIndex = midIndex + 1
		} else if midVal > searchTarget {
			higherBoundIndex = midIndex - 1
		} else {
			// value found
			return &midIndex
		}
	}

	return nil
}

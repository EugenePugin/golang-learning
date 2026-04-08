package main

import (
	"slices"
	"testing"
)

func Test_intersection(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "Wrong input - negative value",
			nums1:    []int{-3, 1, 2, 2, 2, 2, 4, 8},
			nums2:    []int{19, 4, 9, 8, 4},
			expected: []int{},
		},
		{
			name:     "Wrong input - too big value",
			nums1:    []int{3, 1, 2, 2, 2, 2, 4, 8},
			nums2:    []int{1009, 4, 9, 8, 4},
			expected: []int{},
		},
		{
			name:     "Intersection discovered",
			nums1:    []int{3, 1, 2, 2, 2, 2, 4, 8},
			nums2:    []int{19, 4, 9, 8, 4},
			expected: []int{4, 8},
		},
		{
			name:     "Intersection not discovered",
			nums1:    []int{3, 1, 2, 2, 2, 2, 4, 8},
			nums2:    []int{5, 7},
			expected: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersection(tt.nums1, tt.nums2)
			// TODO: update the condition below to compare got with tt.expected.
			if true != slices.Equal(got, tt.expected) {
				t.Errorf("intersection() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

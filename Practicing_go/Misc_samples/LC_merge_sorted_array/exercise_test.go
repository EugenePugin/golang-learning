package LCmsa

import (
	"slices"
	"testing"
)

func Test_mergeTestHelper(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "test",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "test",
			nums1:    []int{1, 0},
			m:        1,
			nums2:    []int{2},
			n:        1,
			expected: []int{1, 2},
		},
		{
			name:     "test",
			nums1:    []int{1, 0},
			m:        -1,
			nums2:    []int{2},
			n:        1,
			expected: nil,
		},
		{
			name:     "test",
			nums1:    []int{1, 0},
			m:        1,
			nums2:    []int{2},
			n:        201,
			expected: nil,
		},
		{
			name:     "test",
			nums1:    []int{1, 0},
			m:        0,
			nums2:    []int{2},
			n:        0,
			expected: nil,
		},
		{
			name:     "test",
			nums1:    []int{1, 0},
			m:        1,
			nums2:    []int{2},
			n:        200,
			expected: nil,
		},

		{
			name:     "test",
			nums1:    []int{1000000001, 0},
			m:        1,
			nums2:    []int{2},
			n:        1,
			expected: nil,
		},

		{
			name:     "test",
			nums1:    []int{1000000000, 0},
			m:        1,
			nums2:    []int{-1000000001},
			n:        1,
			expected: nil,
		},

		{
			name:     "test",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTestHelper(tt.nums1, tt.m, tt.nums2, tt.n)
			if slices.Compare(got, tt.expected) != 0 {
				t.Errorf("actual = %v, expected %v", got, tt.expected)
			}
		})
	}
}

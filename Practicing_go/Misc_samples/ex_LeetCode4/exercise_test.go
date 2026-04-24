package ex4

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "test1",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "test2",
			nums:     []int{2, 7, 11, 15},
			target:   18,
			expected: []int{1, 2},
		},

		{
			name:     "test3",
			nums:     []int{2, 7, 11, 15},
			target:   26,
			expected: []int{2, 3},
		},

		{
			name:     "test4",
			nums:     []int{2, 7, 11, 15},
			target:   17,
			expected: []int{0, 3},
		},
		{
			name:     "test5",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},

		{
			name:     "test6",
			nums:     []int{3, 2},
			target:   5,
			expected: []int{0, 1},
		},
		{
			name:     "test negative number",
			nums:     []int{-3, -2},
			target:   -5,
			expected: []int{0, 1},
		}, {
			name:     "test negative number2",
			nums:     []int{-8, -2},
			target:   -10,
			expected: []int{0, 1},
		},
		{
			name:     "no luck",
			nums:     []int{3, 2, 5},
			target:   -5,
			expected: []int{},
		},
		{
			name:     "negative numbers",
			nums:     []int{-13, 2, 5},
			target:   -8,
			expected: []int{0, 2},
		},
		{
			name:     "input validation <2",
			nums:     []int{3},
			target:   6,
			expected: []int{},
		},
		{
			name:     "input validation 110",
			nums:     []int{3, 110},
			target:   6,
			expected: []int{},
		}, {
			name:     "input validation 110",
			nums:     []int{3, -114},
			target:   6,
			expected: []int{},
		}, {
			name:     "input validation target110",
			nums:     []int{3, 15},
			target:   115,
			expected: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !slicesEqual(got, tt.expected) {
				t.Errorf("got = %v, expected %v", got, tt.expected)
				fmt.Println("got:", got, "expected: ", got, tt.expected)
			}
		})
	}
}

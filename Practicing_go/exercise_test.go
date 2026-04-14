package main

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{{
		"simple test {1}",
		[]int{1},
		1,
	},
		{
			"simple test {1, 2, 2}",
			[]int{1, 2, 2},
			2,
		},
		{
			"simple test {1, 2, 2,3}",
			[]int{1, 2, 2, 3},
			3,
		},

		{
			"simple test {1, 2, 2, 3, 4, 6}",
			[]int{1, 2, 3, 4, 6, 6},
			5,
		},
		{
			"simple test {2, 2, 2,3,2,2}",
			[]int{2, 2, 2, 3, 2, 2},
			2,
		},
		{
			"simple test {0,0,1,1,1,2,2,3,3,4}",
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		}, {
			"input validation {}",
			[]int{},
			-1,
		}, {
			"input validation {-101}",
			[]int{-101},
			-1,
		}, {
			"input validation {101}",
			[]int{101},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
				fmt.Println("removeDuplicates()", got, ", want", tt.want)
			}
		})
	}
}

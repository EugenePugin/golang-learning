package main

import (
	"slices"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "test",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "test",
			nums:   []int{2, 11, 7, 15},
			target: 9,
			want:   []int{0, 2},
		},
		{
			name:   "test",
			nums:   []int{11, 2, 7, 15},
			target: 9,
			want:   []int{1, 2},
		},
		{
			name:   "test",
			nums:   []int{11, 2, 7, 15},
			target: 26,
			want:   []int{0, 3},
		},
		{
			name:   "test",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "test",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !slices.Equal(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

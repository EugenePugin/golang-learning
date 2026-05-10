package LCpl

import (
	"slices"
	"testing"
)

func Test_addTwoValues(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "valid numbers",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "valid numbers",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 7},
			want: []int{0, 1, 8},
		},
		{
			name: "valid numbers",
			l1:   []int{2, 4},
			l2:   []int{5, 6, 4, 4},
			want: []int{8, 6, 6, 5},
		},
		{
			name: "valid numbers",
			l2:   []int{2, 4},
			l1:   []int{5, 6, 4, 4},
			want: []int{8, 6, 6, 5},
		},
		{
			name: "valid numbers",
			l2:   []int{2, 4, 4, 4},
			l1:   []int{5, 2, 4, 4},
			want: []int{8, 8, 6, 7},
		},
		{
			name: "input validation",
			l1:   []int{},
			l2:   []int{5, 6, 4},
			want: nil,
		},
		{
			name: "input validation",
			l1:   []int{5},
			l2:   []int{},
			want: nil,
		},
		{
			name: "input validation",
			l1:   []int{7, 3, 10, 5, 2, 8, 1, 6, 9, 4, 10, 3, 7, 1, 5, 8, 2, 6, 4, 9, 10, 2, 7, 5, 1, 8, 3, 6, 9, 4, 2, 10, 5, 7, 1, 8, 4, 6, 3, 9, 5, 1, 10, 7, 2, 8, 6, 4, 9, 3, 10, 5, 1, 7, 4, 8, 2, 6, 9, 3, 5, 10, 7, 1, 4, 8, 2, 6, 3, 9, 5, 10, 1, 7, 4, 8, 2, 6, 3, 9, 5, 10, 1, 7, 4, 8, 2, 6, 3, 9, 5, 10, 1, 7, 4, 8, 2, 6, 3, 9, 5},
			l2:   []int{5},
			want: nil,
		},
		{
			name: "input validation",
			l1:   []int{},
			l2:   []int{7, 3, 10, 5, 2, 8, 1, 6, 9, 4, 10, 3, 7, 1, 5, 8, 2, 6, 4, 9, 10, 2, 7, 5, 1, 8, 3, 6, 9, 4, 2, 10, 5, 7, 1, 8, 4, 6, 3, 9, 5, 1, 10, 7, 2, 8, 6, 4, 9, 3, 10, 5, 1, 7, 4, 8, 2, 6, 9, 3, 5, 10, 7, 1, 4, 8, 2, 6, 3, 9, 5, 10, 1, 7, 4, 8, 2, 6, 3, 9, 5, 10, 1, 7, 4, 8, 2, 6, 3, 9, 5, 10, 1, 7, 4, 8, 2, 6, 3, 9, 5},
			want: nil,
		},
		{
			name: "input validation",
			l1:   []int{1},
			l2:   []int{-1},
			want: nil,
		},
		{
			name: "input validation",
			l1:   []int{-1},
			l2:   []int{-1},
			want: nil,
		},
		{
			name: "input validation",
			l1:   []int{11},
			l2:   []int{0},
			want: nil,
		},
		{
			name: "input validation",
			l1:   []int{1},
			l2:   []int{14},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoValues(tt.l1, tt.l2)
			if !slices.Equal(got, tt.want) {
				t.Errorf("addTwoValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

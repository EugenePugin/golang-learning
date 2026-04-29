package LCplusone

import (
	"slices"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		digits []int
		want   []int
	}{
		{
			name:   "test",
			digits: []int{1, 2, 3},
			want:   []int{1, 2, 4},
		},
		{
			name:   "test",
			digits: []int{1, 2, 9},
			want:   []int{1, 3, 0},
		},
		{
			name:   "test",
			digits: []int{1, 9, 9},
			want:   []int{2, 0, 0},
		},
		{
			name:   "test",
			digits: []int{9},
			want:   []int{1, 0},
		},
		{
			name:   "test",
			digits: []int{9,9,9,9,9,9},
			want:   []int{1,0,0,0,0,0,0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.digits)
			if 0 != slices.Compare(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

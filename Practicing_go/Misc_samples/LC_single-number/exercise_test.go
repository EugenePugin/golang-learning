package main

import "testing"

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "test",
			nums: []int{2, 2, 3, 2},
			want: 3,
		},
		{
			name: "test",
			nums: []int{0, 1, 0, 1, 0, 1, 99},
			want: 99,
		},
		{
			name: "test",
			nums: []int{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.nums)
			if got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

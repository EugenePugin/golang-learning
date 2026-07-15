package LC_concurrency

import "testing"

func Test_goroutine_runner(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want string
	}{
		{
			name: "input validation",
			nums: []int{1},
			want: "error: nums len must be equal to 3",
		},
		{
			name: "input validation",
			nums: []int{1, 2, 3, 4, 5, 6},
			want: "error: nums len must be equal to 3",
		},
		{
			name: "input validation",
			nums: []int{1, 2, 15},
			want: "error: nums[i] must be equal to 1 or 2 or 3",
		},
		{
			name: "normal case",
			nums: []int{1, 2, 3},
			want: "firstsecondthird",
		},

		{
			name: "normal case",
			nums: []int{2, 3, 1},
			want: "firstsecondthird",
		},

		{
			name: "normal case",
			nums: []int{3, 1, 2},
			want: "firstsecondthird",
		},

		{
			name: "normal case",
			nums: []int{1, 3, 2},
			want: "firstsecondthird",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := goroutine_runner(tt.nums)
			if got != tt.want {
				t.Errorf("goroutine_runner() = %v, want %v", got, tt.want)
			}
		})
	}
}

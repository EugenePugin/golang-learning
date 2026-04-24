package lc_climbstairs

import "testing"

func Test_climbStairs(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{
			name: "test",
			n:    4,
			want: 5,
		},
		{
			name: "test",
			n:    2,
			want: 2,
		},

		{
			name: "test",
			n:    3,
			want: 3,
		},

		{
			name: "test",
			n:    46,
			want: -1,
		},
		{
			name: "test",
			n:    0,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbStairs(tt.n)
			if got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

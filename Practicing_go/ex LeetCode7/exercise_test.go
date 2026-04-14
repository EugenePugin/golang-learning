package ex7

import "testing"

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    int
		want int
	}{
		{
			name: "input validation: negative input",
			x:    -5,
			want: -1,
		},
		{
			name: "special case of 0",
			x:    0,
			want: 0,
		},
		{
			name: "special case of 1",
			x:    1,
			want: 1,
		},
		{
			name: "test 1: 2-1",
			x:    2,
			want: 1,
		}, {
			name: "test 1: 3-1",
			x:    3,
			want: 1,
		}, {
			name: "test 1: 4-2",
			x:    4,
			want: 2,
		}, {
			name: "test 1: 5-2",
			x:    5,
			want: 2,
		}, {
			name: "test 1: 6-2",
			x:    6,
			want: 2,
		}, {
			name: "test 1: 7-2",
			x:    7,
			want: 2,
		}, {
			name: "test 1: 8-2",
			x:    8,
			want: 2,
		},
		{
			name: "test 1: 9-3",
			x:    9,
			want: 3,
		},
		{
			name: "test 1: 10-3",
			x:    10,
			want: 3,
		},
		
		{
			name: "test 1: 255-15",
			x:    255,
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mySqrt(tt.x)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

package LCwc

import "testing"

func Test_maxAreaProxy(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		height []int
		want   int
	}{
		{
			name:   "test",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "test",
			height: []int{1, 3,2},
			want:   2,
		},
		{
			name:   "test",
			height: []int{1, 4, 2,4,1},
			want:   8,
		},
		{
			name:   "test",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "test",
			height: []int{1},
			want:   -1,
		},
		{
			name:   "test",
			height: []int{2, -3},
			want:   -1,
		},
		{
			name:   "test",
			height: []int{2, 10001},
			want:   -1,
		},
		
	// 	{
	// 		name:   "test",
	// 		height: []int{2, 10001},
	// 		want:   -1,
	// 	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxAreaProxy(tt.height)
			if got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

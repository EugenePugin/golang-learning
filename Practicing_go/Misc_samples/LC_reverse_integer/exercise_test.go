package LC_reverse_integer

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    int
		want int
	}{
		{
			name: "test",
			x:    123,
			want: 321,
		},
		{
			name: "test",
			x:    -123,
			want: -321,
		},
		{
			name: "test",
			x:    0,
			want: 0,
		},
		{
			name: "test",
			x:    0,
			want: 0,
		},
		{
			name: "test positive overflow",
			x:    1_247_483_646,
			want: 0,
		},
		{
			name: "test negative overflow",
			x:    -1_247_483_646,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.x)
			if got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

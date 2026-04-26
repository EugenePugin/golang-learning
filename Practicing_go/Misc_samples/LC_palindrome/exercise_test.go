package LCpalindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		x    int
		want bool
	}{
		{
			name:	"test",
			x:		1,
			want:	true,
		},
		{
			name:	"test",
			x:		11,
			want:	true,
		},
		{
			name:	"test",
			x:		121,
			want:	true,
		},
		{
			name:	"test",
			x:		12233221,
			want:	true,
		},
		{
			name:	"test",
			x:		122334221,
			want:	false,
		},
		{
			name:	"test",
			x:		-1221,
			want:	false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.x)
			if got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

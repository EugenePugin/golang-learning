package LC_brackets

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want bool
	}{
		{
			name: "test1",
			s:    "",
			want: false,
		},
		{
			name: "test2",
			s:    "()",
			want: true,
		},
		{
			name: "test3",
			s:    "())",
			want: false,
		},
		{
			name: "test3a",
			s:    "(()",
			want: false,
		},
		
		{
			name: "test3b",
			s:    "(){",
			want: false,
		},
		{
			name: "test4",
			s:    "(()))",
			want: false,
		},
		{
			name: "test5",
			s:    "({{)",
			want: false,
		},
		{
			name: "test6",
			s:    "([)]",
			want: false,
		},
		{
			name: "test7",
			s:    "([])",
			want: true,
		},
		{
			name: "test8",
			s:    "(]",
			want: false,
		},
		{
			name: "test9",
			s:    "()[]{}",
			want: true,
		},

		{
			name: "test10",
			s:    "([)]",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

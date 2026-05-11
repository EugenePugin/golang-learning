package fuzzy_testing

import (
	"testing"
	"unicode/utf8"
)

func Fuzz_reverseString(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		if !utf8.ValidString(orig) {
			t.Skip()
		}
		rev := reverseString(orig)
		doubleRev := reverseString(rev)
		if orig != doubleRev {
			t.Errorf("FAIL! \nOrig: %q (bytes %v)\nDoubleRev: %q (bytes %v)", orig, []byte(orig), doubleRev, []byte(doubleRev))
		}
	})
}
func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want string
	}{
		{
			name: "test",
			s:    "123GO",
			want: "OG321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseString(tt.s)
			if got != tt.want {
				t.Errorf("reverseSring() = %v, want %v", got, tt.want)
			}
		})
	}
}

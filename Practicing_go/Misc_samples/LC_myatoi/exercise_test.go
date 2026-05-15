package LC_myatoi

import "testing"

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{
			name: "test",
			s:    "42",
			want: 42,
		},
		{
			name: "ignore leading whitespace",
			s:    " 42",
			want: 42,
		},
		{
			name: "check valid use of the sign",
			s:    "-42",
			want: -42,
		},
		{
			name: "check valid use of the sign",
			s:    "+42",
			want: 42,
		},
		{
			name: "check invalid use of space symbol",
			s:    "442 45",
			want: 442,
		},
		{
			name: "check invalid use of sign symbol",
			s:    "4-42",
			want: 4,
		},
		{
			name: "check invalid use of sign symbol",
			s:    "+-42",
			want: 0,
		},
		{
			name: "check invalid use of sign symbol",
			s:    "-+42",
			want: 0,
		},
		{
			name: "check invalid use of sign symbol",
			s:    "44+2",
			want: 44,
		},

		{
			name: "non-digit number",
			s:    "-42a",
			want: -42,
		},
		{
			name: "just non-digit number",
			s:    "a",
			want: 0,
		},
		{
			name: "just to confirm",
			s:    "01",
			want: 1,
		},
		{
			name: "just to confirm",
			s:    "-01",
			want: -1,
		},
		{
			name: "check overflow",
			s:    "2147483648", //smth bigger then max int32 2147483647
			want: 2147483647,
		},
		{
			name: "normal case under overflow",
			s:    "2147483646", //smth bigger then max int32 2147483647
			want: 2147483646,
		},
		{
			name: "normal case under overflow",
			s:    "2147483647", //
			want: 2147483647,
		},
		{
			name: "check negative overflow",
			s:    "-2147483649", //
			want: -2147483648,
		},
{
			name: "check normal case near negative overflow",
			s:    "-2147483648", //
			want: -2147483648,
		},
		{
			name: "leading zeros and negative",
			s:    "  -042", //smth bigger then max int32 2147483647
			want: -42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myAtoi(tt.s)
			if got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

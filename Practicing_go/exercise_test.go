package main

import (
	"testing"
)

func Test_numDecodings(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{
			name: "test",
			s:    "12",
			want: 2,
		},
		{
			name: "input validation",
			s:    "",
			want: -1,
		},
			{
			name: "input validation",
			s:    "1A",
			want: -2,
		},
		{
			name: "input validation",
			s:    "01",
			want: -2,
		},
		{
			name: "input validation",
			s:    "10002",
			want: -2,
		},
				{
			name: "special case",
			s:    "5",
			want: 1,
		},
		{
			name: "normal case",
			s:    "1234",
			want: 3,
		},
		{
			name: "normal case",
			s:    "12034",
			want: 1,
		},
		{
			name: "normal case",
			s:    "12014",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numDecodings(tt.s)
			if got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}

package LCZigZag

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s       string
		numRows int
		want    string
	}{
		{
			name:    "normal",
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			name:    "normal",
			s:       "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			name:    "input validation",
			s:       "",
			numRows: 1,
			want:    "",
		},
		{
			name:    "input validation",
			s:       "AP",
			numRows: -2,
			want:    "",
		},
		{
			name:    "special case",
			s:       "A",
			numRows: 1,
			want:    "A",
		},
		{
			name:    "normal case",
			s:       "PAYP",
			numRows: 4,
			want:    "PAYP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convert(tt.s, tt.numRows)
			if got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

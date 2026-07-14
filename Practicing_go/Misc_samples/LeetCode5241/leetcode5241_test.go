package leetcode5241

import (
	"strings"
	"testing"
)

func Test_multiply(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		num1 string
		num2 string
		want string
	}{
		{
			num1: "2",
			num2: "3",
			want: "6",
		},
		{
			num1: "3",
			num2: "4",
			want: "12",
		},
		{
			num1: "4",
			num2: "25",
			want: "100",
		},
		{
			num1: "123",
			num2: "456",
			want: "56088",
		},

		{
			num1: "0",
			num2: "456",
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := multiply(tt.num1, tt.num2)
			// TODO: update the condition below to compare got with tt.want.
			if 0 != strings.Compare(got, tt.want) {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

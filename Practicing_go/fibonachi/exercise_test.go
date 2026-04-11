package fib

import (
	"fmt"
	"testing"
)

func Test_getFibonacchiListItemByOrderNumber(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		order_num uint
		expected  int
	}{
		{
			name:      "wrong input",
			order_num: 0,
			expected:  -1,
		},
		{
			name:      "special case1",
			order_num: 1,
			expected:  1,
		},
		{
			name:      "special case2",
			order_num: 2,
			expected:  1,
		},
		{
			name:      "regular case1",
			order_num: 3,
			expected:  2,
		}, {
			name:      "regular case2",
			order_num: 4,
			expected:  3,
		}, {
			name:      "regular case3",
			order_num: 5,
			expected:  5,
		}, {
			name:      "regular case4",
			order_num: 6,
			expected:  8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFibonacchiListItemByOrderNumber(tt.order_num)
			// fmt.Println("got:", got, "expected: ", got, tt.expected)
			// TODO: update the condition below to compare got with tt.expected.
			if got != tt.expected {
				t.Errorf("got = %v, expected %v", got, tt.expected)
				fmt.Println("got:", got, "expected: ", got, tt.expected)
			}
		})
	}
}

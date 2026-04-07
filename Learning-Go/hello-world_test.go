package main

import "testing"

func TestSome_function_to_test(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    int
		b    int
		want int
	}{
		{
			a:    7,
			b:    7,
			want: 14,
		},
		{
			a:    6,
			b:    6,
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Some_function_to_test(tt.a, tt.b)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("some_function_to_test() = %d, want %d", got, tt.want)
			}
		})
	}
}

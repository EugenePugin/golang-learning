package LC5

import (
	"fmt"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s        string
		expected int
	}{
		{
			name:     "simple checks I",
			s:        "I",
			expected: 1,
		},
		{
			name:     "simple checks II",
			s:        "II",
			expected: 2,
		},
		{
			name:     "simple checks III",
			s:        "III",
			expected: 3,
		},
		{
			name:     "simple checks V",
			s:        "V",
			expected: 5,
		},
		{
			name:     "simple checks X",
			s:        "X",
			expected: 10,
		},
		{
			name:     "simple checks L ",
			s:        "L",
			expected: 50,
		},
		{
			name:     "simple checks C",
			s:        "C",
			expected: 100,
		},

		{
			name:     "simple checks D",
			s:        "D",
			expected: 500,
		},
		{
			name:     "simple checks M",
			s:        "M",
			expected: 1000,
		},
		{
			name:     "special checks IV",
			s:        "IV",
			expected: 4,
		},
		{
			name:     "special checks VI",
			s:        "VI",
			expected: 6,
		},
		{
			name:     "special checks IX",
			s:        "IX",
			expected: 9,
		}, {
			name:     "special checks XIX",
			s:        "XIX",
			expected: 19,
		},
		{
			name:     "special checks XI",
			s:        "XI",
			expected: 11,
		},
		{
			name:     "special checks XXI",
			s:        "XXI",
			expected: 21,
		},
		{
			name:     "special checks XL",
			s:        "XL",
			expected: 40,
		}, {
			name:     "special checks LX",
			s:        "LX",
			expected: 60,
		},
		{
			name:     "special checks XC",
			s:        "XC",
			expected: 90,
		},
		{
			name:     "special checks XCIV",
			s:        "XCIV",
			expected: 94,
		},
		{
			name:     "special checks CX",
			s:        "CX",
			expected: 110,
		},
		{
			name:     "special checks CD",
			s:        "CD",
			expected: 400,
		},
		{
			name:     "special checks DC",
			s:        "DC",
			expected: 600,
		},
		{
			name:     "special checks CM",
			s:        "CM",
			expected: 900,
		},
		{
			name:     "special checks MC",
			s:        "MC",
			expected: 1100,
		},
		{
			name:     "special checks MCM",
			s:        "MCM",
			expected: 1900,
		},
		{
			name:     "special checks DCXXI",
			s:        "DCXXI",
			expected: 621,
		},
		{
			name:     "couple of more LVIII",
			s:        "LVIII",
			expected: 58,
		},
		{
			name:     "couple of more MCMXCIV",
			s:        "MCMXCIV",
			expected: 1994,
		},
		{
			name:     "wrong input - too big length",
			s:        "MCMXCIVMMMMMMMMMMMMMMMM",
			expected: -1,
		},
		{
			name:     "wrong input - incorrect symbol",
			s:        "MCMXCPIV",
			expected: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.s)
			if got != tt.expected {
				t.Errorf("got = %v, expected %v", got, tt.expected)
				fmt.Println("got:", got, "expected: ", got, tt.expected)
			}
		})
	}
}

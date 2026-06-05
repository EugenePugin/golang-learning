package LC_integer_to_Roman

import (
	"testing"
)

func Test_intToRoman(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		num  int
		want string
	}{
		// {
		// 	name: "test",
		// 	num:  3949,
		// 	want: "MMMDCCXLIX",
		// },
		{
			name: "test1",
			num:  2000,
			want: "MM",
		},
		{
			name: "test2",
			num:  500,
			want: "D",
		},

		{
			name: "test3",
			num:  2500,
			want: "MMD",
		},

		{
			name: "test4",
			num:  100,
			want: "C",
		},
		{
			name: "test5",
			num:  1800,
			want: "MDCCC",
		},

		{
			name: "test6",
			num:  50,
			want: "L",
		},
		{
			name: "test7",
			num:  2650,
			want: "MMDCL",
		},

		{
			name: "test8",
			num:  10,
			want: "X",
		},
		{
			name: "test9",
			num:  2680,
			want: "MMDCLXXX",
		},

		{
			name: "test10",
			num:  5,
			want: "V",
		},
		{
			name: "test11",
			num:  2685,
			want: "MMDCLXXXV",
		},

		{
			name: "test12",
			num:  900,
			want: "CM",
		},
		{
			name: "test13",
			num:  930,
			want: "CMXXX",
		},
		{
			name: "test14",
			num:  400,
			want: "CD",
		},
		{
			name: "test",
			num:  465,
			want: "CDLXV",
		},
		{
			name: "test",
			num:  90,
			want: "XC",
		},
		{
			name: "test",
			num:  95,
			want: "XCV",
		},
		{
			name: "test",
			num:  40,
			want: "XL",
		},
		{
			name: "test",
			num:  45,
			want: "XLV",
		},

		{
			name: "test",
			num:  9,
			want: "IX",
		},
		{
			name: "test",
			num:  4,
			want: "IV",
		},

		{
			name: "test",
			num:  8,
			want: "VIII",
		},
		{
			name: "test",
			num:  58,
			want: "LVIII",
		},
		
		{
			name: "test",
			num:  1994,
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToRoman(tt.num)
			if got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

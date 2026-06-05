// https://leetcode.com/problems/integer-to-roman/description/?envType=problem-list-v2&envId=dh5241mj
package LC_integer_to_Roman

import (
	"strings"
)

// Symbol	Value
// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000

// func convertToMDCLX(int base, string symbol, string *resultString) string {
// 	fmt.Println("hey, anonymous:", base, symbol)
// 	return &resultString
// }

func intToRoman(num int) string {
	var result strings.Builder

	number_of_Ms := num / 1000
	// fmt.Println("number_of_Ms:", number_of_Ms)
	if number_of_Ms > 0 {
		for range number_of_Ms {
			result.WriteString("M")
		}
	}
	num -= number_of_Ms * 1000

	// fmt.Println("interim check 1) num:", num)
	if num >= 900 {
		result.WriteString("CM")
		num -= 900
	}
	// fmt.Println("interim check 2) num:", num)

	number_of_Ds := num / 500
	// fmt.Println("number_of_Ds:", number_of_Ds)
	if number_of_Ds > 0 {
		for range number_of_Ds {
			result.WriteString("D")
		}
	}
	num -= number_of_Ds * 500
	// fmt.Println("interim check 3) num:", num)

	if num >= 400 {
		result.WriteString("CD")
		num -= 400
	}

	number_of_Cs := num / 100
	// fmt.Println("number_of_Cs:", number_of_Cs)
	if number_of_Cs > 0 {
		for range number_of_Cs {
			result.WriteString("C")
		}
	}
	num -= number_of_Cs * 100
	// fmt.Println("interim check 4) num:", num)

	if num >= 90 {
		result.WriteString("XC")
		num -= 90
	}
	// fmt.Println("interim check 5) num:", num)

	number_of_Ls := num / 50
	// fmt.Println("number_of_Ls:", number_of_Ls)
	if number_of_Ls > 0 {
		for range number_of_Ls {
			result.WriteString("L")
		}
	}
	num -= number_of_Ls * 50

	// fmt.Println("interim check 6) num:", num)

	if num >= 40 {
		result.WriteString("XL")
		num -= 40
	}
	// fmt.Println("interim check 7) num:", num)

	number_of_Xs := num / 10
	// fmt.Println("number_of_Xs:", number_of_Xs)
	if number_of_Xs > 0 {
		for range number_of_Xs {
			result.WriteString("X")
		}
	}
	num -= number_of_Xs * 10

	// fmt.Println("interim check 7) num:", num)
	if num >= 9 {
		result.WriteString("IX")
		num -= 9
	}
	// fmt.Println("interim check 8) num:", num)

	number_of_Vs := num / 5
	// fmt.Println("number_of_Vs:", number_of_Vs)
	if number_of_Vs > 0 {
		for range number_of_Vs {
			result.WriteString("V")
		}
	}
	num -= number_of_Vs * 5

	// fmt.Println("interim check 9) num:", num)
	if num >= 4 {
		result.WriteString("IV")
		num -= 4
	}
	// fmt.Println("interim check 10) num:", num)
	for range num {
		result.WriteString("I")
	}
	return result.String()
}

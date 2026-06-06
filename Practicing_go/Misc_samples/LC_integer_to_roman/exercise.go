// https://leetcode.com/problems/integer-to-roman/description/?envType=problem-list-v2&envId=dh5241mj
package main

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

var romanValues = [...]struct {
	value       int
	romanString string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	var result strings.Builder
	result.Grow(15)

	for _, rv := range romanValues {
		for num >= rv.value {
			result.WriteString(rv.romanString)
			num -= rv.value
		}
	}

	return result.String()
}

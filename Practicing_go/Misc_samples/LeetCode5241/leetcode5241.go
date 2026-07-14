// https://leetcode.com/problems/multiply-strings/?envType=problem-list-v2&envId=dh5241mj

package leetcode5241

import (
	"strconv"
)

func multiply1(num1 string, num2 string) string {
	multiplier1, _ := strconv.Atoi(num1)
	multiplier2, _ := strconv.Atoi(num2)
	res := multiplier1 * multiplier2

	// fmt.Println(multiplier1, multiplier2, res)
	return strconv.Itoa(res)
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1, len2 := len(num1), len(num2)
	// fmt.Println(len1, len2)
	res := make([]int, len1+len2)
	for i := len1 - 1; i >= 0; i-- {
		// fmt.Println(num1[i] - '0')
		for j := len2 - 1; j >= 0; j-- {
			// fmt.Println(" ", num2[j]-'0')
			mul := int(num2[j]-'0') * int(num1[i]-'0')
			sum := mul + res[i+j+1]
			// fmt.Println(" ", mul, sum)
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	// fmt.Println(res)

	start := 0
	for i := range res {
		if res[i] == 0 {
			start++
		} else {
			break
		}
	}
	// fmt.Println("position to start non-zero values", start)

	bytes := make([]byte, 0)
	for i := start; i < len(res); i++ {
		bytes = append(bytes, byte(res[i]+'0'))
	}
	// 	fmt.Println("converted to slices of bytes", bytes)

	// fmt.Println("converted to string", string(bytes))

	return string(bytes)
}

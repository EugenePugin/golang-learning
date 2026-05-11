// https://leetcode.com/problems/reverse-integer/description/

package LC_reverse_integer

import (
	"math"
)

func reverse(x int) int {
	var digit, result int
	for x != 0 {
		digit = x % 10
		x /= 10

		if (result > math.MaxInt32/10) ||
			(result == math.MaxInt32/10 && digit > 7) {
			// fmt.Println("overflow danger")
			return 0
		}

		if (result < math.MinInt32/10) ||
			((result == math.MinInt32/10) && (digit < -8)) {
			// fmt.Println("overflow danger")
			return 0
		}

		result = result*10 + digit
		// fmt.Println(digit, x, result)
	}

	return result
}

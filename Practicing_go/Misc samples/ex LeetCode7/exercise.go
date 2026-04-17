// https://leetcode.com/problems/sqrtx/description/
package ex7

func mySqrt(x int) int {
	// input validation
	// non-negative
	if x < 0 {
		return -1
	}

	// logic itself

	// special cases: 0 ,1
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	}
	// other check from 2 to evaluate with x2
	max_level := x >> 1
	// fmt.Println("max_level:", max_level)
	for i := 1; i <= max_level; i++ {
		if i*i == x {
			// fmt.Println("Bingo:", i)
			return i
		}
		if i*i > x {
			// fmt.Println("Bingo:", i)
			return i - 1
		}
	}
	// fmt.Println("And so :", max_level)
	return max_level
}

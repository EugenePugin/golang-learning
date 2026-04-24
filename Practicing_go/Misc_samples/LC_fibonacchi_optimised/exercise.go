// https://leetcode.com/problems/climbing-stairs/description/

package lc_climbstairs

func climbStairs_v1(n int) int {
	// input validation
	if n > 45 || n < 1 {
		return -1
	}
	// logic itself
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)

}

var mapInterimResults map[int]int

func climbStairs(n int) int {
	// input validation
	if n > 45 || n < 1 {
		return -1
	}

	if n <= 2 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

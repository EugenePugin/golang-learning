// https://leetcode.com/problems/climbing-stairs/description/

package main

import "fmt"

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

func main() {
	n := 8
	// mapInterimResults = make(map[int]int)

	fmt.Println(climbStairs(n))
}

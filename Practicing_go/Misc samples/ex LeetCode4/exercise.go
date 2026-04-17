// https://leetcode.com/problems/two-sum/

package ex4

import (
	"fmt"
	"math"
)

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func twoSum(nums []int, target int) []int {

	var result []int

	fmt.Println("nums:", nums, "target:", target)
	// inputs validation:
	// 2 <= nums.length <= 10^4
	// -109 <= nums[i] <= 10^9
	// -109 <= target <= 10^9
	if len(nums) < 2 || len(nums) > int(math.Pow(10, 4)) {
		return result
	}
	for i := range len(nums) {
		if math.Abs(float64(nums[i])) > math.Pow(10, 9) {
			return result
		}
	}
	if math.Abs(float64(target)) > math.Pow(10, 9) {
		return result
	}
	// logic itself
	var element0, element1 int
	// var shouldQuit bool
	for i := range len(nums) {
		//fmt.Println("i:", i)
		element0 = nums[i]
		//fmt.Println("element0:", element0)

		for j := i + 1; j < len(nums); j++ {
			// fmt.Println("j:", j)
			element1 = nums[j]
			// fmt.Println("element1:", element1)
			if element0+element1 == target {
				fmt.Println("Bingo!")
				result = make([]int, 2)
				result[0] = i
				result[1] = j
				return result
			}
			//	}

		}
	}

	return result
}

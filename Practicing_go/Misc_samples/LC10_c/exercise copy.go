//https://leetcode.com/problems/remove-element/description/

package LC10_c

import (
	"slices"
)

func removeElement(nums []int, val int) int {
	// input validation check
	if val < 0 || val > 100 {
		return -1
	}
	if len(nums) > 100 {
		return -2
	}
	for i := range len(nums) {
		if nums[i] < 0 || nums[i] > 50 {
			return -3
		}
	}

	// check all slice items
	// it not equal to val, inc non-val counter
	// else replace to spec_symbol
	nonValCounter := 0
	for i := range len(nums) {
		if nums[i] != val {
			nonValCounter++
		}
	}

	if nonValCounter == 0 { //meaning all slices values equal to val
		return nonValCounter
	}

	// replace all items equal to val to special symbol, bigger then all others
	// sort array in asc order
	specSymbol := 101
	for i := range len(nums) {
		if nums[i] == val {
			nums[i] = specSymbol
		}
	}
	slices.Sort(nums)

	return nonValCounter
}

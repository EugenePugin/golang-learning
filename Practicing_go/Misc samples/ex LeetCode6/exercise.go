// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
package ex6

import (
	"fmt"
	"math"
)

// func ifSeen(mapOfItems map[int]int, item int) bool {
// 	// fmt.Println(mapOfItems, item)
// 	for i := range mapOfItems {
// 		if item == mapOfItems[i] {
// 			// fmt.Println("Bingo")
// 			return true
// 		}
// 	}
// 	return false
// }

// func removeDuplicates(nums []int) int {
// 	//input validation:
// 	if len(nums) < 1 || len(nums) > 30000 {
// 		// print("Input validation failed")
// 		return -1
// 	}
// 	for i := range nums {
// 		if math.Abs(float64(nums[i])) > 100 {
// 			// print("Input validation failed")
// 			return -1
// 		}
// 	}
// 	//TODO

// 	mapOfItems := make(map[int]int, 0)
// 	unique_items_cnt := 1
// 	mapOfItems[0] = nums[0]
// 	originalSliceLen := len(nums)

// 	for i := 1; i < len(nums); i++ {
// 		// fmt.Println("Step ", i, ": checking ", nums[i], " unique_items_cnt:", unique_items_cnt)
// 		if ifSeen(mapOfItems, nums[i]) {
// 			nums = slices.Delete(nums, i, i+1)
// 			i--
// 		} else {
// 			unique_items_cnt++
// 			mapOfItems[unique_items_cnt] = nums[i]
// 		}
// 	}

// 	// fill zeroes to keep the origina len
// 	// fmt.Println(nums, originalSliceLen, unique_items_cnt)
// 	for i := 0; i < originalSliceLen-unique_items_cnt; i++ {
// 		nums = append(nums, 1)
// 		nums[len(nums)-1] = 0
// 	}
// 	// fmt.Println(nums)

// 	return unique_items_cnt // number of unique items
// }

func removeDuplicates(nums []int) int {
	//input validation:
	if len(nums) < 1 || len(nums) > 30000 {
		// print("Input validation failed")
		return -1
	}
	for i := range nums {
		if math.Abs(float64(nums[i])) > 100 {
			// print("Input validation failed")
			return -1
		}
	}

	unique_items_cnt := 1
	previous := nums[0]
	max_item_to_check := len(nums)
	for i := 1; i < max_item_to_check; i++ {
		// fmt.Print(nums)
		fmt.Println("Step ", i, ": checking ", nums[i], ".... unique_items_cnt:", unique_items_cnt, "previous:", previous, "max_item_to_check:", max_item_to_check)
		if previous == nums[i] {
			// fmt.Println("Dup discovered - shift to it all the items from the right")
			max_item_to_check--
			for j := i; j < max_item_to_check; j++ {
				nums[j] = nums[j+1]
			}
			i--
			// fmt.Println(nums)
		} else {
			unique_items_cnt++
			previous = nums[i]
		}
	}
	fmt.Print("...", nums, unique_items_cnt, "...")
	return unique_items_cnt // number of unique items
}

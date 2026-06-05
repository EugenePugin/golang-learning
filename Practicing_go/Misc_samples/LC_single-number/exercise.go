// https://leetcode.com/problems/single-number/description/?envType=problem-list-v2&envId=w4eezmsm
// https://leetcode.com/problems/single-number-ii/description/
package LC_single_number

// func singleNumber_v1(nums []int) int {
// 	// fmt.Println(nums)
// 	digits := make(map[int]int)
// 	for i := range nums {
// 		// fmt.Println("i:", i, "nums[i]:", nums[i])
// 		if _, ok := digits[nums[i]]; ok == false {
// 			// fmt.Println("new!")
// 			digits[nums[i]]++
// 		} else {
// 			// fmt.Println("already seen - deleting")
// 			delete(digits, nums[i])
// 		}
// 	}

// 	if len(digits) == 1 {
// 		for key := range digits {
// 			return key
// 		}
// 	}
// 	// fmt.Println(digits)
// 	return -1 // should not happen
// }

// func singleNumber_v2(nums []int) int {
// 	fmt.Println(nums)
// 	digits := make(map[int]int)
// 	for i := range nums {
// 		fmt.Println("i:", i, "nums[i]:", nums[i])
// 		if _, ok := digits[nums[i]]; ok == false {
// 			fmt.Println("new!")
// 			digits[nums[i]]++
// 		} else {
// 			fmt.Println("already seen")
// 			digits[nums[i]]++
// 			if digits[nums[i]] == 3 {
// 				fmt.Println("and now deleting")
// 				delete(digits, nums[i])
// 			}
// 		}
// 	}

// 	if len(digits) == 1 {
// 		for key := range digits {
// 			return key
// 		}
// 	}
// 	// fmt.Println(digits)
// 	return -1 // should not happen
// }

func singleNumber(nums []int) int {
	// fmt.Println(nums)
	digits := make(map[int]int)
	for i := range nums {
		// fmt.Println("i:", i, "nums[i]:", nums[i])
		if _, ok := digits[nums[i]]; ok == false {
			// fmt.Println("new!")
			digits[nums[i]]++
		} else {
			// fmt.Println("already seen")
			digits[nums[i]]++
			if digits[nums[i]] == 3 {
				// fmt.Println("and now deleting")
				delete(digits, nums[i])
			}
		}
	}

	// if len(digits) == 1 {
	for key := range digits {
		return key
	}
	// }
	// fmt.Println(digits)
	return -1 // should not happen
}

// https://leetcode.com/problems/majority-element/description/
package lc9

import (
	"math"
)

func majorityElement(nums []int) int {
	// fmt.Println(nums)
	nums_len := len(nums)
	// fmt.Println(nums_len)
	// input validity check

	// special case
	if 1 == nums_len {
		return nums[0]
	}
	// logic itself

	// func findNumOfOccurences(nums []int) int (

	// )

	// create map with unique digits (key) and number of occurences (value)
	// find the max number of occurences
	// confirm it is a majority
	numsMap := make(map[int]int)
	numsMap[nums[0]] = 1
	//the level of occurrences needs to be exceeded to be considered as a majority

	var signOfMajority int
	if 0 == nums_len%2 {
		signOfMajority = 1 + nums_len/2
	} else {
		signOfMajority = int(math.Ceil(float64(nums_len) / 2))
	}
	// fmt.Println("signOfMajority:", signOfMajority)
	// fmt.Println(numsMap)

	for i := 1; i < nums_len; i++ {
		// fmt.Println("Step ", i)
		value, exists := numsMap[nums[i]]
		// fmt.Println(value, exists, nums[i])
		if exists {
			// fmt.Println("Value of ", nums[i], " was disovered at the map. Need to inc its value")
			numsMap[nums[i]] = value + 1
		} else {
			// fmt.Println("Value of ", nums[i], " was not disovered at the map. Need to add")
			numsMap[nums[i]] = 1
		}
		// fmt.Println(numsMap)
		if numsMap[nums[i]] == signOfMajority {
			return nums[i]
		}
	}
	// fmt.Println(numsMap)

	return 0 //should not be possible as long majority exists
}

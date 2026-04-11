package ex3

import (
	"fmt"
	"slices"
)

// https://leetcode.com/problems/two-sum/
// 2248. Intersection of Multiple Arrays
// Easy
// Topics
// premium lock icon
// Companies
// Hint
// Given a 2D integer array nums where nums[i] is a non-empty array of distinct positive integers, return the list of integers that are present in each array of nums sorted in ascending order.

// Example 1:

// Input: nums = [[3,1,2,4,5],[1,2,3,4],[3,4,5,6]]
// Output: [3,4]
// Explanation:
// The only integers present in each of nums[0] = [3,1,2,4,5], nums[1] = [1,2,3,4], and nums[2] = [3,4,5,6] are 3 and 4, so we return [3,4].
// Example 2:

// Input: nums = [[1,2,3],[4,5,6]]
// Output: []
// Explanation:
// There does not exist any integer present both in nums[0] and nums[1], so we return an empty list [].

// Constraints:

// 1 <= nums.length <= 1000
// 1 <= sum(nums[i].length) <= 1000
// 1 <= nums[i][j] <= 1000
// All the values of nums[i] are unique.

func removeDuplicates(slice []int) []int {
	seen := make(map[int]bool) // ключ — число, значение — флаг «видели ли его»
	result := []int{}

	for _, num := range slice {
		if !seen[num] { // если ещё не видели это число
			seen[num] = true             // отмечаем, что видели
			result = append(result, num) // добавляем в результат
		}
	}
	return result
}

const two_d_array_line_max_size int = 1000           //TODO: to change to 1000
const two_d_array_total_lines_count_limit int = 1000 //TODO: to change to 1000

func validateInputs(nums [][]int) bool {
	//TODO: to change to pass a pointer
	var errorCode int
	// Validation1: all lines are non-empty
	rows_count := len(nums)
	// fmt.Println("rows_count:", rows_count)
	if 0 == rows_count {
		errorCode = 10 //empty 2D-array
	} else {
		for i := 0; i < rows_count; i++ {
			// fmt.Println("Line ", 1, "length: ", len(nums[i]))
			if 0 == len(nums[i]) {
				errorCode = 11 //any line from 2D-array is empty
				break
			}
		}
	}

	if 10 == errorCode || 11 == errorCode {
		switch errorCode {
		case 10:
			fmt.Println("empty 2D-array")
		case 11:
			fmt.Println("any line from 2D-array is empty")
		}
		return false
	}

	// Validation2: all the values of nums[i] are unique
	for i := 0; i < rows_count; i++ {
		original_len := len(nums[i])
		// fmt.Println("original_len:", original_len)
		distinct_only_len := len(removeDuplicates(nums[i]))
		// fmt.Println("distinct_only_len:", distinct_only_len)
		if distinct_only_len != original_len {
			errorCode = 20
			break
		}
	}

	if 20 == errorCode {
		switch errorCode {
		case 20:
			fmt.Println("all the values at any line of nums must be unique")
		}
		return false
	}

	// Validation3: all numbers are positive integers
	for i := range rows_count {
		for j := 0; j < len(nums[i]); j++ {
			// fmt.Println("Cell: [", i,"][",j,"]:",nums[i][j])
			if 0 > nums[i][j] {
				errorCode = 30
				break
			}
		}
	}

	if 30 == errorCode {
		switch errorCode {
		case 30:
			fmt.Println("all numbers must be positive integers")
		}
		return false
	}

	// Validation4: line lentgths: 1 <= nums.length <= 1000
	for i := range rows_count {
		columns_count := len(nums[i])
		// fmt.Println("Line", i, ":", columns_count)
		if columns_count > two_d_array_line_max_size {
			errorCode = 40
			break
		}
	}

	if 40 == errorCode {
		switch errorCode {
		case 40:
			fmt.Println("all the lines length must be below 1000")
		}
		return false
	}

	// Validation5: 1 <= sum(nums[i].length) <= 1000
	sum := 0
	for i := range rows_count {
		sum += len(nums[i])
		// fmt.Println("sum=", sum)
		if sum >= two_d_array_total_lines_count_limit {
			errorCode = 50
			break
		}
	}

	if 50 == errorCode {
		switch errorCode {
		case 50:
			fmt.Println("sum of the lines must be under 1000")
		}
		return false
	}
	// Validation6: 1 <= nums[i][j] <= 1000
	for i := range rows_count {
		for j := 0; j < len(nums[i]); j++ {
			// fmt.Println("nums[", i, "][", j, "]", nums[i][j])
			if nums[i][j] > 1000 {
				errorCode = 60
				break
			}
		}
	}

	if 60 == errorCode {
		switch errorCode {
		case 60:
			fmt.Println("all the cell numbers must be under 1000")
		}
		return false
	}
	return true
}

func ifDiscovered(value int, slice []int) bool {
	for i := range len(slice) {
		if value == slice[i] {
			return true
		}
	}
	return false
}

func intersection(nums [][]int) []int {
	var result []int
	fmt.Println(nums)
	if !validateInputs(nums) {
		fmt.Println("Input validation: FAILED - function aborted")
		return result
	}
	fmt.Println("Input validation: PASSED")

	// for each element of the line1, check the occurence at all other lines
	// sort the resulting slice in ascending order

	intersectionSlice := make([]int, 0)
	// fmt.Println("intersectionSlice:", intersectionSlice)

	// check whether all values of tmpSlice discovered at all other lines
	for i := range nums[0] {
		numberOfOccurence := 0
		for j := 1; j < len(nums); j++ {
			// fmt.Println("Checking for ", nums[0][i], " at line", j)
			if true == ifDiscovered(nums[0][i], nums[j]) {
				// fmt.Println("Bingo! ", nums[0][i])
				numberOfOccurence++
			} else {
				// fmt.Println("Value ", nums[0][i], "was not discovered at line", j)
			}
			// fmt.Println("numberOfOccurence:", numberOfOccurence)
			if numberOfOccurence == len(nums)-1 {
				// fmt.Println("Voila", numberOfOccurence)
				intersectionSlice = append(intersectionSlice, nums[0][i])
			}
		}
	}

	// fmt.Println("intersectionSlice:", intersectionSlice)
	// and do sort in asc

	slices.Sort(intersectionSlice)
	// fmt.Println("intersectionSlice:", intersectionSlice)

	result = intersectionSlice
	return result
}

/// how to run it
// func main() {

// 	start := time.Now()
// 	var nums [][]int
// 	nums = [][]int{{3, 1, 2}, {1, 2, 3}, {3, 4, 2, 15, 1}, {3, 1, 2, 15, 5}}

// 	// nums = [][]int{{3, 1, 2, 4, 5}, {}, {3, 4, 5, 6}}

// 	result := intersection(nums)
// 	elapsed := time.Since(start)
// 	fmt.Println(result)

// 	fmt.Printf("Время выполнения: %s\n", elapsed)

// }

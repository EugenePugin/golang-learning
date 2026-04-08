package main

import "fmt"

// https://leetcode.com/problems/intersection-of-two-arrays/description/
// 349. Intersection of Two Arrays
// Easy
// Topics
// premium lock icon
// Companies
// Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.

// Example 1:

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]
// Example 2:

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.

// Constraints:

// 1 <= nums1.length, nums2.length <= 1000
// 0 <= nums1[i], nums2[i] <= 1000

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

func intersection(nums1 []int, nums2 []int) []int {
	var result []int

	// Input parameters check
	errorCode := 0

	//input arrays size len checks
	len1 := len(nums1)
	len2 := len(nums2)
	switch {
	case len1 < 1:
		errorCode = 1
	case len2 < 1:
		errorCode = 2
	case len1 > 1000:
		errorCode = 3
	case len2 > 1000:
		errorCode = 4
	}

	if errorCode > 0 {
		switch errorCode {
		case 1:
			fmt.Println("len1 <1")
		case 2:
			fmt.Println("len2 <1")
		case 3:
			fmt.Println("len1 > 1000")
		case 4:
			fmt.Println("len12> 1000")
		}
		return result
	}

	//input arrays members checks
	// fmt.Println("nums1:", nums1)
	nums1u := removeDuplicates(nums1)
	// fmt.Println("nums1u:", nums1u)
	// fmt.Println("nums2:", nums2)
	nums2u := removeDuplicates(nums2)
	// fmt.Println("nums2u:", nums2u)
	for i := 0; i < len(nums1u); i++ {
		// fmt.Println("Assessing ", nums1u[i])
		switch {
		case nums1u[i] < 0:
			errorCode = 5
		case nums1u[i] > 1000:
			errorCode = 6
		}
		if errorCode == 5 || errorCode == 6 {
			break
		}
	}

	if errorCode == 5 || errorCode == 6 {
		switch errorCode {
		case 5:
			fmt.Println("nums1u < 0")
		case 6:
			fmt.Println("nums1u > 1000")
		}
		return result
	}

	for j := 0; j < len(nums2u); j++ {
		switch {
		case nums2u[j] < 0:
			errorCode = 7
		case nums2u[j] > 1000:
			errorCode = 8
		}
		if errorCode == 7 || errorCode == 8 {
			break
		}
	}

	if errorCode == 7 || errorCode == 8 {
		switch errorCode {
		case 7:
			fmt.Println("nums2u < 0")
		case 8:
			fmt.Println("nums2u > 1000")
		}
		return result
	}

	// logic itself
	// create bucket for resulting array

	// intersectionSlice := make([]int, max(len(nums1u), len(nums2u)))
	intersectionSlice := make([]int, min(len(nums1u), len(nums2u)))

	intersectionSliceIndex := 0
	// loop slice1 to discover elements of slice2 => once discover, add to the final slice
	for i := 0; i < len(nums1u); i++ {
		// fmt.Println("Consider ", nums1u[i], "from Slice1")
		for j := 0; j < len(nums2u); j++ {
			// fmt.Println("\tand ", nums2u[j], "from Slice2")
			if nums1u[i] == nums2u[j] {
				// fmt.Println("Bingo! intersectionSliceIndex=", intersectionSliceIndex)
				intersectionSlice[intersectionSliceIndex] = nums1u[i]
				intersectionSliceIndex++
			}
		}
	}
	// fmt.Println("intersectionSliceIndex:", intersectionSliceIndex)
	// fmt.Println("intersectionSlice:", intersectionSlice)

	if intersectionSliceIndex == 0 {
		// fmt.Println("No intersection discovered! ")
		// fmt.Println("result:", result)
		return result
	} else {
		// fmt.Println("Intersection discovered! ")
	}

	cellsToTruncate := len(intersectionSlice) - intersectionSliceIndex
	// fmt.Println("cellsToTruncate:", cellsToTruncate)

	truncated := intersectionSlice[:len(intersectionSlice)-cellsToTruncate]
	// fmt.Println("truncated:", truncated)

	result = removeDuplicates(truncated)
	// fmt.Println("result:", result)

	return result
}

// func discoverIntersection(nums1slicePtr, nums2slicePtr *[]int) {
// 	//
// 	nums1slice := *nums1slicePtr
// 	nums2slice := *nums2slicePtr
// 	// fmt.Println("Array1:")
// 	// for i1, value1 := range nums1slice {
// 	// 	fmt.Printf("Index %d: %d\n", i1, value1)
// 	// }
// 	// fmt.Println("Array2:")
// 	// for i2, value2 := range nums2slice {
// 	// 	fmt.Printf("Index %d: %d\n", i2, value2)
// 	// }

// 	// create bucket for resulting array
// 	lenIntersectionSlice := min(len(nums1slice), len(nums2slice))
// 	// fmt.Println("IntersectionSlice len = ", lenIntersectionSlice)
// 	intersectionSlice := make([]int, lenIntersectionSlice)
// 	// fmt.Println("intersectionSlice:")
// 	// for i3, value3 := range intersectionSlice {
// 	// 	fmt.Printf("Index %d: %d\n", i3, value3)
// 	// }

// 	intersectionSliceIndex := 0
// 	// loop slice1 to discover elements of slice2 => once discover, add to the final slice
// 	for i := 0; i < len(nums1slice); i++ {
// 		// fmt.Println("Consider ", nums1slice[i], "from Slice1")
// 		for j := 0; j < len(nums2slice); j++ {
// 			// fmt.Println("\tand ", nums2slice[j], "from Slice2")
// 			if nums1slice[i] == nums2slice[j] {
// 				// fmt.Println("Bingo! intersectionSliceIndex=", intersectionSliceIndex)
// 				intersectionSlice[intersectionSliceIndex] = nums1slice[i]
// 				intersectionSliceIndex++
// 			}
// 		}
// 	}

// 	// fmt.Println("Resulting intersectionSlice:")
// 	// for i, v := range intersectionSlice {
// 	// 	fmt.Printf("Index %d: %d\n", i, v)
// 	// }
// 	// fmt.Println("intersectionSlice:", intersectionSlice)

// 	unique := removeDuplicates(intersectionSlice)
// 	fmt.Println("unique:", unique)

// 	// return -1 // stub for now
// }

// https://leetcode.com/problems/merge-sorted-array/description/
package LCmsa

import (
	"slices"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := range nums2 {
		nums1[m+i] = nums2[i]
	}
	n = 0 //useless action just to suppress the warning
	slices.Sort(nums1)
	// fmt.Println(nums1, n)
}

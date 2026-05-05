package LCmsa

import (
	"fmt"
	"math"
)

func mergeTestHelper(nums1 []int, m int, nums2 []int, n int) []int {
	// resultSlice := slices.Clone(nums1)
	// input validation
	if m < 0 ||
		n > 200 ||
		m+n < 1 ||
		m+n > 200 {
		return nil
	}

	for i := range nums1 {
		if math.Abs(float64(nums1[i])) > math.Pow10(9) {
			return nil
		}
	}
	for i := range nums2 {
		if math.Abs(float64(nums2[i])) > math.Pow10(9) {
			return nil
		}
	}

	merge(nums1, m, nums2, n)
	// fmt.Println(resultSlice, nums1)
	// resultSlice = []int{1, 2, 2, 3, 5, 6}
	return nums1
}
func main() {
	fmt.Println("Hey!")

	m := 3
	n := 3

	nums1 := make([]int, m+n)
	nums2 := make([]int, n)

	nums1 = []int{1, 2, 3, 0, 0, 0}
	nums2 = []int{2, 5, 6}

	// fmt.Println("Before:\t", nums1, nums2)
	mergeTestHelper(nums1, m, nums2, n)
	// fmt.Println(mergeTestHelper(nums1, m, nums2, n))

}

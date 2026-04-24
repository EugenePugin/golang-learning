package LC100

import "fmt"

func main() {
	fmt.Println("Hey!")

	nums_array := [...]int{3, 2, 2, 3} // Input array
	nums := make([]int, len(nums_array))
	nums = nums_array[:]
	val := 3 // Value to remove
	// var expectedNums []int   // The expected answer with correct length.
	// It is sorted with no values equaling val.
	fmt.Println(nums, val)
	k := removeElement(nums, val) // Calls your implementation
	fmt.Println(k, nums)
}

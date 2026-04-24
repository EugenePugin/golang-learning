// create two singly linked lists
// create a singly linked list, which contains a sum of value in two lists of equal size

package lists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// array1 := [...]int{}
	// array2 := [...]int{}
	// expected []
	// array1 := [...]int{}
	// array2 := [...]int{1}
	// expected [1]
	// array1 := [...]int{1, 2, 4}
	// array2 := [...]int{1, 3, 4}
	// expected [1,1,2,3,4,4]

	// array1 := [...]int{1, 2, 4}
	// array2 := [...]int{1, 3, 4}
	// expected [1,1,2,3,4,4]

	array1 := [...]int{1, 2, 4}
	array2 := [...]int{5}
	// expected [1,5,2,4]

	slice1 := make([]int, len(array1))
	slice1 = array1[:]
	// fmt.Println(slice1)
	list1 := reverseTheWholeList(IntSinglyListBySlice(slice1))
	DisplayList(list1)

	slice2 := make([]int, len(array2))
	slice2 = array2[:]
	// fmt.Println(slice2)
	list2 := reverseTheWholeList(IntSinglyListBySlice(slice2))
	DisplayList(list2)

	merged_list := mergeTwoLists(list1, list2)
	DisplayList(merged_list)
	fmt.Println(list1, list2, merged_list)
}

package lists

import "fmt"

/*
type ListNode struct {
	Val  int
	Next *ListNode
}
*/
func AddToTop(head *ListNode, value int) *ListNode {
	newNode := &ListNode{Val: value, Next: head}
	return newNode
}

func IntSinglyListBySlice(slice []int) *ListNode {
	var head *ListNode = nil
	for _, val := range slice {
		head = AddToTop(head, val)
	}
	return head
}

func DisplayList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Println(current.Val, "->")
		current = current.Next
	}
	fmt.Println()
}

func reverseTheWholeList(list *ListNode) *ListNode {
	var head *ListNode = nil
	curList := list
	for curList != nil {
		head = AddToTop(head, curList.Val)
		curList = curList.Next
	}
	return head
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// input validation
	if list1 == nil && list2 != nil {
		return list2
	}
	if list2 == nil && list1 != nil {
		return list1
	}
	//
	var head *ListNode = nil // head of merged list
	list1cur := list1
	list2cur := list2
	for list1cur != nil || list2cur != nil {
		fmt.Println(list1cur, list2cur)
		switch {
		case (list2cur == nil):
			{
				head = AddToTop(head, list1cur.Val)
				list1cur = list1cur.Next
			}
		case (list1cur == nil):
			{
				head = AddToTop(head, list2cur.Val)
				list2cur = list2cur.Next
			}
		case (list1cur.Val <= list2cur.Val):
			{
				fmt.Println("a")
				head = AddToTop(head, list1cur.Val)
				head = AddToTop(head, list2cur.Val)
				list1cur = list1cur.Next
				list2cur = list2cur.Next
			}
		case (list1cur.Val > list2cur.Val):
			{
				fmt.Println("b")

				head = AddToTop(head, list2cur.Val)
				head = AddToTop(head, list1cur.Val)
				list1cur = list1cur.Next
				list2cur = list2cur.Next
			}
		}

	}

	return reverseTheWholeList(head)
}

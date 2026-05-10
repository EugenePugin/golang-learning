package LCpl

import (
	"fmt"
)

/**
 * Definition for singly-linked list. */
type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (list *LinkedList) Add(value int) {
	newListNode := &ListNode{Val: value}
	// Если список пуст, новый элемент становится головой
	if list.head == nil {
		list.head = newListNode
		return
	}
	// Иначе находим последний элемент и добавляем к нему новый
	current := list.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newListNode
}

func (list *LinkedList) Println() {
	current := list.head
	for current.Next != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println(current.Val)
}

func sliceToSinglyLinkedList(slice []int) *LinkedList {
	var linkedList LinkedList
	for i := range slice {
		linkedList.Add(slice[i])
	}
	return &linkedList
}

func singlyLinkedListToSlice(linkedList *LinkedList) []int {
	slice := make([]int, 0)
	current := linkedList.head
	for current.Next != nil {
		slice = append(slice, current.Val)
		current = current.Next
	}
	slice = append(slice, current.Val)
	return slice
}

func main() {

	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}

	fmt.Println(addTwoValues(l1, l2))
	// slice := singlyLinkedListToSlice(l2List)
	// fmt.Println(slice)

	// var l2List LinkedList
	// for i := range l2 {
	// 	l2List.Add(l2[i])
	// }
	// l2List.Println()

}

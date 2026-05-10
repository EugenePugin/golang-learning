// https://leetcode.com/problems/zigzag-conversion/description/

package LCpl

import (
	"fmt"
	"math"
	"slices"
)

func addTwoValues(l1 []int, l2 []int) []int { // for unit-testing

	// input validation
	if len(l1) == 0 ||
		len(l1) > 100 ||
		len(l2) == 0 ||
		len(l2) > 100 {
		return nil
	}
	for i := range l1 {
		if l1[i] < 0 || l1[i] > 9 {
			return nil
		}
	}
	for i := range l2 {
		if l2[i] < 0 || l2[i] > 9 {
			return nil
		}
	}

	// slices to linked lists conversion

	l1reverse := slices.Clone(l1)
	slices.Reverse(l1reverse)
	l1List := sliceToSinglyLinkedList(l1reverse)
	// fmt.Println(l1)
	l1List.Println()

	l2reverse := slices.Clone(l2)
	slices.Reverse(l2reverse)
	l2List := sliceToSinglyLinkedList(l2reverse)
	// fmt.Println(l2)
	l2List.Println()

	// summing itself
	summand1 := l1List.head
	summand2 := l2List.head
	var lastLap1, lastLap2 bool
	var counter int
	var sum *ListNode
	var trailingOne bool
	result := make([]int, max(len(l1), len(l2)))
	for {
		// fmt.Println("counter:", counter) //dbg only
		if lastLap1 || lastLap2 {
			// fmt.Println("exit from the loop")
			break
		}
		if summand1.Next == nil {
			lastLap1 = true
			// fmt.Println("lastLap1 set")
		}
		if summand2.Next == nil {
			lastLap2 = true
			// fmt.Println("lastLap2 set")
		}
		sum = addTwoNumbers(summand1, summand2)
		if trailingOne {
			fmt.Println("trailingOne time")
			sum.Val++
			trailingOne = false
		}
		if sum.Val >= 10 {
			fmt.Println("boop")
			sum.Val -= 10
			trailingOne = true
		}
		// fmt.Println(summand1.Val, summand2.Val, sum.Val)

		result[counter] = sum.Val
		summand1 = summand1.Next
		summand2 = summand2.Next
		counter++
	}

	fmt.Println("boop:", result)
	tmpSlice := make([]int, int(math.Abs(float64(len(l1)-len(l2)))))

	if lastLap1 {
		// fmt.Println(l1, l2, tmpSlice)
		for i := range tmpSlice {
			tmpSlice[i] = l2[i]
		}
		slices.Reverse(tmpSlice)
		// fmt.Println(tmpSlice)
		copy(result[len(tmpSlice):], tmpSlice)
	} else if lastLap2 {
		fmt.Println(l1, l2, tmpSlice)
		for i := range tmpSlice {
			tmpSlice[i] = l1[i]
		}
		slices.Reverse(tmpSlice)
		// fmt.Println(tmpSlice)
		copy(result[len(tmpSlice):], tmpSlice)
	}
	fmt.Println(result)

	return result // tbd
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmpListNode := ListNode{}
	tmpListNode.Val = l1.Val + l2.Val //tbd
	return &tmpListNode
}

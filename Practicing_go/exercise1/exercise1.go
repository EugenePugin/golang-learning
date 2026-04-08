package exercise1

import "fmt"

// Count Odd Numbers in an Interval Range
// Given two non-negative integers low and high. Return the count of odd numbers between low and high (inclusive).

//  https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/description/?q=go+interview

// Example 1:

// Input: low = 3, high = 7
// Output: 3
// Explanation: The odd numbers between 3 and 7 are [3,5,7].
// Example 2:

// Input: low = 8, high = 10
// Output: 1
// Explanation: The odd numbers between 8 and 10 are [9].

// Constraints:

// 0 <= low <= high <= 10^9

func oddsNumberCounter1(low int, high int) int {
	oddsNumberCounter := 0
	for i := low; i <= high; i++ {
		//fmt.Println("i=", i)
		// if i%2 == 1 {
		if i&1 == 1 {
			// fmt.Println("Odd number discovered!")
			oddsNumberCounter++
		}
	}
	// fmt.Println("Odd number counter: ", oddsNumberCounter)
	return oddsNumberCounter
}

func oddsNumberCounter2(low int, high int) int {
	var firstOddNumber, lastOddNumber int
	// discover first odd
	if low&1 == 1 {
		firstOddNumber = low
	} else {
		firstOddNumber = low + 1
	}
	// discover last odd
	if high&1 == 1 {
		lastOddNumber = high
	} else {
		lastOddNumber = high - 1
	}

	// fmt.Println("firstOddNumber=", firstOddNumber, "lastOddNumber=", lastOddNumber)

	// calculate number of odd number based on arythmetic progression
	oddsNumberCounter := (lastOddNumber-firstOddNumber)/2 + 1
	return oddsNumberCounter
}

func countOdds(low int, high int) int {
	// input checks
	switch {
	case low < 0:
		{
			fmt.Println("Constraint violation: ", "low < 0", " - function aborted")
			return -1
		}
	case low > high:
		{
			fmt.Println("Constraint violation: ", "low > high", " - function aborted")
			return -2
		}
	case high > 1000000000:
		{
			fmt.Println("Constraint violation: ", "high exceeds 10^9", " - function aborted")
			return -3
		}
	}

	// logic itself
	oddsNumberCounter := oddsNumberCounter2(low, high)
	return oddsNumberCounter
}

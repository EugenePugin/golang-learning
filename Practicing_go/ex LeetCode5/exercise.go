// https://leetcode.com/problems/roman-to-integer/description/

package main

import (
	"fmt"
	"strings"
)

func confirmValidCharacters(s string) bool {
	validSymbols := []string{"I", "V", "X", "L", "C", "D", "M"}
	var validSymbolsLen uint
	for i := range len(validSymbols) {
		if strings.Count(s, validSymbols[i]) > 0 {
			validSymbolsLen += uint(strings.Count(s, validSymbols[i]))
		}
	}
	// fmt.Println("validSymbolsLen: ", validSymbolsLen)

	if validSymbolsLen != uint(len(s)) {
		// fmt.Println("Oops ")
		return false
	}
	return true
}

func romanToInt(s string) int {
	var result int
	fmt.Println(s)
	// input validation
	if len(s) > 15 || len(s) < 1 {
		// fmt.Println("Input validation failed: incorrect length")
		return -1
	}

	// ensure s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
	if !confirmValidCharacters(s) {
		// fmt.Println("Input validation failed: incorrect symbols detected")
		return -1
	}

	// logic itself
	count_1000 := strings.Count(s, "M")
	// fmt.Println("Ms:", count_1000)
	count_500 := strings.Count(s, "D")
	// fmt.Println("Ds:", count_500)
	count_100 := strings.Count(s, "C")
	// fmt.Println("Cs:", count_100)
	count_50 := strings.Count(s, "L")
	// fmt.Println("Ls:", count_50)
	count_10 := strings.Count(s, "X")
	// fmt.Println("Xs:", count_10)
	count_5 := strings.Count(s, "V")
	// fmt.Println("Vs:", count_5)
	count_1 := strings.Count(s, "I")
	// fmt.Println("Is:", count_1)

	// common case
	result = count_1000*1000 + count_500*500 + count_100*100 + count_50*50 + count_10*10 + count_5*5 + count_1
	// fmt.Println("result", result)

	var smallerSymbolindex, biggerSymbolIndex int
	//special cases
	if count_1 == 1 && count_5 >= 1 {
		smallerSymbolindex = strings.Index(s, "I")
		biggerSymbolIndex = strings.LastIndex(s, "V")
		if smallerSymbolindex+1 == biggerSymbolIndex { //smaller symbol from the left side
			result -= count_5*5 + count_1 //rework common case calculation
			count_5--
			count_1--
			result += count_5*5 + count_1 + (5 - 1)
		}
	} else {
		if count_1 == 1 && count_10 >= 1 {
			// fmt.Println("trace A... result=", result)
			smallerSymbolindex = strings.Index(s, "I")
			biggerSymbolIndex = strings.LastIndex(s, "X")

			if smallerSymbolindex+1 == biggerSymbolIndex { //smaller symbol from the left side
				result -= count_10*10 + count_1 //rework common case calculation
				count_10--
				count_1--
				result += count_10*10 + count_1 + (10 - 1)
			}
			// fmt.Println("trace A... result=", result)
		}
	}

	if count_10 == 1 && count_50 >= 1 {
		smallerSymbolindex = strings.Index(s, "X")
		biggerSymbolIndex = strings.LastIndex(s, "L")
		if smallerSymbolindex+1 == biggerSymbolIndex { //smaller symbol from the left side
			result -= count_50*50 + count_10*10 //rework common case calculation
			count_50--
			count_10--
			result += count_50*50 + count_10*10 + (50 - 10)
		}
	} else {
		if count_10 == 1 && count_100 >= 1 {
			// fmt.Println("trace B... result=", result)
			smallerSymbolindex = strings.Index(s, "X")
			biggerSymbolIndex = strings.LastIndex(s, "C")
			if smallerSymbolindex+1 == biggerSymbolIndex { //smaller symbol from the left side
				result -= count_100*100 + count_10*10 //rework common case calculation
				count_100--
				count_10--
				result += count_100*100 + count_10*10 + (100 - 10)
			}
			// fmt.Println("trace B... result=", result)
		}
	}

	if count_100 == 1 && count_500 >= 1 {
		smallerSymbolindex = strings.Index(s, "C")
		biggerSymbolIndex = strings.LastIndex(s, "D")
		if smallerSymbolindex+1 == biggerSymbolIndex { //smaller symbol from the left side
			// fmt.Println("trace...")
			result -= count_500*500 + count_100*100 //rework common case calculation
			count_500--
			count_100--
			result += count_500*500 + count_100*100 + (500 - 100)
		}
	} else {
		if count_100 == 1 && count_1000 >= 1 {
			// fmt.Println("trace...")

			smallerSymbolindex = strings.Index(s, "C")
			biggerSymbolIndex = strings.LastIndex(s, "M")
			// fmt.Println("trace.C.. result=", result, smallerSymbolindex, biggerSymbolIndex)
			// if smallerSymbolindex < biggerSymbolIndex { //smaller symbol from the left side
			if smallerSymbolindex+1 == biggerSymbolIndex { // smaller symbol from the left side
				// fmt.Println("trace...")

				result -= count_1000*1000 + count_100*100 //rework common case calculation
				count_1000--
				count_100--
				result += count_1000*1000 + count_100*100 + (1000 - 100)
			}
			//fmt.Println("trace... result=", result)
		}
	}

	return result
}

package main

import (
	"fmt"
	"strconv"
)

func one10g(dbgMode int) {
	var a, b int
	// var digitsA = make([]int, 0) // clice for value A digits
	// var digitsB = make([]uint, 0) // clice for value A digits
	// var digit int
	// var yearsCount uint
	// var accountBalance float64

	if dbgMode == 1 { //dbg mode
		fmt.Println("Hey, it is Debug mode")
		a = 564
		b = 8954
	} else { //release mode
		fmt.Scan(&a, &b)
		// 	fmt.Scan(&sampleValue)
		// 	if sampleValue < 10 {
		// 		continue
		// 	}
		// 	if sampleValue > 100 {
		// 		break
		// 	}
		// 	mySlice = append(mySlice, int(sampleValue))
		// 	sliceSize++
		// }
	}

	if a >= 10000 || b > 10000 {
		fmt.Println("Incorrect input data")
		return
	}

	// fmt.Println("a=", a, "b=", b)
	stringA := strconv.Itoa(a)
	runeASlice := []rune(stringA)
	sliceASize := uint(len(stringA))
	// fmt.Println(uint(runeASlice[0]-'0'), uint(runeASlice[1]-'0'), sliceASize)

	var i, j uint
	var uintASlice = make([]uint, 0)
	for i = 0; i < sliceASize; i++ {
		uintASlice = append(uintASlice, uint(runeASlice[i]-'0'))
		// fmt.Println(uintASlice[i])
	}

	stringB := strconv.Itoa(b)
	runeBSlice := []rune(stringB)
	sliceBSize := uint(len(stringB))

	var uintBSlice = make([]uint, 0)
	for i = 0; i < sliceBSize; i++ {
		uintBSlice = append(uintBSlice, uint(runeBSlice[i]-'0'))
		// fmt.Println(uintBSlice[i])
	}

	var uintBothABSlice = make([]uint, 0)
	var bothABSliceSize uint

	for i = 0; i < sliceASize; i++ {
		// fmt.Println("OUT: i=", i, "j=", j)
		if i > sliceASize {
			break
		}
		for j = 0; j < sliceBSize; j++ {
			// fmt.Println("IN: i=", i, "j=", j)
			if j > sliceBSize {
				break
			}
			if uintASlice[i] == uintBSlice[j] {
				// fmt.Println("Bingo!")
				uintBothABSlice = append(uintBothABSlice, uintASlice[i])
				bothABSliceSize++
			}
		}
		j = 0
	}

	// fmt.Println(bothABSliceSize)

	for i = 0; i < bothABSliceSize; i++ {
		stringToOutput := fmt.Sprintf("%d ", uintBothABSlice[i])
		fmt.Print(stringToOutput)
	}
}

func main() {
	one10g(0)
}

package main

import (
	"fmt"
)

func one12a() {
	var workArray [10]uint8
	var swap1, swap2, swap3 [2]uint8

	// if dbgMode == 1 { //dbg mode
	// 	workArray = [10]uint8{99, 151, 137, 71, 117, 187, 20, 93, 187, 67}
	// 	swap1 = [2]uint8{1, 2}
	// 	swap2 = [2]uint8{3, 5}
	// 	swap3 = [2]uint8{7, 8}
	// } else {

	// fmt.Println("Hey, I am release mode")
	for i := 0; i < 10; i++ {
		fmt.Scan(&workArray[i])
	}
	fmt.Scan(&swap1[0])
	fmt.Scan(&swap1[1])
	fmt.Scan(&swap2[0])
	fmt.Scan(&swap2[1])
	fmt.Scan(&swap3[0])
	fmt.Scan(&swap3[1])
	// }s

	// fmt.Println(workArray)
	tmp := workArray[swap1[0]]
	workArray[swap1[0]] = workArray[swap1[1]]
	workArray[swap1[1]] = tmp
	// fmt.Println(workArray)
	tmp = workArray[swap2[0]]
	workArray[swap2[0]] = workArray[swap2[1]]
	workArray[swap2[1]] = tmp
	// fmt.Println(workArray)
	tmp = workArray[swap3[0]]
	workArray[swap3[0]] = workArray[swap3[1]]
	workArray[swap3[1]] = tmp
	// fmt.Println(workArray)

	for i := 0; i < 10; i++ {
		if i == 9 {
			fmt.Print(workArray[i])
		} else {
			fmt.Print(workArray[i], " ")
		}
	}

	formattedString := fmt.Sprintf("%T", workArray)
	if formattedString == "[10]uint8" {
		fmt.Println(" type ok")
	}
}

func one12(dbgMode uint) {
	// var workArray [10]uint8
	// var swap1, swap2, swap3 [2]uint8
	var N int
	mySlice := make([]int, N) // Declares a nil slice of integers
	var sliceElement int

	if dbgMode == 1 { //dbg mode
		fmt.Println("Hey, I am debug mode")
		N = 5
		mySlice = append(mySlice, 41, -231, 24, 49, 6)
	} else {
		// fmt.Println("Hey, I am release mode")
		fmt.Scan(&N)
		for i := 0; i < N; i++ {
			fmt.Scan(&sliceElement)
			mySlice = append(mySlice, sliceElement)
		}
	}

	// input data check
	if N < 4 {
		fmt.Println("Input data are not valid")
		return
	}
	// fmt.Println(N)
	// for i := 0; i < N; i++ {
	// 	fmt.Print(mySlice[i])
	// }

	fmt.Print(mySlice[3])

}

func one12e() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	max := array[0]
	for i := 1; i < 5; i++ {
		if max < array[i] {
			max = array[i]
		}
	}
	// здесь ваш код
	fmt.Println(max)

}

func one12e(mode int) {
	var N, countOfPositives int
	mySlice := make([]int, N)
	var sliceElement int

	if mode == 1 {
		// fmt.Println("Hey, I am dbg mode")
		N = 5
		mySlice = append(mySlice, 1, 2, 3, -1, -4)

	} else {
		// fmt.Println("Hey, I am release mode")
		fmt.Scan(&N)
		for i := 0; i < N; i++ {
			fmt.Scan(&sliceElement)
			mySlice = append(mySlice, sliceElement)
		}
	}

	for i := 0; i < N; i++ {
		if mySlice[i] > 0 {
			countOfPositives++
		}
	}

	fmt.Println(countOfPositives)

}

package main

import "fmt"

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

func main() {
	one12e(0)
}

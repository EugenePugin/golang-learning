package main

import (
	"fmt"
)

func one13(mode int) {
	var n uint

	if mode == 1 {
		// fmt.Println("DBG mode welcomes you!")
		n = 2
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&n)
	}
	// input data check
	if n < 1 {
		fmt.Printf("Input data are out of range")
		return
	}
	//

	var fibonacchiNumber uint
	var prev1FibonacchiNumber, prev2FibonacchiNumber uint
	// var countInFibonacchi uint
	var counter uint
	counter = 0
	for {
		switch {
		case counter == 0:
			{
				fibonacchiNumber = 0
			}
		case counter == 1:
			{
				fibonacchiNumber = 1
			}
		case counter == 2:
			{
				fibonacchiNumber = 1
				prev1FibonacchiNumber = 1
				prev2FibonacchiNumber = 0
			}
		default:
			{
				fibonacchiNumber = prev1FibonacchiNumber + prev2FibonacchiNumber
			}
		}

		// fmt.Println(counter, fibonacchiNumber, n)

		if n == fibonacchiNumber {
			fmt.Println(counter)
			return
		} else {
			if n < fibonacchiNumber {
				fmt.Println(-1)
				return
			}
		}

		counter++
		prev2FibonacchiNumber = prev1FibonacchiNumber
		prev1FibonacchiNumber = fibonacchiNumber
	}
	// fmt.Println("n=", n, "countInFibonacchi=", countInFibonacchi)

}

func main() {
	// var i uint
	// for i = 0; i < 40; i++ {
	// 	one13(1, i)
	// }
	one13(0)
}

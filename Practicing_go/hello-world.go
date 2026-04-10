package main

import (
	"fmt"
)

func main() {
	n := 123456789

	removeOddNumbers := func(x int) int {
		//  на вход получает целое положительное число
		if x <= 0 {
			fmt.Println("Input validation failed")
			return 0
		}
		//возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0
		digitsReverseSlice := make([]uint, 0)
		for i := 0; x > 0; i++ {
			digitsReverseSlice = append(digitsReverseSlice, 1)
			digitsReverseSlice[i] = uint(x - 10*(x/10))
			x /= 10
		}
		// fmt.Println(digitsReverseSlice, len(digitsReverseSlice))

		digitsSlice := make([]uint, len(digitsReverseSlice))
		for i := range len(digitsSlice) {
			digitsSlice[len(digitsSlice)-i-1] = digitsReverseSlice[i]
		}
		// fmt.Println("digitsSlice", digitsSlice)

		digitsEvenOnly := make([]uint, 0)
		for i := range len(digitsSlice) {
			// fmt.Println("i:", i)
			if i%2 == 1 {
				// fmt.Println("index odd :", i)
				digitsEvenOnly = append(digitsEvenOnly, 1)
				// fmt.Println("index to write:", i-1)
				digitsEvenOnly[len(digitsEvenOnly)-1] = digitsSlice[i]

				// fmt.Println(digitsEvenOnly)
			}
		}

		// fmt.Println(digitsEvenOnly)

		// fmt.Println("digitsSliceRecovery", digitsSliceRecovery)

		var v1 int
		for i := range len(digitsEvenOnly) {
			v1 = int(digitsEvenOnly[uint(i)]) + 10*v1
		}
		// fmt.Println("Recovered", v1)

		return v1
	}

	removeOddNumbersClosure := func() int {
		//возвращает число того же типа в котором отсутствуют нечетные цифры и цифра 0
		digitsReverseSlice := make([]uint, 0)
		for i := 0; n > 0; i++ {
			digitsReverseSlice = append(digitsReverseSlice, 1)
			digitsReverseSlice[i] = uint(n - 10*(n/10))
			n /= 10
		}
		// fmt.Println(digitsReverseSlice, len(digitsReverseSlice))

		digitsSlice := make([]uint, len(digitsReverseSlice))
		for i := range len(digitsSlice) {
			digitsSlice[len(digitsSlice)-i-1] = digitsReverseSlice[i]
		}
		// fmt.Println("digitsSlice", digitsSlice)

		digitsEvenOnly := make([]uint, 0)
		for i := range len(digitsSlice) {
			// fmt.Println("i:", i)
			if i%2 == 1 {
				// fmt.Println("index odd :", i)
				digitsEvenOnly = append(digitsEvenOnly, 1)
				// fmt.Println("index to write:", i-1)
				digitsEvenOnly[len(digitsEvenOnly)-1] = digitsSlice[i]

				// fmt.Println(digitsEvenOnly)
			}
		}

		// fmt.Println(digitsEvenOnly)

		// fmt.Println("digitsSliceRecovery", digitsSliceRecovery)

		var v1 int
		for i := range len(digitsEvenOnly) {
			v1 = int(digitsEvenOnly[uint(i)]) + 10*v1
		}
		// fmt.Println("Recovered", v1)

		return v1
	}

	removeEvenNumbers := func(x int) int {
		//  на вход получает целое положительное число
		if x <= 0 {
			fmt.Println("Input validation failed: ", x)
			return 0
		}
		//возвращает число того же типа в котором отсутствуют четные цифры и цифра 0
		digitsReverseSlice := make([]uint, 0)
		for i := 0; x > 0; i++ {
			digitsReverseSlice = append(digitsReverseSlice, 1)
			digitsReverseSlice[i] = uint(x - 10*(x/10))
			x /= 10
		}

		digitsSlice := make([]uint, len(digitsReverseSlice))
		for i := range len(digitsSlice) {
			digitsSlice[len(digitsSlice)-i-1] = digitsReverseSlice[i]
		}

		digitsOddOnly := make([]uint, 0)
		for i := range len(digitsSlice) {
			// fmt.Println("i:", i)
			if i%2 == 0 {
				// fmt.Println("index odd :", i)
				digitsOddOnly = append(digitsOddOnly, 1)
				// fmt.Println("index to write:", i-1)
				digitsOddOnly[len(digitsOddOnly)-1] = digitsSlice[i]
			}
		}

		// fmt.Println(digitsOddOnly)

		// fmt.Println("digitsSliceRecovery", digitsSliceRecovery)

		var v1 int
		for i := range len(digitsOddOnly) {
			v1 = int(digitsOddOnly[uint(i)]) + 10*v1
		}
		// fmt.Println("Recovered", v1)

		return v1
	}

	fmt.Println(n)
	a := removeOddNumbersClosure()
	fmt.Println(n)

	fmt.Println(a)
	b := removeEvenNumbers(n) //expect 246
	fmt.Println(b)
	c := removeOddNumbers(n) //expect 246
	fmt.Println(c)

}

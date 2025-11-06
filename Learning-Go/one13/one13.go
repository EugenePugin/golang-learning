package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getDigitsString(n int) []int {
	// Convert the integer to a string
	s := strconv.Itoa(n)

	// Handle potential negative sign
	startIndex := 0
	if s[0] == '-' {
		startIndex = 1
	}

	var digits []int
	// Iterate over the string characters (runes)
	for i := startIndex; i < len(s); i++ {
		// Convert the character back to an integer
		digit, _ := strconv.Atoi(string(s[i]))
		digits = append(digits, digit)
	}

	return digits
}

func one13jj(mode int) {
	var a, b int
	// var mySlice []int

	if mode == 1 {
		// fmt.Println("DBG mode welcomes you!")
		a = -30
		b = 0
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&a, &b)
	}
	// input data check
	//
	var maxDividedBySeven int
	solutionFound := false
	for i := b; i >= a; i-- {
		// fmt.Println(a, b, i)
		if i%7 == 0 {
			maxDividedBySeven = i
			solutionFound = true
			break
		}
	}
	if !solutionFound {
		fmt.Println("NO")
	} else {
		fmt.Println(maxDividedBySeven)
	}
}

func one13hh(mode int) {
	var N uint
	// var mySlice []int

	if mode == 1 {
		// fmt.Println("DBG mode welcomes you!")
		N = 3456
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&N)
	}
	// input data check
	//

	var interimValue, sum uint
	interimValue = N
	for {
		sum = 0
		digits := getDigitsString(int(interimValue))
		for i := 0; i < len(digits); i++ {
			// fmt.Println(digits[i])
			sum += uint(digits[i])
		}
		// fmt.Println("sum: :", sum, "iteration=", i)
		interimValue = sum
		if interimValue < 10 {
			break
		}
	}
	digitRoot := interimValue
	fmt.Println(digitRoot)
}

func one13bb(mode int) {
	var N int
	if mode == 1 {
		fmt.Println("DBG mode welcomes you!")
		N = 981
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&N)
	}

	// input data check
	if !(N > 99 && N < 999) {
		fmt.Println("Input data are out of range")
		return
	}

	digits := getDigitsString(N)
	if digits[2] == 0 {
		fmt.Println("Input data are out of range")
		return
	}

	for i := len(digits) - 1; i >= 0; i-- {
		fmt.Print(digits[i])
	}
}

func one13a(mode int) {
	var N int
	if mode == 1 {
		fmt.Println("DBG mode welcomes you!")
		N = 987
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&N)
	}

	// input data check
	if !(N > 99 && N < 999) {
		fmt.Println("Input data are out of range")
		return
	}

	digits := getDigitsString(N)
	var sum int
	for i := 0; i < len(digits); i++ {
		sum += digits[i]
	}
	fmt.Println(sum)

}

func main() {
	one13a(0)
}

func one13gg(mode int) {
	var N uint
	var mySlice []int
	var sliceElement int

	if mode == 1 {
		// fmt.Println("DBG mode welcomes you!")
		N = 5
		// mySlice = make([]int, N)
		mySlice = append(mySlice, 1, 0, 100, 0, 12)
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&N)
		for i := 0; i < int(N); i++ {
			fmt.Scan(&sliceElement)
			mySlice = append(mySlice, sliceElement)
		}
	}
	// input data check
	// if !(a > 0 && b > 0) {
	// 	fmt.Println("Input data are out of range")
	// 	return
	// }

	minValue := mySlice[0]
	countOfMins := 1

	for i := 1; i < int(N); i++ {
		if mySlice[i] < minValue {
			minValue = mySlice[i]
			countOfMins = 1
		} else {
			if mySlice[i] == minValue {
				countOfMins++
			}
		}
	}

	fmt.Println(countOfMins)

}

func one13ee(mode int) {
	var a, b, c uint8

	if mode == 1 {
		fmt.Println("DBG mode welcomes you!")
		a = 3
		b = 0
		c = 5
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&a, &b, &c)
	}
	// input data check
	if !(a > 0 && b > 0 && c > 0) {
		fmt.Println("Input data are out of range")
		return
	}

	if a+b > c && a+c > b && b+c > a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}

}

func one13dd(mode int) {
	var a, b, c uint8

	if mode == 1 {
		fmt.Println("DBG mode welcomes you!")
		a = 3
		b = 0
		c = 5
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&a, &b, &c)
	}
	// input data check
	if !(a > 0 && b > 0 && c > 0) {
		fmt.Println("Input data are out of range")
		return
	}

	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}

}

func one13cc(mode int) {
	var k uint32
	const secondsInHour = 3600
	const secondsInMin = 60
	if mode == 1 {
		fmt.Println("DBG mode welcomes you!")
		k = 13257
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&k)
	}
	// input data check
	if !(k > 0 && k < 86399) {
		fmt.Println("Input data are out of range")
		return
	}
	hh := k / secondsInHour
	fmt.Print("It is ", hh, " hours ")
	mm := (k - hh*secondsInHour) / secondsInMin
	fmt.Println(mm, "minutes.")

	// digits := getDigitsString(N)
	// if digits[2] == 0 {
	// 	fmt.Println("Input data are out of range")
	// 	return
	// }

	// for i := len(digits) - 1; i >= 0; i-- {
	// 	fmt.Print(digits[i])
	// }
}

func korov_ending(n uint) string {
	const korov = "korov"
	const korova = "korova"
	const korovy = "korovy"
	korovX := "not set"

	switch {
	case n == 0 || n >= 5 && n <= 20:
		{
			korovX = korov
		}
	case n == 1:
		{
			korovX = korova
		}
	case n >= 2 && n < 5:
		{
			korovX = korovy
		}
	}
	return korovX
}

func one13m1m(mode int) {
	var n, m int

	if mode == 1 {
		// fmt.Println("DBG mode welcomes you!")
		n = 2343434
		m = 4
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&n, &m)
	}
	// input data check
	if n < 1 || m < 1 {
		fmt.Printf("Input data are out of range")
		return
	}
	// demount value to digits
	strSliceRaw := strings.Split(strconv.Itoa(n), "")
	// fmt.Println("DBG: ", strSliceRaw) // Output: [1 2 3 4 5]
	var strSliceProcessed []string
	for i := 0; i < len(strSliceRaw); i++ {
		// fmt.Println("DBG: ", m)
		if strconv.Itoa(m) != strSliceRaw[i] {
			strSliceProcessed = append(strSliceProcessed, strSliceRaw[i])
		}
	}

	// for i := 0; i < len(strSliceProcessed); i++ {
	// 	fmt.Println("DBG: ", strSliceProcessed[i])
	// }

	fmt.Println(strings.Join(strSliceProcessed, ""))
	// remove the 1st occurence of the specified digit
	// 	TODO

}

func one13k1k(mode int) {
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

func one13mm(mode int) {
	var n uint

	if mode == 1 {
		// fmt.Println("DBG mode welcomes you!")
		n = 50
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&n)
	}
	// input data check
	// if n >= 100 {
	// 	fmt.Printf("Input data are out of range")
	// 	return
	// }
	//
	var pow2result float64
	var i uint

	if n == 1 {
		fmt.Print(0)
		return
	}
	for i = 0; i <= n; i++ {
		pow2result = math.Pow(2.0, float64(i))
		if pow2result > float64(n) {
			break
		}
		fmt.Print(pow2result, " ")
	}

}
func one13kk(mode int, korovNum uint) {
	var n uint

	if mode == 1 {
		// fmt.Println("DBG mode welcomes you!")
		n = korovNum
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&n)
	}
	// input data check
	if n >= 100 {
		fmt.Printf("Input data are out of range")
		return
	}
	//
	m := n
	if n > 20 {
		tensNumber := (uint)(n / 10)
		// fmt.Println("tensNumber:", tensNumber)
		m = n - tensNumber*10
		// fmt.Println("and now:", m)
	}
	fmt.Println(n, korov_ending(m))

}

func sumInt(inputs ...int) (int, int) {

	// input data check
	// if n < 1 {
	// 	fmt.Printf("Input data are out of range")
	// 	return
	// }
	//

	amountOfArguments := len(inputs)

	var sumOfArguments int
	for i := 0; i < amountOfArguments; i++ {
		sumOfArguments += inputs[i]
	}

	return amountOfArguments, sumOfArguments

}

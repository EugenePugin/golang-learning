package one10a

import (
	"fmt"
	"math"
	"strconv"
)

func one10a() {

	var a, b int
	// tmp
	fmt.Scan(&a)
	fmt.Scan(&b)

	// a = 1
	// b = 5

	// input data check
	if !(a < b) {
		fmt.Println("A must be < then B")
		return
	}
	if !(a < 100) || !(b < 100) {
		fmt.Println("Both A and B must be below 100")
		return
	}

	sum := a
	// functional part
	for i := a + 1; i <= b; i++ {
		sum += i
	}

	fmt.Println(sum)

}

func one10b() {
	var sliceLen uint
	var index uint

	mySlice := make([]int, 0) // Declares a nil slice of integers
	dbgMode := 0

	if dbgMode == 1 { //dbg mode
		sliceLen = 5
		mySlice = append(mySlice, 38)
		mySlice = append(mySlice, 24)
		mySlice = append(mySlice, 800)
		mySlice = append(mySlice, 8)
		mySlice = append(mySlice, 16)
	} else { //release mode
		fmt.Scan(&sliceLen)
		for index = 0; index < sliceLen; index++ {
			var sliceElement int
			fmt.Scan(&sliceElement)
			mySlice = append(mySlice, sliceElement)
		}
	}

	// fmt.Println("sliceLen", sliceLen, "\nslice elements:")
	// for index = 0; index < sliceLen; i++ {
	// 	fmt.Println(mySlice[i])
	// }
	// fmt.Println("====")
	var sum int
	for index = 0; index < sliceLen; index++ {
		if mySlice[index] >= 10 &&
			mySlice[index] < 100 &&
			mySlice[index]%8 == 0 {
			sum += mySlice[index]
		}
	}
	fmt.Println(sum)
}

func one10с(dbgMode int) {
	// var sliceLen uint
	var index uint

	mySlice := make([]int, 0) // Declares a nil slice of integers
	// dbgMode := 0

	if dbgMode == 1 { //dbg mode
		// sliceLen = 5
		mySlice = append(mySlice, 1)
		mySlice = append(mySlice, 3)
		mySlice = append(mySlice, 3)
		mySlice = append(mySlice, 1)
		mySlice = append(mySlice, 0) //признак конца
	} else { //release mode
		// fmt.Scan(&sliceLen)
		var sliceElement int = -1
		for index = 0; sliceElement != 0; index++ {
			fmt.Scan(&sliceElement)
			mySlice = append(mySlice, sliceElement)

			// if sliceElement != 0 {
			// 	mySlice = append(mySlice, sliceElement)
			// }
		}
	}
	// for index = 0; mySlice[index] != 0; index++ {
	// 	fmt.Println(mySlice[index])
	// }
	// fmt.Println("====")
	var max int
	for index = 0; mySlice[index] != 0; index++ {
		if mySlice[index] > max {
			max = mySlice[index]
		}
	}
	var countOfMax int
	for index = 0; mySlice[index] != 0; index++ {
		if mySlice[index] == max {
			countOfMax++
		}
	}
	// fmt.Println(max)
	fmt.Println(countOfMax)
}

func one10d(dbgMode int) {
	var n, c, d uint
	var index uint

	if dbgMode == 1 { //dbg mode
		n = 20
		c = 3
		d = 5
	} else { //release mode
		fmt.Scan(&n, &c, &d)
	}
	// fmt.Println("n=", n, "c=", c, "d=", d)
	for index = c; index <= n; index++ {
		if index%c == 0 && index%d != 0 {
			fmt.Println(index)
			break
		}
	}
}

func one10e(dbgMode int) {
	var sampleValue int
	var mySlice = make([]int, 0) // Declares a nil slice of integers
	var sliceSize uint
	var index uint

	if dbgMode == 1 { //dbg mode
		// n = 20
		// c = 3
		// d = 5
	} else { //release mode
		for index = 0; ; index++ {
			fmt.Scan(&sampleValue)
			if sampleValue < 10 {
				continue
			}
			if sampleValue > 100 {
				break
			}
			mySlice = append(mySlice, int(sampleValue))
			sliceSize++
		}
	}

	for index = 0; index < sliceSize; index++ {
		fmt.Println(mySlice[index])
	}
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func one10f(dbgMode int) {
	var x, p, y uint
	var yearsCount uint
	var accountBalance float64

	if dbgMode == 1 { //dbg mode
		x = 100
		p = 10
		y = 200
	} else { //release mode
		fmt.Scan(&x, &p, &y)
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

	// fmt.Println("x=", x, "p=", p, "y", y)
	accountBalance = float64(x) //1st year
	for accountBalance < float64(y) {
		accountBalance *= (1 + float64(p)/100) //annual change
		// fmt.Println("After the year=", accountBalance)
		accountBalance = roundFloat(accountBalance, 2) // 12.35
		// fmt.Println("After the year=", accountBalance)
		yearsCount++
	}
	// fmt.Println("accountBalance=", accountBalance)
	fmt.Println(yearsCount)
}

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

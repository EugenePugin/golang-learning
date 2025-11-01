package main

import (
	"fmt"
	"strconv"
)

func one9a() {
	var a int
	var first, second, third int
	fmt.Scan(&a)
	first = a / 100
	second = (a - first*100) / 10
	third = a - first*100 - second*10
	// fmt.Println(first, second, third)

	if first == second || first == third || second == third {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func one9b() {
	var a int //На вход дается натуральное число, не превосходящее 10000.
	fmt.Scan(&a)
	first_digit := a
	//	checking for the input range is intentionally skipped here
	switch {
	case a == 10000:
		first_digit = 1
	case a > 1000:
		first_digit = a / 1000
	case a > 100:
		first_digit = a / 100
	case a > 10:
		first_digit = a / 10
	}
	fmt.Println(first_digit)
}

func one9c() {
	var sixDigitsStr string
	// TMP
	fmt.Scan(&sixDigitsStr)
	// TMP
	//sixDigitsStr = "234567"

	runes := []rune(sixDigitsStr)
	digit1 := string(runes[0])
	digit2 := string(runes[1])
	digit3 := string(runes[2])
	digit4 := string(runes[3])
	digit5 := string(runes[4])
	digit6 := string(runes[5])

	//fmt.Println("Result: ", sixDigitsStr, "contains: ", digit1, digit2, digit3, digit4, digit5, digit6)

	digit1num, err1 := strconv.Atoi(digit1) // Atoi assumes base 10 and returns an int
	digit2num, err2 := strconv.Atoi(digit2) // Atoi assumes base 10 and returns an int
	digit3num, err3 := strconv.Atoi(digit3) // Atoi assumes base 10 and returns an int
	digit4num, err4 := strconv.Atoi(digit4) // Atoi assumes base 10 and returns an int
	digit5num, err5 := strconv.Atoi(digit5) // Atoi assumes base 10 and returns an int
	digit6num, err6 := strconv.Atoi(digit6) // Atoi assumes base 10 and returns an int

	//	fmt.Println(digit1num,err1)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
		fmt.Println("Error converting string to int:")
	} // else {
	// 	//fmt.Println("Converted Int: ", digit1num, digit2num, digit3num, digit4num, digit5num, digit6num)
	// }

	if digit1num+digit2num+digit3num == digit4num+digit5num+digit6num {
		fmt.Println("YES") //Happy
	} else {
		fmt.Println("NO") //UnHappy
	}
}

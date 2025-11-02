package main

import (
	"fmt"
)

func ifLeapYear(year int) bool {

	// condition1 := false
	// condition2 := false
	// condition3 := true

	// switch {
	// case year%400 == 0:
	// 	condition1 = true
	// 	// fmt.Println("condition1+ - кратен 400")
	// 	// fallthrough
	// case year%4 == 0:
	// 	condition2 = true
	// 	// fmt.Println("condition2+ - кратен 4")
	// 	fallthrough
	// case year%100 == 0:
	// 	condition3 = false
	// 	// fmt.Println("condition3+ - кратен 100")
	// }
	// if condition1 || (condition2 && !condition3)

	if (year%400 == 0) || ((year%4 == 0) && (year%100 != 0)) {
		// fmt.Println(year, "YES") //высокосный
		return true //высокосный
	} else {
		// fmt.Println(year, "NO") // не высокосный
		return false //не высокосный
	}

}

func ifLeapYearUnitTest() {
	var testYears = [...]int{2025, 2000, 2004, 2008, 2012, 2016, 2020, 2024, 2028, 2027, 2026}
	var testIfLeap = [...]bool{false, true, true, true, true, true, true, true, true, false, false}
	for i := 0; i < len(testYears); i++ {
		if testIfLeap[i] == ifLeapYear(testYears[i]) {
			fmt.Println(testYears[i], "+")
		} else {
			fmt.Println(testYears[i], "-")
		}
	}
}

func main() {
	// ifLeapYearUnitTest()

	var yearNum int
	fmt.Scan(&yearNum)

	if ifLeapYear(yearNum) {
		fmt.Println("YES") //высокосный
	} else {
		fmt.Println("NO") //не высокосный
	}
}

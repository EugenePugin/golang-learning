package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// "fmt"

func work(x int) int {
	return x * x
}

func main3() {
	groupCity := map[int][]string{
		10:   {"Ивановка", "Грязь"},      // города с населением 10-99 тыс. человек
		100:  {"Королев", "Электроугли"}, // города с населением 100-999 тыс. человек
		1000: {"Москва"},                 // города с населением 1000 тыс. человек и более
	}
	cityPopulation := map[string]int{
		"Королев": 300000,
		"Москва":  20000000,
	}

	// fmt.Println("Before", cityPopulation)

	for key := range cityPopulation {
		// fmt.Println(key, value)
		// if contains(groupCity[10], key) || contains(groupCity[1000], key) {
		for _, v := range groupCity[10] {
			if key == v {
				// fmt.Print(key, "duplicate detected!")
				delete(cityPopulation, key)
			} else {
				// fmt.Println(key, "is unique")
			}
			for _, v1 := range groupCity[1000] {
				if key == v1 {
					// fmt.Print(key, "duplicate detected!")
					delete(cityPopulation, key)
				} else {
					// fmt.Println(key, "is unique")
				}
			}
		}

	}

	// fmt.Println("After", cityPopulation)
}

func main2() {
	var arrayInt [10]int
	for i := range arrayInt {
		fmt.Scan(&arrayInt[i])
	}

	myMap := make(map[int]int)
	for i := range len(arrayInt) {
		if value, ok := myMap[arrayInt[i]]; ok {
			// fmt.Printf("Using cached %d", value)
			myMap[arrayInt[i]] = value
		} else {
			// fmt.Println("No cached version is avaialble")
			myMap[arrayInt[i]] = work(arrayInt[i])
		}
	}

	// fmt.Println(arrayInt)
	// fmt.Println(myMap)
	for i := range len(arrayInt) {
		fmt.Print(myMap[arrayInt[i]], " ")
	}
}

func adding(x, y string) int64 {
	var xInt, yInt int
	var xCleaned, yCleaned string

	// fmt.Println(x, y)

	var resultx strings.Builder
	for _, rx := range x {
		// Check if the rune is a digit
		if unicode.IsDigit(rx) {
			resultx.WriteRune(rx)
		}
	}
	xCleaned = resultx.String()

	var resulty strings.Builder
	for _, ry := range y {
		// Check if the rune is a digit
		if unicode.IsDigit(ry) {
			resulty.WriteRune(ry)
		}
	}
	yCleaned = resulty.String()

	// fmt.Println(xCleaned, yCleaned)
	xInt, _ = strconv.Atoi(xCleaned)
	yInt, _ = strconv.Atoi(yCleaned)
	return int64(xInt) + int64(yInt)
}

func main4() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	separator := ";"
	decimalSeparatorSymbolRussia := ","
	decimalSeparatorSymbolInternational := "."

	parts := strings.Split(s, separator)

	// fmt.Println("Raw string A", parts[0])
	purified := strings.ReplaceAll(parts[0], " ", "")
	purified = strings.ReplaceAll(purified, decimalSeparatorSymbolRussia, decimalSeparatorSymbolInternational)
	// fmt.Println("Spaces cleanup", purified)
	floatValueA, _ := strconv.ParseFloat(purified, 64)

	// fmt.Println("Raw string B", parts[1])
	purified = strings.ReplaceAll(parts[1], " ", "")
	purified = strings.ReplaceAll(purified, decimalSeparatorSymbolRussia, decimalSeparatorSymbolInternational)
	// fmt.Println("Spaces cleanup", purified)
	floatValueB, _ := strconv.ParseFloat(purified, 64)
	floatDiv := floatValueA / floatValueB
	fmt.Printf("%.4f", floatDiv)

}

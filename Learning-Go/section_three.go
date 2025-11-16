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

func main55() {

	fn := func(a uint) uint {
		s := strconv.Itoa(int(a))
		var digits []int
		var index int
		for _, runeValue := range s {
			digit := int(runeValue - '0')
			index++
			if digit == 0 || digit%2 == 1 {
				continue
			}
			// fmt.Println(digit)
			digits = append(digits, digit)
		}
		stringSlice := make([]string, len(digits))
		for i, dig := range digits {
			stringSlice[i] = strconv.Itoa(dig) // Itoa means "Integer to ASCII"
		}
		var combinedString string // := strings.Join(stringSlice, "")
		for _, s := range stringSlice {
			// This line creates a brand new string variable in memory every time
			combinedString += s
		}
		combinedStringInt, _ := strconv.ParseUint(combinedString, 10, 16)

		if combinedStringInt == 0 {
			return 100
		} else {
			return uint(combinedStringInt)
		}

	}
	var x uint
	fmt.Scan(&x)
	fmt.Println(fn(x))
}

func isEqualAny(input string, candidates []string) bool {
	for _, candidate := range candidates {
		if input == candidate {
			return true // Found a match, stop and return true
		}
	}
	return false // No matches found after checking all candidates
}

// DetermineStringType checks if the input string can be parsed as a bool, int, or float.
func DetermineStringType(input string) string {
	// 1. Check if it's a boolean
	if _, err := strconv.ParseBool(input); err == nil {
		return "bool"
	}

	// 2. Check if it's an integer
	// We use base 10 and 64-bit size.
	if _, err := strconv.ParseInt(input, 10, 64); err == nil {
		return "int"
	}

	// 3. Check if it's a float
	// We use 64-bit size. If ParseInt failed, this might still pass if it has a decimal point.
	if _, err := strconv.ParseFloat(input, 64); err == nil {
		return "float64"
	}

	// 4. If all fail, it's just a general string
	return "string"
}

func readTaskA() (interface{}, interface{}, interface{}) {
	var x, y, operation interface{}

	readerX := bufio.NewReader(os.Stdin)
	// fmt.Print("DBG: Enter a value (e.g., 10, 3.14, or hello): ")
	input, _ := readerX.ReadString('\n')
	input = strings.TrimSpace(input) // Clean up newline characters and spaces
	switch DetermineStringType(input) {
	case "float64":
		{
			// fmt.Println("Float64 detected")
			x, _ = strconv.ParseFloat(input, 64)
		}
	case "bool":
		{
			x, _ := strconv.ParseBool(input)
			fmt.Printf("value=%t: %T\n", x, x)
			return x, y, operation
		}
	case "int":
		{
			x, _ := strconv.ParseInt(input, 10, 64)
			fmt.Printf("value=%d: %T\n", x, x)
			return x, y, operation
		}
	default: // assuming string
		{
			x = input
			fmt.Printf("value=%s: %T\n", x, x)
			return x, y, operation
		}
	}

	/////////////////

	readerY := bufio.NewReader(os.Stdin)
	// fmt.Print("DBG: Enter a value (e.g., 10, 3.14, or hello): ")
	input, _ = readerY.ReadString('\n')
	input = strings.TrimSpace(input) // Clean up newline characters and spaces
	switch DetermineStringType(input) {
	case "float64":
		{
			// fmt.Println("Float64 detected")
			y, _ = strconv.ParseFloat(input, 64)
		}
	case "bool":
		{
			y, _ := strconv.ParseBool(input)
			fmt.Printf("value=%t: %T\n", y, y)
			return x, y, operation
		}
	case "int":
		{
			y, _ := strconv.ParseInt(input, 10, 64)
			fmt.Printf("value=%d: %T\n", y, y)
			return x, y, operation
		}
	default: // assuming string
		{
			y = input
			fmt.Printf("value=%s: %T\n", input, input)
			return x, y, operation
		}
	}

	////////////

	readerOperation := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter a value of operation: ")
	input, _ = readerOperation.ReadString('\n')
	input = strings.TrimSpace(input) // Clean up newline characters and spaces
	// Сheck for symbols of operations
	operationSigns := []string{"+", "-", "*", "/"}
	if isEqualAny(input, operationSigns) {
		// fmt.Printf("DBG'%s' is one of the valid options.\n", input)
		operation = input
	} else {
		fmt.Printf("неизвестная операция\n")
		return x, y, operation
	}

	return x, y, operation
}
func readTask() (interface{}, interface{}, interface{}) {
	return 3.14, 4.13, false
}

func main66() {
	value1, value2, operation := readTaskA() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса

	// fmt.Println("The End:", value1, value2, operation)

	switch value1.(type) { // The '.(type)' syntax initiates a type switch
	case float64:
		{
			// fmt.Println("Float64 detected")
		}
	default: // assuming string
		{
			fmt.Printf("value=%t: %T\n", value1, value1)
			return
		}
	}

	switch value2.(type) { // The '.(type)' syntax initiates a type switch
	case float64:
		{
			// fmt.Println("Float64 detected")
		}
	default: // assuming string
		{
			fmt.Printf("value=%t: %T\n", value2, value2)
			return
		}
	}

	switch operation.(type) { // The '.(type)' syntax initiates a type switch
	case string:
		{
			// fmt.Println("Float64 detected")
		}
	default: // assuming string
		{
			fmt.Printf("неизвестная операция")
			return
		}
	}

	switch operation {
	case "+":
		{
			fmt.Printf("%.4f", value1.(float64)+value2.(float64))
		}
	case "-":
		{
			fmt.Printf("%.4f", value1.(float64)-value2.(float64))
		}
	case "*":
		{
			fmt.Printf("%.4f", value1.(float64)*value2.(float64))
		}
	case "/":
		{
			fmt.Printf("%.4f", value1.(float64)/value2.(float64))
		}
	default:
		{
			fmt.Printf("неизвестная операция") // unsupported value
		}
	}
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

func ProcessInputString() {
	// 1. Create a Reader that wraps standard input (os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	// 2. Read the input up to the newline character ('\n')
	// ReadString blocks until the delimiter is found.
	input, _ := reader.ReadString('\n')

	// Trim whitespace (including newline) before processing
	input = strings.TrimSpace(input)

	runeSlice := []rune(input)
	// fmt.Println("DBG2:", input)
	// fmt.Println("DBG:", string(runeSlice[utf8.RuneCountInString(input)-1]))

	if string(runeSlice[utf8.RuneCountInString(input)-1]) == "." && unicode.IsUpper(runeSlice[0]) {
		fmt.Printf("Right")
	} else {
		fmt.Printf("Wrong")
	}
}

// ReverseString takes a string and returns its character-wise reversal.
func ReverseString(s string) string {
	// 1. Convert string to a slice of runes
	runes := []rune(s)

	// 2. Loop from both ends towards the middle, swapping elements
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Perform the swap
	}

	// 3. Convert the reversed rune slice back to a string
	return string(runes)
}
func ifPalindrome() bool {
	// 1. Create a Reader that wraps standard input (os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	// 2. Read the input up to the newline character ('\n')
	// ReadString blocks until the delimiter is found.
	input, _ := reader.ReadString('\n')
	// Trim whitespace (including newline) before processing
	input = strings.TrimSpace(input)

	input2 := ReverseString(input)
	// fmt.Println("DBG2:", input)
	// fmt.Println("DBG2:", input2)

	// fmt.Println("DBG:", string(runeSlice[utf8.RuneCountInString(input)-1]))

	if input == input2 {
		return true
	} else {
		return false
	}
}

func toFindIndex() {
	var X, S string
	fmt.Scan(&X)
	// reader := bufio.NewReader(os.Stdin)
	// X, _ := reader.ReadString('\n')
	X = strings.TrimSpace(X)
	// fmt.Println("DBG1:", X)

	// reader2 := bufio.NewReader(os.Stdin)
	// S, _ := reader2.ReadString('\n')
	fmt.Scan(&S)
	S = strings.TrimSpace(S)
	// fmt.Println("DBG2:", S)

	index := strings.Index(X, S)

	if index != -1 {
		// fmt.Printf("'%s' starts at byte index %d.\n", S, index)
		fmt.Println(index)
	} else {
		// fmt.Printf("'%s' was not found.\n", S)
		fmt.Println(-1)

	}

}

func RemoveOddSymbols(s string) string {
	var builder strings.Builder
	runeSlice := []rune(s)

	// Iterate over each rune (character) in the input string
	for i := 0; i < len(s); i++ {
		// fmt.Println("DBGodd:", i, s[i], runeSlice[i])

		if i%2 == 0 {
			builder.WriteRune(runeSlice[i])
		}
	}
	// Return the new string built from valid characters
	return builder.String()
}

func ToLeaveOddOnly() {
	var X, S string
	fmt.Scan(&X)
	// reader := bufio.NewReader(os.Stdin)
	// X, _ := reader.ReadString('\n')
	X = strings.TrimSpace(X)
	// fmt.Println("DBG1:", X)

	S = RemoveOddSymbols(X)
	fmt.Println(S)

}

func ToRemoveSymbolsMoreThenOnce() {
	var X string
	fmt.Scan(&X)
	X = strings.TrimSpace(X)
	runeX := []rune(X)
	var runeY []rune
	// fmt.Println("DBG1:", runeX)
	// fmt.Println("DBG1:", runeX[0])

	for i := range runeX {
		// fmt.Println(strings.Count(X, string(runeX[i])))
		if strings.Count(X, string(runeX[i])) == 1 {
			runeY = append(runeY, runeX[i])
		}
	}
	resultingString := string(runeY)
	fmt.Println(resultingString)

}

func divide__(a int, b int) (int, error) {
	var err error
	if b == 0 {
		err = errors.New("на ноль делить нельзя")
		return 0, err
	} else {
		return a / b, nil
	}
}

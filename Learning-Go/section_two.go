package main

import (
	"bufio"
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

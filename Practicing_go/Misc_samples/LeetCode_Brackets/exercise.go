// https://leetcode.com/problems/valid-parentheses/description/
package LC_brackets

import (
	"slices"
	"strings"
)

// type Stack struct {
// 	items []string
// }

// func NewStack(s *Stack) {
// 	s.items = make([]string, 0)
// }

// func Push(s *Stack, item string) {
// 	s.items = append(s.items, item)
// }

// func Pop(s *Stack) (string, bool) {
// 	if len(s.items) == 0 {
// 		return "", false
// 	}
// 	last_index := len(s.items) - 1
// 	itemToReturn := s.items[last_index]
// 	s.items = s.items[:last_index]
// 	return itemToReturn, true //return last items
// }

// func validate_stack(symbols []string) bool {
// 	var s Stack
// 	NewStack(&s)
// 	for i := range symbols {
// 		Push(&s, symbols[i])
// 	}
// 	fmt.Println(s.items)
// 	fmt.Println("and now")
// 	// vat nextSymbolExpected []string
// 	for range symbols {
// 		symbol, ok := Pop(&s)
// 		if ok != false {
// 			fmt.Println(symbol)
// 		} else {
// 			fmt.Println("incorrect method usage")
// 		}

// 	}
// 	return true //tbd

// }

func validation_order_of_brackets(s string) bool {
	// fmt.Println(s)
	pairToDetect := [3]string{"()", "{}", "[]"}
	// let's exclude all valid pairs of brackets
	for i := 0; i < 3; i++ {
		// fmt.Println(i, "checking for", pairToDetect[i])
		for {
			if true == strings.Contains(s, pairToDetect[i]) {
				s = strings.ReplaceAll(s, pairToDetect[i], "")
				i = 0 // start over in this case to capture brackets out of the nested ones
				// fmt.Println(s)
			} else {
				break //tbd
			}
		}
	}
	// fmt.Println(s)
	if s == "" {
		return true //all brackets are in the right order
	}
	return false
}
func validateBracketType(OpenBracket string, ClosedBracket string, symbols []string) bool {
	var openBracketind, closedBracketind,
		openBracketcnt, closedBracketcnt int
	for {
		if openBracketind = slices.Index(symbols[openBracketind:], OpenBracket); openBracketind == -1 {
			// fmt.Println("no more open bracket found")
			// break
		} else {
			openBracketcnt++
		}
		if closedBracketind = slices.Index(symbols[closedBracketind:], ClosedBracket); closedBracketind == -1 {
			// fmt.Println("no closed bracket found ")
			// break
		} else {
			closedBracketcnt++
		}
		// fmt.Println(openBracketind, openBracketcnt, closedBracketind, closedBracketcnt)
		if openBracketind == -1 && closedBracketind != -1 {
			// fmt.Println("closed bracket exist, but no open bracket found")
			return false
		}
		if closedBracketind <= openBracketind {
			// fmt.Println("closed bracket must be after the open one")
			return false
		}
		// if openBracketcnt != closedBracketcnt {
		// 	fmt.Println("number of open and closed brackets should be the same")
		// 	return false
		// }
		openBracketind++
		closedBracketind++
		if openBracketind == len(symbols) {
			break
		}
		if closedBracketind == len(symbols) {
			break
		}
	}
	return true
}

func isValid(s string) bool {
	symbols := make([]string, len(s))
	symbols = strings.Split(s, "")
	// fmt.Println(symbols)

	// input validation - false, in case any other symbol is detected
	validSymbols := make([]string, 6)
	validSymbols = []string{"(", ")", "{", "}", "[", "]"}
	var validSymbolsCounter int
	if len(symbols) == 0 {
		return false
	}
	for i := range symbols {
		validSymbolsCounter = 0
		// fmt.Println("Assessing", symbols[i])
		for j := range validSymbols {
			// fmt.Println("comparing with", validSymbols[j])
			if symbols[i] == validSymbols[j] {
				validSymbolsCounter++
			}
		}
		if validSymbolsCounter == 0 {
			return false //invalid symbol discovered
		}
	}

	// Open brackets must be closed by the same type of brackets.
	// Open brackets must be closed in the correct order.
	// Every close bracket has a corresponding open bracket of the same type.
	// var ind int
	// var openBracketind, closeBracketind int
	probeType0 := slices.Index(symbols, "(")
	probeType1 := slices.Index(symbols, "{")
	probeType2 := slices.Index(symbols, "[")
	// fmt.Println("probes:", probeType0, probeType1, probeType2)

	var result bool
	if probeType0 != -1 {
		result = validateBracketType("(", ")", symbols)
		// fmt.Println("Validation for ():", result)
	}
	if probeType1 != -1 {
		result = result || validateBracketType("{", "}", symbols)
		// fmt.Println("Validation for {}:", result)
	}
	if probeType2 != -1 {
		result = result || validateBracketType("[", "]", symbols)
		// fmt.Println("Validation for []:", result)
	}

	if result == false {
		return false
	}

	return validation_order_of_brackets(s)
}

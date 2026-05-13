// https://leetcode.com/problems/decode-ways/description/?envType=problem-list-v2&envId=dh5241mj

package LC_decodeways

import (
	"slices"
)



var validCodes []string
func isValidSymbolCode(symbol string) bool {
	// fmt.Println(validCodes)
	// fmt.Println("validation of:", symbol)
	return slices.Contains(validCodes, symbol)
}

func initValidCodesArrary() {

	validCodes = make([]string, 0)
	validCodes = append(validCodes, "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"10", "11", "12", "13", "14", "15", "16", "17", "18", "19",
		"20", "21", "22", "23", "24", "25", "26")
}

func numDecodings(s string) int {
	// fmt.Println(s)
	if len(s) < 1 || len(s) > 100 {
		return 0
	}
	initValidCodesArrary()

	for i := range s {
		if string(s[i]) == "0" {
			if i > 0 {
				// fmt.Println("a")
				if !slices.Contains(validCodes[9:], string(s[i-1])+string(s[i])) {
					// fmt.Println("b")
					return 0
				}
			} else {
				// fmt.Println("c")
				return 0 // no leading 0s allowed
			}
		} else if !slices.Contains(validCodes[:9], string(s[i])) {
			// fmt.Println("d")
			return 0
		}
	}

	// special case
	if len(s) == 1 {
		return 1
	}
	var result int
	numDecodingSlice := make([]int, len(s)+1)
	numDecodingSlice[0] = 1
	numDecodingSlice[1] = 1
	// check from 1st symbol and iterate through the whole string

	for i := 2; i < len(s)+1; i++ {

		// fmt.Println("i:", i,"result:",numDecodingSlice)
		oneSymbol := string(s[i-1])
		twoSymbols := string(s[i-2]) + string(s[i-1])
		// fmt.Println(oneSymbol, twoSymbols)
		if isValidSymbolCode(oneSymbol) {
			numDecodingSlice[i] += numDecodingSlice[i-1]
			// fmt.Println("=>", numDecodingSlice[i])
		}
		if isValidSymbolCode(twoSymbols) {
			numDecodingSlice[i] += numDecodingSlice[i-2]
			// fmt.Println("=>", numDecodingSlice[i])
		}
	}

	// check each symbol and each pair of symbols of s
	// if discovered as valid count on that

	// var counter int
	// result := 1 //	as all symbols are validated at this point, at least 1symbol decoding is in place
	result = numDecodingSlice[len(numDecodingSlice)-1]
	// fmt.Println(result)

	return result // tbd
}

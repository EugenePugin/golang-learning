// https://leetcode.com/problems/decode-ways/description/?envType=problem-list-v2&envId=dh5241mj

package main

func myAtoi(s string) int {
	var result, tmp int
	isNegative := false
	isLeading := true
	var symbol string
	for i := range len(s) {
		symbol = string(s[i])
		switch symbol {
		case " ":
			if !isLeading {
				goto outFromTheLoop //quit in case of not leading
			}
			continue
		case "+":
			if !isLeading {
				goto outFromTheLoop //quit in case of not leading
			}
			isLeading = false
			continue
		case "-":
			if !isLeading {
				goto outFromTheLoop //quite in case of not leading
			}
			isLeading = false
			isNegative = true
			continue
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			isLeading = false
			tmp = int(s[i]) - 48 //используем кодировку UTF-8
		default: //non digit symbol detected
			goto outFromTheLoop
		}
		// math.MaxInt32
		if !isNegative && (1<<31-1)-10*result < tmp {
			result = 1<<31 - 1
		} else if isNegative && 1<<31-10*result < tmp {
			result = 1 << 31
		} else {
			result = 10*result + tmp
		}
	}
outFromTheLoop:
	if isNegative {
		result *= -1
	}
	return result
}

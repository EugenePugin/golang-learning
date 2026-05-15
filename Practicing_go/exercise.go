// https://leetcode.com/problems/decode-ways/description/?envType=problem-list-v2&envId=dh5241mj

package main

import "math"

func myAtoi(s string) int {
	var result int64
	var digit int64
	isNegative := false
	isLeading := true
	for i := range s {
		switch {
		case s[i] == ' ' && isLeading:
			continue
		case (s[i] == '+' || s[i] == '-') && isLeading:
			isLeading = false
			isNegative = (s[i] == '-')
			continue
		case s[i] >= '0' && s[i] <= '9':
			isLeading = false
			digit = int64(s[i]) - '0' //используем кодировку UTF-8
			if isNegative {
				if result > (-int64(math.MinInt32)-digit)/10 {
					return math.MinInt32
				}
			} else {
				if result > (int64(math.MaxInt32)-digit)/10 {
					return math.MaxInt32
				}
			}

			result = 10*result + digit

		default: //non digit symbol either non-leading sign detected
			goto done
		}
	} // loop cycle
done:
	if isNegative {
		return int(-result)
	}
	return int(result)
}

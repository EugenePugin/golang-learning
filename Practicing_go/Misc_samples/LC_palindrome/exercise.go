// https://leetcode.com/problems/palindrome-number/

package LCpalindrome

func isPalindrome(x int) bool {
	// negative numbers can't be palindromes
	// 1-digit numbers are always palindromes
	// fmt.Println(x)
	switch {
	case x < 0:
		return false
	case x <= 9 && x >= 0:
		return true
	}

	// split number by digits
	sliceDigits := make([]int, 0)
	for x > 0 {
		digit := x % 10 // get digit from the right
		// fmt.Println(digit)
		sliceDigits = append(sliceDigits, digit)
		x = x / 10 // divide by 10
		// fmt.Println(x)
	}
	// fmt.Println(sliceDigits)

	var formLeftToRight, formRightToLeft int
	multiplyer := 1
	for i := 0; i < len(sliceDigits); i++ {
		formLeftToRight += sliceDigits[i] * multiplyer
		formRightToLeft += sliceDigits[len(sliceDigits)-i-1] * multiplyer
		multiplyer *= 10
	}
	if formLeftToRight == formRightToLeft {
		return true
	}
	return false
}

// https://leetcode.com/problems/plus-one/description/
package LCplusone

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digitToInc := digits[i]
		if digitToInc != 9 {
			digits[i] = digitToInc + 1
			return digits
		} else { // if 9, let's check elder digits
			digitToInc = 0
			digits[i] = digitToInc
			if i == 0 { //старшая цифра
				newSlice := make([]int, len(digits)+1)
				newSlice[0] = 1
				copy(newSlice[1:], digits)
				// fmt.Println("newSlice:", newSlice)
				return newSlice
			}
		}
		// fmt.Println("moving1:", moving1)
		// fmt.Println("digits:", digits)
	}

	return nil //tbd
}

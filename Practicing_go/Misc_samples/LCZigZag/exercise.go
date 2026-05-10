// https://leetcode.com/problems/zigzag-conversion/description/

package LCZigZag

import "strings"

func convert(s string, numRows int) string {
	// input validation
	if numRows < 1 ||
		numRows > 1000 {
		return ""
	}
	if len(s) < 1 ||
		len(s) > 1000 {
		return ""
	}
	// special case
	if numRows == 1 {
		return s
	}
	// create 2d map of indeces
	// fmt.Println("ok, let's go - #3")
	slice2d := make([][]rune, numRows)
	for i := range numRows {
		slice2d[i] = make([]rune, len(s))
	}
	for row := range numRows {
		for col := range len(s) {
			slice2d[row][col] = 0
		}
	}
	// printSlice2d(slice2d)
	// fmt.Println()
	// go there from up to down, once achieve the limit to up and left

	var row, col int
	var zigzagMode bool
	for i := range len(s) {
		// fmt.Println("i:", i)
		slice2d[row][col] = rune(s[i])
		// fmt.Println(string(slice2d[row][col]))
		// printSlice2d(slice2d)
		if i%(numRows-1) == 0 {
			zigzagMode = true
		}
		if i%(numRows+numRows-2) == 0 {
			zigzagMode = false
		}
		// fmt.Println("\tzigzagMode:", zigzagMode)
		if zigzagMode {
			row--
			col++
		} else {
			row++
		}
		// fmt.Println("\trow:", row, "col:", col)
	}

	// compact the slice reading line by line

	var result strings.Builder
	for row := range numRows {
		for col := range slice2d[row] {
			if slice2d[row][col] != 0 {
				result.WriteString(string(slice2d[row][col]))
				// fmt.Println(result, slice2d[row][col])
			}
		}

	}

	return result.String()
}

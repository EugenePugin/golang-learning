package main

import (
	"fmt"
)

func one12a(dbgMode uint) {
	var workArray [16]uint

	if dbgMode == 1 { //dbg mode
		workArray = [16]uint{99, 151, 137, 71, 117, 187, 20, 93, 187, 67, 1, 2, 3, 5, 7, 8}
	} else {

	}

	// if R <= 0 {
	// 	formattedString := fmt.Sprintf("%4.2f", R)
	fmt.Println(workArray)
	// } else if R > 10000 {
	// 	formattedString := fmt.Sprintf("%e", R)
	// 	fmt.Println(formattedString)
	// } else {
	// 	convertedR := (float64)((int64)(R*R*10000)) / 10000
	// 	formattedString := fmt.Sprintf("%.4f", convertedR)
	// 	fmt.Println(formattedString)
	// }
}

func main() {
	one12a(1)
}

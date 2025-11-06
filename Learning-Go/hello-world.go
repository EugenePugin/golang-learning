package main

import (
	"fmt"
	"strconv"
	"strings"
)

func one13(mode int) {
	var n, m int

	if mode == 1 {
		// fmt.Println("DBG mode welcomes you!")
		n = 2343434
		m = 4
	} else {
		// fmt.Println("RELEASE mode welcomes you!")
		fmt.Scan(&n, &m)
	}
	// input data check
	if n < 1 || m < 1 {
		fmt.Printf("Input data are out of range")
		return
	}
	// demount value to digits
	strSliceRaw := strings.Split(strconv.Itoa(n), "")
	// fmt.Println("DBG: ", strSliceRaw) // Output: [1 2 3 4 5]
	var strSliceProcessed []string
	for i := 0; i < len(strSliceRaw); i++ {
		// fmt.Println("DBG: ", m)
		if strconv.Itoa(m) != strSliceRaw[i] {
			strSliceProcessed = append(strSliceProcessed, strSliceRaw[i])
		}
	}

	// for i := 0; i < len(strSliceProcessed); i++ {
	// 	fmt.Println("DBG: ", strSliceProcessed[i])
	// }

	fmt.Println(strings.Join(strSliceProcessed, ""))
	// remove the 1st occurence of the specified digit
	// 	TODO

}

func main() {
	one13(0)
}

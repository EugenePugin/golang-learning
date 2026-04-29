package main

import (
	"fmt"
)

func main() {
	digits := [...]int{9}

	fmt.Println(digits)
	/*result := */ plusOne(digits[:])
	// fmt.Println(result)
	// fmt.Println(slices.Compare(result, digits[:]))
}

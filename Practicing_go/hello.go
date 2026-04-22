package main

import (
	"fmt"
	"time"
)

func main() {
	array := [...]int{11, 2, 7, 15}
	slice := make([]int, len(array))
	slice = array[:]

	target := 26

	fmt.Println(slice, target)
	time0 := time.Now()
	result := twoSum(slice, target)
	time1 := time.Since(time0) * time.Microsecond
	fmt.Println(result, time1)
}

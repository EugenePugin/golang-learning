// practice for generics

// given a slice of int64 and the same for float64
// implement a function to sort them
// using generics

package generics

import (
	"fmt"
	"math/rand"
)

func sortInt(sliceInt []int64) error {
	for i := 0; i < len(sliceInt)-1; i++ {
		for j := 0; j < len(sliceInt)-1-i; j++ {
			if sliceInt[j] > sliceInt[j+1] {
				tmp := sliceInt[j+1]
				sliceInt[j+1] = sliceInt[j]
				sliceInt[j] = tmp
			}
		}
	}
	return nil
}

type Number interface {
	int64 | float64
}

func sort[T Number](slice []T) error {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				tmp := slice[j+1]
				slice[j+1] = slice[j]
				slice[j] = tmp
			}
		}
	}
	return nil
}

func main() {
	fmt.Println("Hey!")
	var sliceInt []int64
	sliceInt = make([]int64, 5)
	for i := range sliceInt {
		sliceInt[i] = rand.Int63n(10)
	}
	fmt.Printf("%d", sliceInt)
	fmt.Println()
	if ok := sortInt(sliceInt); ok == nil {
		fmt.Println(sliceInt)
	}
	sliceFloat := make([]float64, 5)
	for i := range sliceFloat {
		sliceFloat[i] = rand.Float64()
	}
	fmt.Printf("%.2f", sliceFloat)
	fmt.Println()
	if ok := sort(sliceFloat); ok == nil {
		fmt.Printf("%.2f", sliceFloat)
	}

	var sliceInt32 []int32
	sliceInt32 = make([]int32, 5)
	for i := range sliceInt {
		sliceInt32[i] = rand.Int31n(10)
	}
	fmt.Printf("%d", sliceInt32)
	fmt.Println()
	if ok := sortInt(sliceInt32); ok == nil { //compile issues
		fmt.Println(sliceInt32)
	}

}

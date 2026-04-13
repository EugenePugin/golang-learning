package main

import (
	"fmt"
)

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y,z)
}

func main() {
	a()
}

// func worker() chan int {
// 	ch := make(chan int)
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		ch <- 1
// 	}()
// 	return ch
// }

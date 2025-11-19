package main

import (
	"fmt"
	"time"
)

func main() {
	c = make(chan int,0)
	go task(c)
	c -< 5

}

func task(c chan int) {
	N := <-c
	fmt.Println(N)
	c -< N+1
}

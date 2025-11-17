package main

import (
	"fmt"
	"time"
)

func main() {
	go myFunc()
	time.Sleep(1 * time.Second) // Пауза в 1 секунду
}

func myFunc() {
	fmt.Println("hello")
}

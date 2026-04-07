package hello_ex

import (
	"fmt"
	"strconv"
)

func main_a() {
	fmt.Println("main() started")
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream) // start goroutine
	// go outputStreamPrinter(outputStream)

	i := 1
	inputStream <- strconv.Itoa(i)

	i = 2
	inputStream <- strconv.Itoa(i)

	i = 3
	inputStream <- strconv.Itoa(i)

	i = 3
	inputStream <- strconv.Itoa(i)

	i = 4
	inputStream <- strconv.Itoa(i)

	// close(inputStream)
	// close(outputStream) // close channel

	// fmt.Println("main() stopped")
}

// func outputStreamPrinter(outputStream chan string) {
// 	defer close(outputStream)
// 	for {
// 		val, ok := <-outputStream
// 		if !ok {
// 			fmt.Println(val, ok, "outputStream empty!")
// 			break // exit break loop
// 		} else {
// 			fmt.Println("Read from outputStream: ", val, ok)
// 		}
// 	}
// 	// close(outputStream) // close channel
// }

var previousValueInInputStream string

func removeDuplicates(inputStream chan string, outputStream chan string) {

	for v := range inputStream {
		if v != previousValueInInputStream {
			outputStream <- v
		}
		previousValueInInputStream = v
	}
	close(outputStream)
	// // for {
	// // 	val, ok := <-inputStream
	// // 	if !ok {
	// // 		fmt.Println(val, ok, "<-- loop broke!")
	// // 		break // exit break loop
	// // 	} else {
	// // 		fmt.Println("Read from inputStream: ", val, ok)
	// // 		if val != previousValueInInputStream {
	// // 			outputStream <- val
	// // 			previousValueInInputStream = val
	// // 		}
	// // 	}
	// }
}

func main_b() {
	// <-myFunc()
	done := make(chan struct{})

	go func(d chan struct{}) {
		work_b()
		// return d
		close(d)
	}(done)

	<-done
}

func work_b() {
	fmt.Println("hello work")
}

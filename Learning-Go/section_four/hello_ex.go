package main1

import (
	"fmt"
	"sync"
	"time"
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	// A wait group to ensure all calculations are complete before potentially closing 'out'
	var wg sync.WaitGroup
	wg.Add(n) // We will launch N concurrent operations
	// for value := range bufferedChan {
	// 	fmt.Printf("Received: %d\n", value)
	// }
	for i := 0; i < n; i++ {
		// Launch a new goroutine for each pair to calculate fn concurrently
		go func() {
			defer wg.Done()
			x1 := <-in1 // Read sequentially from input channels
			x2 := <-in2
			// fmt.Println("DBG:", i, x1, x2)
			// Calculate both fn(val1) and fn(val2) concurrently if possible within fn itself,
			// or simply run the operations in parallel across different cores.
			result1 := fn(x1)
			result2 := fn(x2)
			// Send the final result to the output channel
			// NOTE: If 'out' is unbuffered and there are many goroutines, this might block
			// until a receiver is ready.
			out <- result1 + result2
		}()
	}

	// Optional: Launch a separate goroutine to wait for all tasks and close the out channel
	// if the caller expects the channel to be closed.
	go func() {
		wg.Wait()
		// close(out) // Only close if the caller specifically requires it.
	}()
}

func slowFunction(x int) int {
	fmt.Println("slowFunction() started")
	time.Sleep(time.Second * 1)
	return x * x
}

func main() {
	// fmt.Println("Main() started")
	const N = 3

	// Use a buffered channel for arguments since we send multiple values quickly
	arguments1 := make(chan int, N)
	arguments2 := make(chan int, N)
	// We will close arguments manually after all sends are complete

	// Initialize the result channel before using it
	result := make(chan int, N)
	// We will close it manually

	for i := 0; i < N; i++ {
		// Send 5 integers to the arguments channel
		arguments1 <- i
		arguments2 <- N - i
		fmt.Println("Arguments: ", i, N-i)
	}

	// Close the arguments channel to signal that no more data is coming
	close(arguments1)
	close(arguments2)

	// Since we closed arguments, the calculator function will finish naturally
	// using the `!ok` check in the select statement. We don't necessarily need
	// to use the 'done' channel in this specific flow.

	// Reworking the main logic to wait for the final accumulated result:
	// We pass the channel of arguments and the 'done' channel (though it won't be used
	// in this revised flow, it fulfills the function signature) to a *single* instance
	// of the calculator function that runs to completion.

	// Call calculator once and read the final result
	merge2Channels(slowFunction, arguments1, arguments2, result, N)
	for i := 0; i < N; i++ {
		res := <-result
		fmt.Printf("DBG: results: %d\n", res)
	}

	// The deferred close(done) from the original code is no longer necessary
	// as we never sent anything to it in this fixed logic.
}

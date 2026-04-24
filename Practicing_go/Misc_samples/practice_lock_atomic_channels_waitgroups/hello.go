// implement counting from 0 till 100
// accelerate with goroutines
// a) using channels and b) wait groups
// a) using lock and b) using atomics
// implemenet timeout

package plac //main

import (
	"context"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Counter struct {
	counter int64
	// mu      sync.Mutex
}

func NewCounter(counterStructPtr *Counter, counterValue int64) {
	counterStructPtr.counter = counterValue

}

// func CounterInc(counterStructPtr *Counter, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	mu.Lock()
// 	defer mu.Unlock()
// 	counterStructPtr.counter++

// 	// atomic.AddInt64(&counterStructPtr.counter,1)

// 	// time.Sleep(100 * time.Millisecond)
// 	// fmt.Println("counter value:", counterStructPtr.counter)
// }

// func CounterInc(counterStructPtr *Counter, ch chan bool) {
// 	atomic.AddInt64(&counterStructPtr.counter, 1)
// 	ch <- true
// }

func CounterInc(ctx context.Context, counterStructPtr *Counter, ch chan bool) {
	timeToSleep := time.Duration(rand.Intn(10))*time.Millisecond*10 + 950*time.Millisecond
	// fmt.Println("timeToSleep:", timeToSleep)
	// time.Sleep(timeToSleep)

	select {
	case <-ctx.Done():
		{
			// fmt.Println(ctx.Err())
			ch <- false
		}
	case <-time.After(timeToSleep):
		{
			// fmt.Println("all is good")
			atomic.AddInt64(&counterStructPtr.counter, 1)
			ch <- true
		}
	}
}

func GetCounter(counterStructPtr *Counter) int64 {
	return counterStructPtr.counter
}

func main() {
	// fmt.Println("Hey")

	var counter Counter
	N := 10
	time0 := time.Now()
	// option1: nearly 1000ms
	// for i := 0; i < N; i++ {
	// 	CounterInc(&counter)
	// }
	// var wg sync.WaitGroup
	// wg.Add(N)
	// for range N {
	// 	go CounterInc(&counter, &wg)
	// }
	// wg.Wait()

	NewCounter(&counter, 0)
	ch := make(chan bool)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for range N {
		go CounterInc(ctx, &counter, ch)
	}
	for range N {
		if <-ch == true {
			fmt.Println("thumbs up")
		} else {
			fmt.Println("timeout")
		}
	}

	time1 := time.Since(time0).Milliseconds()
	fmt.Println(GetCounter(&counter), " - should be equal to", N)
	fmt.Println("Execution time:", time1, "ms")
}

package more_goroutines

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// let's have 2 functions:
// 1st generate random number with 500ms delay, 2nd generate some hash-function from it with the similar delay
// 3rd build a map with the 1st number is a key, and the 2nd as a value
const randomNumberMax = 100
const timeDelayMultiplyer = 1
const mapSize = 10000000

var myMap map[int]int
var wg sync.WaitGroup
var mutex sync.Mutex
var mutexRW sync.RWMutex

func numberGenerator() int {
	result := rand.Intn(randomNumberMax)
	time.Sleep(timeDelayMultiplyer * 100 * time.Millisecond)
	// fmt.Println("function 1st: ", result)
	return result
}

func hashFunction(x int) int {
	result := int(math.Pow(float64(x), 2))
	time.Sleep(timeDelayMultiplyer * 500 * time.Millisecond)
	// fmt.Println(result)
	return result
}

func mapBuilder1(key, value int) {
	if _, ok := myMap[key]; ok {
		fmt.Println("the key", key, "is already there, skipping")
		return
	}
	myMap[key] = value
}

func goroutines_practice1() { //6 seconds
	fmt.Println("Hey, I am goroutines_practice 1!")

	myMap = make(map[int]int)

	time0 := time.Now()
	for range mapSize {
		x := numberGenerator()
		hashX := hashFunction(x)
		mapBuilder1(x, hashX)
	}
	duration := time.Since(time0)

	fmt.Println(myMap)
	fmt.Println("It took", int(duration.Seconds()), "seconds")
	// time.Sleep(3 * time.Second)

}

func mapBuilder2() {
	defer wg.Done()
	key := numberGenerator()
	mutexRW.RLock()
	if _, ok := myMap[key]; ok {
		mutexRW.RUnlock()
		// fmt.Println("the key", key, "is already there, skipping")
		return
	}
	mutexRW.RUnlock()

	value := hashFunction(key)
	mutexRW.Lock()
	if _, ok := myMap[key]; ok {
		mutexRW.Unlock()
		// fmt.Println("the key", key, "was just used by someone in parallel, skipping")
		return
	}
	myMap[key] = value
	mutexRW.Unlock()
}

func goroutines_practice2() { //0 seconds
	// fmt.Println("Hey, I am goroutines_practice 2!")

	myMap = make(map[int]int)

	time0 := time.Now()
	for range mapSize {
		wg.Add(1)
		go mapBuilder2()
	}
	wg.Wait()
	duration := time.Since(time0)

	// fmt.Println(myMap)
	fmt.Println("It took", int(duration.Seconds()), "seconds")
	// time.Sleep(3 * time.Second)

}

func numberGeneratorB(ch chan<- int) {
	result := rand.Intn(randomNumberMax)
	time.Sleep(timeDelayMultiplyer * 100 * time.Millisecond)
	// fmt.Println("function 1st: ", result)
	ch <- result
}

func hashFunctionB(chIn <-chan int, chOut chan<- int) {
	x := <-chIn
	result := int(math.Pow(float64(x), 2))
	time.Sleep(timeDelayMultiplyer * 500 * time.Millisecond)
	chOut <- result
	// fmt.Println(result)
	// return result
}

func mapBuilderB() {
	var chAB, chBC chan int
	defer wg.Done()

	chAB = make(chan int)
	chBC = make(chan int)

	defer close(chAB)
	defer close(chBC)

	go numberGeneratorB(chAB)
	key := <-chAB
	mutexRW.RLock()
	if _, ok := myMap[key]; ok {
		mutexRW.RUnlock()
		// fmt.Println("the key", key, "is already there, skipping")
		return
	}
	mutexRW.RUnlock()

	go hashFunctionB(chAB, chBC)
	chAB <- key
	value := <-chBC

	mutexRW.Lock()
	if _, ok := myMap[key]; ok {
		mutexRW.Unlock()
		// fmt.Println("Woops, someone just created the map entry with this key, skipping")
		return
	}
	myMap[key] = value
	mutexRW.Unlock()

}

func goroutinesPracticeB() {
	// implement goroutines sequence via channels
	// including closure on timeout
	fmt.Println("Hey, I am goroutines_practice B!")

	myMap = make(map[int]int)

	time0 := time.Now()
	for range mapSize {
		wg.Add(1)
		go mapBuilderB()
	}
	wg.Wait()
	duration := time.Since(time0)

	// fmt.Println(myMap)
	fmt.Println("It took", int(duration.Seconds()), "seconds")
	// time.Sleep(3 * time.Second)
}

func longlongFunc(ctx context.Context, pause int) {
	defer wg.Done()
	fmt.Println("longlongfunc")
	// time.Sleep(time.Duration(pause) * time.Second)
	timer := time.NewTimer(time.Duration(pause) * time.Second)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		fmt.Println("longlongFunc ended with ctx.Done: ", ctx.Err())
		return
	case <-timer.C:
		fmt.Println("longlongFunc completed by itself")
		return
	}
}

func goroutinesPractice3() {
	fmt.Println("Hey! #3")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	wg.Add(1)
	go longlongFunc(ctx, 5)
	wg.Wait()
	fmt.Println("The End")

}

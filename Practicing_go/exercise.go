package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
)

const mapSize = 10

type MyMap struct {
	myMap map[int]int
	mu    sync.Mutex
}

func NewMyMap() *MyMap {
	var obj MyMap
	obj.myMap = make(map[int]int, mapSize)
	return &obj
}

func goroutine12(ctx context.Context, obj *MyMap, mul int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hey, I am goroutine with mul", mul)
	var key int
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		key = rand.IntN(mapSize)

		obj.mu.Lock()
		if len(obj.myMap) >= mapSize {
			obj.mu.Unlock()
			fmt.Println("the map is already filled")
			return
		}
		if _, ok := obj.myMap[key]; !ok {
			obj.myMap[key] = rand.IntN(mul * mapSize)
			obj.mu.Unlock()
			break
		}
		obj.mu.Unlock()
	}
}

// func main() {
// 	sampleMap := NewMyMap()
// 	// fmt.Println(sampleMap.myMap)

// 	ctx := context.Background()

// 	var wg sync.WaitGroup

// 	for i := range mapSize / 2 {
// 		fmt.Println("iteration", i)
// 		wg.Add(1)
// 		go goroutine12(ctx, sampleMap, 2, &wg)
// 		wg.Add(1)
// 		go goroutine12(ctx, sampleMap, 4, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("mission complete: ", sampleMap.myMap)
// }

func practiceWithMaps() {
	sampleMap := NewMyMap()
	// fmt.Println(sampleMap.myMap)

	ctx := context.Background()

	var wg sync.WaitGroup

	for i := range mapSize / 2 {
		fmt.Println("iteration", i)
		wg.Add(1)
		go goroutine12(ctx, sampleMap, 2, &wg)
		wg.Add(1)
		go goroutine12(ctx, sampleMap, 4, &wg)
	}
	wg.Wait()
	fmt.Println("mission complete: ", sampleMap.myMap)
}

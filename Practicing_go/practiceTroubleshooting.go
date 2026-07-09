package main

import (
	"context"
	"fmt"
	"math/rand"
	_ "net/http/pprof" // Регистрирует обработчики на /debug/pprof
	"sync"
	"time"
)

type MapStruct struct {
	myMap    map[int]int
	mapSize  int
	mapMutex sync.Mutex
}

// let's have two gorotines, adding random items to the same map

func generator1() int {
	return rand.Intn(10)
}

func generator2() int {
	return rand.Intn(100)
}

func adder(ctx context.Context, obj *MapStruct, key int, f func() int, wg *sync.WaitGroup) {
	// fmt.Println(mapToAdd)
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("Termination via ctx:", ctx.Err())
		return
	default:
		obj.mapMutex.Lock()
		defer obj.mapMutex.Unlock()
		if _, ok := obj.myMap[key]; ok {
			fmt.Println("the key", key, "is already there, skipping")
			return
		}
		obj.myMap[key] = f()
	}
}

func troubleshooting_practice() {

	var obj MapStruct
	obj.myMap = make(map[int]int)
	obj.mapSize = 1000
	//
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup

	for i := range obj.mapSize {
		wg.Add(1)
		if rnd := rand.Intn(2); rnd > 0 {
			go adder(ctx, &obj, i, generator1, &wg)
			continue
		}
		go adder(ctx, &obj, i, generator2, &wg)
	}
	wg.Wait()
	fmt.Println(obj.myMap)

}

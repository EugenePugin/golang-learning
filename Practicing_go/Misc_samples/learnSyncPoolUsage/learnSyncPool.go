package learnsyncpoolusage

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	_ "net/http/pprof" // ВАЖНО: анонимный импорт регистрирует роуты /debug/pprof
	"sync"
)

//1.  to implement an infinite loop, where 10 goroutines are actively consuming memory:
// getting a blob of data, filling by random numbers, returning its sum, and fill the map
//2.  confirm mem consumption with flamechart
//3.  to rework #1 with sync.pool

type learnSyncPoolSample struct {
	mu        sync.Mutex
	wg        sync.WaitGroup
	mapOfSums map[int]int
	pool      sync.Pool
}

func New(mapOfSums map[int]int, poolSize int) *learnSyncPoolSample {
	return &learnSyncPoolSample{
		mapOfSums: mapOfSums,
		pool: sync.Pool{
			// New указывает, что делать, если ящик пуст
			New: func() any {
				return make([]int, poolSize)
			},
		},
	}
}

func (s *learnSyncPoolSample) syncPoolUsageSample(ind int, usePool bool) {
	defer s.wg.Done()
	var sum int
	var sliceOfData []int
	if usePool {
		sliceOfData = s.pool.Get().([]int)
		defer s.pool.Put(sliceOfData)
	} else {
		sliceOfData = make([]int, blobSize)
	}

	for j := range blobSize {
		sliceOfData[j] = rand.IntN(blobSize)
		sum += sliceOfData[j]
	}
	// fmt.Println(sliceOfData, sum)
	s.mu.Lock()
	s.mapOfSums[ind] = sum
	s.mu.Unlock()
}

const blobSize = 1000

// var slicePool sync.Pool

func LearnSyncPoolUsage() {
	fmt.Println("hey, I am LearnSyncPoolUsage!")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	const mapSize = 10

	mapOfSums := make(map[int]int, mapSize)
	syncPoolObj := New(mapOfSums, blobSize)

	fmt.Println(mapOfSums)

	for {
		for i := range mapSize {
			syncPoolObj.wg.Add(1)
			go syncPoolObj.syncPoolUsageSample(i, false)
			// go syncPoolObj.syncPoolUsageSample(i, true)

		}
		syncPoolObj.wg.Wait()
	}
}

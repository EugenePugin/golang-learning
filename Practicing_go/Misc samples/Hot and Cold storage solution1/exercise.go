// https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/description/?envType=daily-question&envId=2026-04-15
package HCS1

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

type FullName struct {
	FirstName, LastName string
}

var hotStorage map[int]FullName
var hotStorageIdx int
var hotStorageIdxSlice []int

var coldStorage map[int]FullName

// var coldStorageIdx int
var coldStorageIdxSlice []int

const HOT_STORAGE_TYPE int = 0
const COLD_STORAGE_TYPE int = 1
const SUCCESS int = 0
const FAILURE int = 1

var wg sync.WaitGroup
var mu sync.RWMutex

func NewStorage() int {
	hotStorage = make(map[int]FullName)
	hotStorageIdxSlice = make([]int, 0)
	coldStorage = make(map[int]FullName)
	coldStorageIdxSlice = make([]int, 0)

	// fmt.Println(hotStorage, hotStorageIdxSlice)
	return SUCCESS
}

func Store(value FullName) int {
	// fmt.Println("Store...")
	// fmt.Println(hotStorage, hotStorageIdx, hotStorageIdxSlice)
	mu.Lock()
	idx := hotStorageIdx
	hotStorage[idx] = value
	hotStorageIdxSlice = append(hotStorageIdxSlice, idx)
	hotStorageIdx++
	mu.Unlock()

	wg.Add(1)
	go func(key int, value FullName) {
		// var mu sync.Mutex
		defer wg.Done()
		time.Sleep(1 * time.Second) // emulating async
		mu.Lock()
		coldStorage[key] = value
		coldStorageIdxSlice = append(coldStorageIdxSlice, key)
		mu.Unlock()
	}(idx, value)

	return SUCCESS
}

func dbgCacheInvalidationEmu() {
	var rndNum int
	var percentage float32
	// getting percentage 1...len(hotStorageIdxSlice)
	for {
		rndNum = rand.Intn(len(hotStorageIdxSlice))
		percentage = float32(rndNum) / float32(len(hotStorageIdxSlice))
		if percentage != 0 {
			break
		}
	}
	fmt.Println("Deleting the random", 100*percentage, "% of records from the hot storage")
	rndUniqueNums := make([]int, 0)
	// fmt.Println(percentage, rndUniqueNums)
	var i int
	for {
		rndNum = rand.Intn(len(hotStorageIdxSlice))
		if !slices.Contains(rndUniqueNums, rndNum) {
			rndUniqueNums = append(rndUniqueNums, rndNum)
			i++
		}
		if i > int(percentage*float32(len(hotStorage))) {
			break
		}
	}

	// fmt.Println(rndUniqueNums, len(rndUniqueNums), percentage)

	// var mu sync.Mutex
	mu.Lock()
	for i := range rndUniqueNums {
		// fmt.Println("Deleting key ", rndUniqueNums[i], "... ")
		_, exists := hotStorage[rndUniqueNums[i]]
		if exists {

			delete(hotStorage, rndUniqueNums[i])
		} else {
			// fmt.Println("Key ", rndUniqueNums[i], " is not found")
		}
	}
	mu.Unlock()

}

func Retrieve(storageType int, idx int) (FullName, int) {
	var empty_record FullName

	mu.RLock()
	defer mu.RUnlock()

	// Пытаемся взять из HOT
	if val, ok := hotStorage[idx]; ok {
		return val, SUCCESS
	}

	// Если не нашли, значит запись могла быть инвалидирована (удалена)
	// Пытаемся взять из COLD
	// fmt.Print("now getting from COLD")

	if val, ok := coldStorage[idx]; ok {
		return val, SUCCESS
	}

	return empty_record, FAILURE
}

func PrintTheStorageContent() int {
	wg.Wait()
	if 0 == len(hotStorageIdxSlice) {
		return 0
	}
	fmt.Println("Checking the storage...")
	for i := range len(hotStorageIdxSlice) {
		value, discovered := Retrieve(HOT_STORAGE_TYPE, hotStorageIdxSlice[i])
		if SUCCESS == discovered {
			fmt.Println("[", hotStorageIdxSlice[i], "]", value)
		} else {
			fmt.Println("[", hotStorageIdxSlice[i], "]", "not found")
		}
	}
	return 0
}

func StorageSize() int {
	wg.Wait()
	if 0 == len(hotStorageIdxSlice) {
		return 0
	}
	var theStorageItemsCount int
	// fmt.Println("Checking the storage items...")
	for i := range len(hotStorageIdxSlice) {
		_, discovered := Retrieve(HOT_STORAGE_TYPE, hotStorageIdxSlice[i])
		if SUCCESS == discovered {
			theStorageItemsCount++
			// fmt.Println("[", hotStorageIdxSlice[i], "]", value)
		} else {
			// fmt.Println("[", hotStorageIdxSlice[i], "]", "not found")
		}
	}
	return theStorageItemsCount
}

func dbg_PrintTheColdStorageContent() int {
	wg.Wait()
	if 0 == len(coldStorageIdxSlice) {
		return 0
	}
	for i := range len(coldStorageIdxSlice) {
		fmt.Print("Checking the cold storage...")
		value, discovered := Retrieve(COLD_STORAGE_TYPE, coldStorageIdxSlice[i])
		if SUCCESS == discovered {
			fmt.Println("[", coldStorageIdxSlice[i], "]", value)
		}
	}
	return 0
}

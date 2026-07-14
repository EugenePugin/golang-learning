package exercise

import (
	"context"
	"errors"
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
	mapSync  sync.Map
}

// MapItem связывает ключ и сгенерированное значение для передачи по каналу
type MapItem struct {
	Key   int
	Value int
}

func New(mapSize int) (*MapStruct, error) {
	var obj MapStruct
	if mapSize <= 0 {
		return nil, errors.New("Error: mapSize must be positive")
	}
	if mapSize <= 100 {
		return nil, errors.New("Error: mapSize should be above 100 to make this program relevant to its goal")
	}
	obj.myMap = make(map[int]int)
	obj.mapSize = mapSize
	return &obj, nil
}

// let's have two gorotines, adding random items to the same map

func generator1() int {
	return rand.Intn(10)
}

func generator2() int {
	return rand.Intn(100)
}

func adder0(ctx context.Context, obj *MapStruct, key int, f func() int, wg *sync.WaitGroup) {
	// fmt.Println(mapToAdd)
	defer wg.Done()
	select {
	case <-ctx.Done():
		// fmt.Println("Termination via ctx:", ctx.Err())
		return
	default:
		value := f()
		obj.mapMutex.Lock()
		defer obj.mapMutex.Unlock()
		if _, ok := obj.myMap[key]; ok {
			// fmt.Println("the key", key, "is already there, skipping")
			return
		}
		obj.myMap[key] = value
	}
}

func adder1(ctx context.Context, obj *MapStruct, key int, f func() int, wg *sync.WaitGroup) {
	// fmt.Println(mapToAdd)
	defer wg.Done()
	select {
	case <-ctx.Done():
		// fmt.Println("Termination via ctx:", ctx.Err())
		return
	default:
		obj.mapSync.LoadOrStore(key, f())
	}
}

// adder2 (Fan-In воркер) — ничего не знает про мапу и мьютексы.
// Он просто отправляет сгенерированные данные в канал-семафор.
func adder2(ctx context.Context, key int, f func() int, ch chan<- MapItem, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	default:
		// Вычисляем значение и отправляем в общий канал
		ch <- MapItem{
			Key:   key,
			Value: f(),
		}
	}
}

func troubleshooting_practice(mapSize int, mode int) error { //naive implementation
	obj, err := New(mapSize)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(mode)
	switch mode {
	case 0: // naive implementation
		var wg sync.WaitGroup
		for range obj.mapSize {
			key := rand.Intn(obj.mapSize / 10)
			wg.Add(1)
			if key%2 == 0 {
				go adder0(ctx, obj, key, generator1, &wg)
			} else {
				go adder0(ctx, obj, key, generator2, &wg)
			}
		}
		wg.Wait()
		// fmt.Println(obj.myMap)
	case 1: // using sync.Map
		var wg sync.WaitGroup
		for range obj.mapSize {
			key := rand.Intn(obj.mapSize / 10)
			wg.Add(1)
			if key%2 == 0 {
				go adder1(ctx, obj, key, generator1, &wg)
			} else {
				go adder1(ctx, obj, key, generator2, &wg)
			}
		}
		wg.Wait()
	// obj.mapSync.Range(func(key, value any) bool {
	// 	fmt.Printf("Ключ: %v, Значение: %v\n", key, value)
	// 	return true // Продолжить итерацию (false для прерывания)
	// })
	case 2: // Паттерн Fan-In (Каналы, БЕЗ блокировок мапы)
		var wg sync.WaitGroup

		// Создаем буферизованный канал для слияния результатов работы воркеров
		ch := make(chan MapItem, obj.mapSize)

		// Запускаем воркеров
		for range obj.mapSize {
			wg.Add(1)
			key := rand.Intn(obj.mapSize / 10)
			if key%2 == 0 {
				go adder2(ctx, key, generator1, ch, &wg)
			} else {
				go adder2(ctx, key, generator2, ch, &wg)
			}
		}

		// Запускаем фоновую горутину, которая закроет канал результатов
		// строго после того, как ВСЕ воркеры завершат свою работу.
		go func() {
			wg.Wait()
			close(ch)
		}()

		// Главная горутина последовательно вычитывает данные из канала.
		// Только этот поток имеет доступ на запись к obj.myMap.
		for item := range ch {
			if _, ok := obj.myMap[item.Key]; !ok {
				obj.myMap[item.Key] = item.Value
			}
		}

	default:
		return fmt.Errorf("The value of mode %d is not yet implemented", mode)
	}

	return nil
}

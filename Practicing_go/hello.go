package main

import (
	"fmt"
	_ "net/http/pprof" // Регистрирует эндпоинты pprof под капотом
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

// Имитируем тяжелую функцию для наглядности профилирования
func heavyWork() {
	for {
		var s []string
		// Бесконечно аллоцируем память и грузим процессор
		for i := 0; i < 100000; i++ {
			s = append(s, fmt.Sprintf("number: %d", i))
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// 1. Включаем сбор профилей по мьютексам и блокировкам
	runtime.SetMutexProfileFraction(1) // Логировать 100% событий конкуренции мьютексов
	runtime.SetBlockProfileRate(1)     // Логировать 100% событий блокировок горутин

	// 2. Запускаем нашу логику
	// check overall time on
	time0 := time.Now()
	if err := troubleshooting_practice(1000, 0); err != nil {
		// mode 0: 1.25ms
		// mode 1: 4+ using sync.map
		// mode 2: fan-in
		panic(err)
	}

	duration := time.Since(time0)

	fmt.Println("Done in", duration)
	// 3. Записываем накопленный профиль мьютексов в файл перед выходом
	f, err := os.Create("mutex.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if p := pprof.Lookup("mutex"); p != nil {
		_ = p.WriteTo(f, 0)
	}
	fmt.Println("Профиль мьютексов сохранен в mutex.pprof")

	runtime.GC()

	f1, err := os.Create("mem.pprof")
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	if err := pprof.WriteHeapProfile(f1); err != nil {
		panic(err)
	}
}

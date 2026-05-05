package practice_goroutines

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type Lottery struct {
	winningCombination []int
}

const LotterySize = 100_000
const MaxLenToShow = 10

func New(lptr *Lottery) bool {
	lptr.winningCombination = make([]int, LotterySize)
	for i := range len(lptr.winningCombination) {
		lptr.winningCombination[i] = 1 + rand.Intn(9)
	}
	return true
}

func (lptr *Lottery) Print(top int) {
	if top < len(lptr.winningCombination) {
		fmt.Println(lptr.winningCombination[:top], "...")
	} else {
		fmt.Println(lptr.winningCombination)
	}
}

type Result struct {
	index int
	value int
}

func crack_func(ind int, lptr *Lottery) int {
	for probe := 1; probe <= 9; probe++ {
		if lptr.winningCombination[ind] == probe {
			return probe
		}
	}
	return -1
}

func worker(lptr *Lottery, tasks <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for index := range tasks {
		results <- Result{
			index: index,
			value: crack_func(index, lptr),
		}
	}
}

func main() {
	fmt.Println("Запуск программы...")

	var l Lottery
	lptr := &l
	if ok := New(lptr); !ok {
		fmt.Println("Initialization error")
		return
	}

	show_topN := func(slice []int, top int) {
		if top < len(slice) {
			fmt.Print(slice[:top], " ...")
		} else {
			fmt.Print(slice)
		}
	}

	fmt.Print("Win combination to crack:")
	lptr.Print(MaxLenToShow)

	winCombination := make([]int, len(lptr.winningCombination))
	tasks := make(chan int, LotterySize)
	results := make(chan Result, LotterySize)
	time0 := time.Now()

	var numWorkers = runtime.NumCPU() // Используем все ядра CPU
	var wg sync.WaitGroup

	// Запускаем рабочих
	for range numWorkers {
		wg.Add(1)
		go worker(lptr, tasks, results, &wg)
	}

	// Отправляем задачи в отдельный горутине, чтобы не блокировать основной поток
	go func() {
		for i := range LotterySize {
			tasks <- i
		}
		close(tasks) // Закрываем канал задач после отправки всех индексов
	}()

	// Собираем результаты
	for range LotterySize {
		res := <-results
		winCombination[res.index] = res.value
	}

	wg.Wait()      // Ждём завершения всех рабочих
	close(results) // Закрываем канал результатов

	time1 := time.Since(time0)

	fmt.Print("Win combination cracked: ")
	show_topN(winCombination, MaxLenToShow)
	fmt.Println(" by", time1)

	// Дополнительная проверка: убедимся, что всё совпало
	matchCount := 0
	for i := 0; i < LotterySize; i++ {
		if winCombination[i] == lptr.winningCombination[i] {
			matchCount++
		}
	}
	fmt.Printf("Совпало %d из %d элементов\n", matchCount, LotterySize)
}

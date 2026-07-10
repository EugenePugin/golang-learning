package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)

	return rnd
}

type PredictableFuncReturnType struct {
	value int64
	err   error
}

// Нужно изменить функцию обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.
func predictableFunc() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result_chan := make(chan int64, 1)

	time0 := time.Now()
	go func() {
		result_chan <- unpredictableFunc()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Error: time is up", time.Since(time0))
		return 0, ctx.Err()
	case res := <-result_chan:
		fmt.Println(time.Since(time0))
		return res, nil
	}
}

func buro_task() {
	fmt.Println("started")

	result, err := predictableFunc()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result:", result)
	}

	fmt.Println("ended")
}

package main

import (
	"fmt"
	"math/rand"
)

func arrays() {
	// создай массив, заполни рандомными значениями, выведи его членов на экран + длину
	var array1 [5]int

	for i := range array1 {
		array1[i] = rand.Intn(10)
		fmt.Print(array1[i], " ")
	}
	fmt.Println()
	array2 := [...]int{array1[0], array1[2], array1[4]}
	for i := range array2 {
		fmt.Print(array2[i], " ")
	}

}

func slices() {
	// create array with len of 5, and fill with rnd integers
	// create slice from 0 till 3
	var array [5]int
	for i := range array {
		array[i] = rand.Intn(10)
	}

	fmt.Println(len(array), array)
	slice := array[:3]
	fmt.Println(len(slice), cap(slice), slice)
	slice = append(slice, rand.Intn(10))
	fmt.Println(len(slice), cap(slice), slice)
	slice = append(slice, rand.Intn(10))
	fmt.Println(len(slice), cap(slice), slice)
	slice = append(slice, rand.Intn(10))
	fmt.Println(len(slice), cap(slice), slice)
	slice = append(slice, rand.Intn(10))
	fmt.Println(len(slice), cap(slice), slice)
	fmt.Println(len(array), array)

}

func maps() {
	// create map
	// create slice with len of 10, filled with random numbers
	// fill the map with unique values only
	var slice []int
	for range 10 {
		slice = append(slice, rand.Intn(10))
	}
	fmt.Println(slice)
	var m = make(map[int]int)

	for i := range slice {
		if val, ok := m[slice[i]]; ok {
			// fmt.Println("the key ", m[slice[i]], "exists")
			m[slice[i]] = val + 1
		} else {
			// fmt.Println("the key ", m[slice[i]], "does not exist - creating that now")
			m[slice[i]] = 1
		}
	}
	fmt.Println(m)
}

func sort_exercise() (result int) {
	// let's have a slice, filled with rand numbers
	// copy it to another slices, sort this 2nd one  and expose
	defer func() {
		result++
	}()

	slice1 := make([]int, 0)
	for range 10 {
		slice1 = append(slice1, rand.Intn(10))
	}
	fmt.Println(slice1)

	slice2 := make([]int, 10)
	copy(slice2, slice1)

	// sort.Ints(slice2)
	// slices.Sort(slice2)

	fmt.Println(slice2)
	for i := range len(slice2) {
		fmt.Print(slice2[len(slice2)-1-i], " ")
	}
	return 0
}

func ptr_exercise2(ptr_array *[10]int) {
	ptr_array[1] = 255
}
func ptr_exercise3(slice []int) {
	slice[2] = 777
}

func ptr_exercise() {
	// creare array, fil with rnd numbers
	// transfer to a function via ptr, change some interim item there
	// validate the change at the original function
	var array [10]int
	for i := range len(array) {
		array[i] = rand.Intn(10)
	}
	fmt.Println(array)
	ptr_exercise2(&array)
	fmt.Println(array)

	slice := make([]int, 10)
	slice = array[:]
	ptr_exercise3(slice)
	fmt.Println(array)

}

func anonymous_functions_with_closure() {
	// create slice1, fill with rnd numbers
	// implement an. function, replacing odd numbers to 1, and even numbers to 0
	// use it to fill the values of slice2
	slice1 := make([]int, 0)
	for range 10 {
		slice1 = append(slice1, rand.Intn(10))
	}
	fmt.Println(slice1)

	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	odd_even_processor := func() {
		slice := slice2[:]
		for i := range slice {
			if 0 == slice[i]%2 {
				slice[i] = 0
			} else {
				slice[i] = 1
			}
		}
	}

	odd_even_processor()

	// func(slice []int) {
	// 	for i := range len(slice) {
	// 		if 0 == slice[i]%2 {
	// 			slice[i] = 0
	// 		} else {
	// 			slice[i] = 1
	// 		}
	// 	}
	// }(slice2)

	// func() {
	// 	slice := slice2[:]
	// 	for i := range len(slice) {
	// 		if 0 == slice[i]%2 {
	// 			slice[i] = 0
	// 		} else {
	// 			slice[i] = 1
	// 		}
	// 	}
	// }()

	fmt.Println(slice2)

}

func mapsPractice() {
	// let's have a map[int]*int
	// organize its filling by random number
	const mapSize = 20
	myMap := make(map[int]int, mapSize)

	for range mapSize {
		key := rand.Intn(mapSize)
		value := rand.Intn(10)
		if _, ok := myMap[key]; !ok {
			myMap[key] = value
		}
	}

	fmt.Println(myMap)

}

package main

import (
	"fmt"
)

// func countNumbers(nums []int) string {
// 	// Ваш код
// 	var countAboveZero, countBelowZero, countEqualToZero int

// 	for i := 0; i < len(nums); i++ {
// 		switch {
// 		case nums[i] > 0:
// 			countAboveZero++
// 		case nums[i] < 0:
// 			countBelowZero++
// 		default:
// 			countEqualToZero++
// 		}
// 	}

// 	TextToReturn := "выше нуля: " + strconv.Itoa(countAboveZero) + ", ниже нуля: " + strconv.Itoa(countBelowZero) + ", равна нулю: " + strconv.Itoa(countEqualToZero)

// 	return TextToReturn
// }

func Some_function_to_test(a, b int) int {
	// if a == 0 {
	// 	return 5
	// }
	return a + b
}

// func test_practice() {
// 	fmt.Println("Let's test!")
// 	var myArrayA = [3]int{-2, 0, 5}
// 	var myArrayB = [3]int{-3, 1, 6}

// 	for i := 0; i < 3; i++ {
// 		if myArrayA[i]+myArrayB[i] == Some_function_to_test(myArrayA[i], myArrayB[i]) {
// 			fmt.Println("Test[", i, "] passed")
// 		} else {
// 			fmt.Println("Test[", i, "] failed")
// 		}
// 	}

// 	// input, err := readWithTimeout(5 * time.Second)
// 	//     if err != nil {
// 	//         fmt.Println(err)
// 	//     } else {
// 	//         fmt.Printf("Вы ввели: %s\n", input)
// 	//     }
// }

// func ifDiscovered(value int, map1 map[int]int) bool {
// 	for i := 0; i < len(map1); i++ {
// 		if value == map1[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

func main() {

	// workMap := make(map[int]int)

	// workMap[0] = 1
	// workMap[1] = 14
	// workMap[2] = 44

	// fmt.Println(workMap)
	// fmt.Println(len(workMap))

	// fmt.Println("Enter a value to look for")
	// var n int
	// fmt.Scan(&n)

	// var msg string
	// if ifDiscovered(n, workMap) {
	// 	msg = "Bingo!"
	// } else {
	// 	msg = "Pardon moi..."
	// }
	// fmt.Println(msg)

	var friends0fDima []string
	friends0fSemyon := make([]string, 3)
	friends0fDima = append(friends0fDima, "Костя", "Семён", "Таня")
	friends0fSemyon = append(friends0fSemyon, "Валера", "Таня", "Дима")
	friends := map[string][]string{
		"Dima":   friends0fDima,
		"Semyon": friends0fSemyon,
		"Костя":  nil,
	}
	var value bool
	fmt.Println(value)
	fmt.Println(friends["Костя"])
	_, value = friends["Костя"]
	fmt.Print(value, " ")
	delete(friends, "Dima")
	fmt.Print(friends0fSemyon[3])
}

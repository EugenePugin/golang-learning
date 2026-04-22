//https://leetcode.com/problems/two-sum/description/

package main

import (
	"fmt"
	"math"
)

func twoSum(nums []int, target int) []int {
	// input validation check
	if len(nums) < 2 || len(nums) > 10000 {
		return nil
	}
	if math.Abs(float64(target)) > math.Pow(10, 9) {
		return nil
	}

	for i := range len(nums) {
		if math.Abs(float64(nums[i])) > math.Pow(10, 9) {
			return nil
		}
	}

	// logic itself

	// result := make([]int, 2)
	// for i := range len(nums) - 1 {
	// 	// fmt.Println("i=", i)
	// 	firstItem := nums[i]
	// 	// fmt.Println("    firstItem=", firstItem)
	// 	for j := i + 1; j < len(nums); j++ {
	// 		// fmt.Println("    j=", j)
	// 		// fmt.Println("    item to check:", nums[j])
	// 		if firstItem+nums[j] == target {
	// 			result[0] = i
	// 			result[1] = j
	// 			return result
	// 		}
	// 	}
	// }

	// 2. Создаем мапу: ключ — значение числа, значение — его индекс в массиве
	// map[Value]Index
	seenNumbers := make(map[int]int)

	// 3. Проходим по массиву один раз
	for i, currentNum := range nums {
		fmt.Println("i=", i)
		// Вычисляем, какое число нам нужно найти, чтобы получить target
		complement := target - currentNum
		fmt.Println("complement=", complement)

		// Проверяем, встречали ли мы это "дополнение" ранее
		if complementIndex, ok := seenNumbers[complement]; ok {
			// Если нашли, возвращаем индексы:
			// индекс найденного ранее числа и текущий индекс
			fmt.Println("Bingo!")
			return []int{complementIndex, i}
		}

		// Если не нашли, сохраняем текущее число и его индекс в мапу
		seenNumbers[currentNum] = i
	}

	// Если решение не найдено (по условию LeetCode оно всегда есть)
	return nil
	// for i:=range len(mapFacility) {
	// 	fmt.Println("i:",i)
	// 	firstItem := mapFacility[i]
	// 	for j:=1;j<len(mapFacility);j++} {
	// 		if val, ok = mapFacility[j] == target - firstItem {
	// 			result[0]
	// 		}
	// 	}

	// }

}

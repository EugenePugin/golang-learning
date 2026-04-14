// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
package lc8

// func FindMinAndGetItsIndex(prices []int) (int, int) {
// 	// fmt.Println("len(prices):", len(prices))
// 	bestPriceToBuy := prices[0]
// 	var bestDayToBuy int
// 	for i := 1; i < len(prices); i++ {
// 		if bestPriceToBuy > prices[i] {
// 			bestPriceToBuy = prices[i]
// 			bestDayToBuy = i
// 		}
// 	}
// 	switch bestDayToBuy {
// 	case len(prices) - 1: // the solution with the last day should be disqualified
// 		// fmt.Println("min solution is not qualified - need to find a better one")
// 		bestPriceToBuy, bestDayToBuy = FindMinAndGetItsIndex(prices[:(len(prices) - 1)])
// 	case 0:
// 		break
// 	}

// 	// fmt.Println(bestPriceToBuy, bestDayToBuy)

// 	return bestPriceToBuy, bestDayToBuy
// }

// func FindMaxAndGetItsIndex(slice []int) (int, int) {
// 	// var bestDayToBuy
// 	maxPrice := slice[0]
// 	var maxPriceIdx int
// 	for i := 1; i < len(slice); i++ {
// 		if maxPrice < slice[i] {
// 			maxPrice = slice[i]
// 			maxPriceIdx = i
// 		}
// 	}

// 	return maxPrice, maxPriceIdx
// }

// func generateSlicesOfProfits(slice []int) []int {
// 	sliceProfits := make([]int, len(slice)-1)
// 	for i := 1; i < len(slice); i++ {
// 		sliceProfits[i-1] = slice[i] - slice[0]
// 	}
// 	return sliceProfits
// }

// func getMaxProfitForThisSlice_v0(slice []int) int { //works well, but deadly slow
// 	sliceProfits := make([]int, len(slice)-1)

// 	for i := 1; i < len(slice); i++ {
// 		sliceProfits[i-1] = slice[i] - slice[0]
// 	}
// 	maxProfit := slices.Max(sliceProfits)
// 	if maxProfit <= 0 {
// 		return 0
// 	}
// 	return maxProfit
// }

func getMaxProfitForThisSlice(slice []int) int {
	var maxProfit int
	for i := 1; i < len(slice); i++ {
		if maxProfit < slice[i]-slice[0] {
			maxProfit = slice[i] - slice[0]
		}
	}
	return maxProfit
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// fmt.Println("Step #", i)

		// Если нашли цену ниже текущего минимума, обновляем его
		if prices[i] < minPrice {
			// fmt.Println("Wow, better price discovered")
			minPrice = prices[i]
		} else {
			// Иначе проверяем, получим ли мы больше прибыли, продавая сегодня
			currentProfit := prices[i] - minPrice
			// fmt.Println("currentProfit", currentProfit)
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
		// fmt.Println("maxProfit", maxProfit)
	}

	return maxProfit
}

func maxProfit_v0(prices []int) int {
	// if len(prices) > 100000 || len(prices) == 0 {
	// 	return 0
	// }

	// if slices.Min(prices) == slices.Max(prices) {
	// 	return 0
	// }

	// if isDescending(prices) { //eсли массив не является возрастющим , то профита нет
	// 	return 0
	// }
	// time_tracking := make([]int,len(slices))
	var maxProfitValue int
	// time0 := time.Now()

	for i := 0; i < len(prices)-1; i++ {
		// fmt.Println("Step #", i)
		profitValue := getMaxProfitForThisSlice(prices[i:])
		if maxProfitValue < profitValue {
			maxProfitValue = profitValue
		}
		// fmt.Println(sliceCurrent, "maxProfitValue:", maxProfitValue, "profitValue:", profitValue)
	}
	// fmt.Println(sliceCurrent, "maxProfitValue:", maxProfitValue, "profitValue:", profitValue)
	// time1 := time.Since(time0).Milliseconds()
	// fmt.Println(time1)

	return maxProfitValue
}

// func isDescending(arr []int) bool {
// 	for i := 0; i < len(arr)-1; i++ {
// 		if arr[i] <= arr[i+1] {
// 			return false
// 		}
// 	}
// 	return true
// }

// func maxProfit_v1(prices []int) int {

// 	if slices.Min(prices) == slices.Max(prices) {
// 		return 0
// 	}
// 	if isDescending(prices) { //eсли массив не является возрастющим , то профита нет
// 		return 0
// 	}
// 	// fmt.Println("There is a solution here!")

// 	var MinValue, MinValueIdx int

// 	MinValue, MinValueIdx = FindMinAndGetItsIndex(prices)
// 	fmt.Println("Mins:", MinValue, MinValueIdx)

// 	var MaxValueR, MaxValueRIdx int
// 	MaxValueR, MaxValueRIdx = FindMaxAndGetItsIndex(prices[(MinValueIdx + 1):])
// 	MaxValueRIdx += MinValueIdx + 1
// 	fmt.Println("Max's:", MaxValueR, MaxValueRIdx)
// 	maxProfitValue := MaxValueR - MinValue

// 	// fmt.Println("maxProfitValue:", maxProfitValue)
// 	return maxProfitValue
// }

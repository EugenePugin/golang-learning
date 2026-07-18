package HH_Go_exam

import (
	"fmt"
	"strconv"
	"strings"
)

func task1(url string) string {
	// find value of user parameter OR not found
	errReturnString := "not found"
	if len(url) == 0 {
		return errReturnString
	}
	_, query, ok := strings.Cut(url, "?")
	if !ok {
		return errReturnString
	}

	// Разбиваем строку запроса на отдельные параметры по символу "&"
	for _, param := range strings.Split(query, "&") {
		// Разделяем каждый параметр на "ключ" и "значение"
		key, value, found := strings.Cut(param, "=")

		// Проверяем строгое совпадение ключа
		if found && key == "user" {
			if value == "" {
				return errReturnString
			}
			return value
		}
	}
	return errReturnString
}

func AlaMainT1() {
	// fmt.Println("Hey!")
	url := "https://shop.com/orders?user=789&sort=date "
	fmt.Println(task1(url))
}

// Описываем структуру диапазона оценок
type MarkRange struct {
	Min   int
	Max   int
	Label string
}

// Декларативно объявляем все наши диапазоны.
// Если в будущем добавятся новые диапазоны, достаточно будет просто дописать строку сюда!
var markRanges = []MarkRange{
	{Min: 0, Max: 20, Label: "от 0 до 20"},
	{Min: 21, Max: 40, Label: "от 21 до 40"},
	{Min: 41, Max: 60, Label: "от 41 до 60"},
	{Min: 61, Max: 80, Label: "от 61 до 80"},
	{Min: 81, Max: 100, Label: "от 81 до 100"},
}

func task2(allTheMarks string) []string {
	parts := strings.Fields(allTheMarks)
	if len(parts) == 0 {
		return nil
	}
	counts := make([]int, len(markRanges))
	var total int
	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil || n < 0 || n > 100 {
			return nil
		}
		for i, r := range markRanges {
			if n >= r.Min && n <= r.Max {
				counts[i]++
				total++
				break
			}
		}
	}
	result := make([]string, len(markRanges))

	for i, r := range markRanges {
		ratio := 100 * float64(counts[i]) / float64(total)
		result[i] = fmt.Sprintf("число оценок %s: %d доля: %.1f%%", r.Label, counts[i], ratio)
	}
	return result
}

func task2old(allTheMarks string) []string {
	// return num of marks and their ratio
	parts := strings.Fields(allTheMarks)
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		nums = append(nums, n)
	}

	var nums0_20, nums21_40, nums41_60, nums61_80, nums81_100 int
	for i := range nums {
		if nums[i] < 0 {
			return nil
		}
		if nums[i] >= 0 && nums[i] <= 20 {
			nums0_20++
		}
		if nums[i] >= 21 && nums[i] <= 40 {
			nums21_40++
		}
		if nums[i] >= 41 && nums[i] <= 60 {
			nums41_60++
		}
		if nums[i] >= 61 && nums[i] <= 80 {
			nums61_80++
		}
		if nums[i] >= 81 && nums[i] <= 100 {
			nums81_100++
		}
	}
	total := nums0_20 + nums21_40 + nums41_60 + nums61_80 + nums81_100
	rat0_20 := float64(nums0_20) / float64(total)
	rat21_40 := float64(nums21_40) / float64(total)
	rat41_60 := float64(nums41_60) / float64(total)
	rat61_80 := float64(nums61_80) / float64(total)
	rat81_100 := float64(nums81_100) / float64(total)

	result := make([]string, 5)
	result[0] = "число оценок от 0 до 20: " + strconv.Itoa(nums0_20) + " доля: " + strconv.FormatFloat(rat0_20, 'f', 1, 64)
	result[1] = "число оценок от 21 до 40: " + strconv.Itoa(nums21_40) + " доля: " + strconv.FormatFloat(rat21_40, 'f', 1, 64)
	result[2] = "число оценок от 41 до 60: " + strconv.Itoa(nums41_60) + " доля: " + strconv.FormatFloat(rat41_60, 'f', 1, 64)
	result[3] = "число оценок от 61 до 80: " + strconv.Itoa(nums61_80) + " доля: " + strconv.FormatFloat(rat61_80, 'f', 1, 64)
	result[4] = "число оценок от 81 до 100: " + strconv.Itoa(nums81_100) + " доля: " + strconv.FormatFloat(rat81_100, 'f', 1, 64)

	return result
}

func AlaMainT2() {
	allTheMarks := "21 14 55 23 66 44 99 84 100"
	t2Result := task2(allTheMarks)
	for i := range len(t2Result) {
		fmt.Println(t2Result[i])
	}
}

type Route struct {
	from string
	to   string
}

type RouteStat struct {
	maxRouteCnt         int
	uniqueRoutesCnt     int
	routesWith1transfer int
}

func t3(routes []Route) RouteStat {
	var result RouteStat
	routeStatsMap := make(map[Route]int) // transfer per route
	for i := range routes {
		if routes[i].from == "" || routes[i].to == "" {
			fmt.Println("Нет корректных данных")
			return result
		}
		routeStatsMap[routes[i]]++
	}

	var routesWith1Transfer, maxRoutesCnt int
	for _, v := range routeStatsMap {
		if maxRoutesCnt < v {
			maxRoutesCnt = v
		}
		if v == 1 {
			routesWith1Transfer++
		}
	}

	result = RouteStat{
		maxRouteCnt:         maxRoutesCnt,
		uniqueRoutesCnt:     len(routeStatsMap),
		routesWith1transfer: routesWith1Transfer,
	}
	return result
}

func AlaMainT3() {
	transferCount := 7
	routes := make([]Route, transferCount)
	routes[0] = Route{from: "Moscow", to: "SPb"}
	routes[1] = Route{from: "SPb", to: "Moscow"}
	routes[2] = Route{from: "Moscow", to: "SPb"}
	routes[3] = Route{from: "Kazan", to: "Moscow"}
	routes[4] = Route{from: "Moscow", to: "SPb"}
	routes[5] = Route{from: "Kazan", to: "Moscow"}
	routes[6] = Route{from: "Kazan", to: "Moscow"}
	for i := range routes {
		fmt.Println(routes[i])
	}

	result := t3(routes)

	localrouteStatsMap := make(map[Route]int) // transfer per route
	for i := range routes {
		localrouteStatsMap[routes[i]]++
	}
	for k, v := range localrouteStatsMap {
		if v == result.maxRouteCnt {
			fmt.Println(k.from, k.to, v)
		}
	}
	fmt.Println("Уникальных маршрутов: ", result.uniqueRoutesCnt)
	fmt.Println("Маршрутов с одной перевозкой: ", result.routesWith1transfer)

}

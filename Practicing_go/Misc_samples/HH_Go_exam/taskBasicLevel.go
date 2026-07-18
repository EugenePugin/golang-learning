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

func t3a(routes []Route) RouteStat {
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

func AlaMainT3a() {
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

	result := t3a(routes)

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

func isValidIP(ip string) bool {
	// fmt.Println("DBG_:",ip)
	parts := strings.Split(ip, ".")
	// fmt.Println("DBG_",parts)
	if len(parts) != 4 {
		return false
	}
	for r := range parts {
		if r < 0 || r > 255 {
			return false
		}
	}
	return true
}

func listSpamIPs(ips string, threshold int) {
	parts := strings.Split(ips, " ")
	ipMap := make(map[string]int) //ip counter
	for _, r := range parts {
		if !isValidIP(r) {
			continue
		}
		ipMap[r]++
	}
	
	for k, v := range ipMap {
		if v < threshold {
			continue
		}
		fmt.Println(k, v)
	}
}

func AlaMainT3b() {
	ips := "192.168.1.1 192.168.1.1 10.0.0.5 8.8.8.8 172.16.0.100 172.16.2220.100 192.168.1.1 53.2"
	threshold := 3
	listSpamIPs(ips, threshold)

}

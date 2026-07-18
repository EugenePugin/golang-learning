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
	// fmt.Println(qPos)
	if !ok {
		return errReturnString
	}
	start := strings.Index(query, "user=")
	if start == -1 {
		return errReturnString
	}
	start += len("user=")

	end := strings.Index(query[start:], "&")
	// fmt.Println(end)

	if end == -1 {
		return query[start:]
	} else {
		return query[start : start+end]
	}
}

func AlaMainT1() {
	// fmt.Println("Hey!")
	url := "https://shop.com/orders?user=789&sort=date "
	fmt.Println(task1(url))
}

func task2(allTheMarks string) []string {
	// return num of marks and their ratio
	parts := strings.Fields(allTheMarks)
	nums := make([]int, 0, len(parts))
	for _, p := range parts {
		n, _ := strconv.Atoi(p)
		nums = append(nums, n)
	}
	var nums0_20, nums21_40, nums41_60, nums61_80, nums81_100 int
	for i := range nums {
		if nums[i] >= 0 && nums[i] <= 20 {
			nums0_20++
		}
		if nums[i] >= 20 && nums[i] <= 40 {
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

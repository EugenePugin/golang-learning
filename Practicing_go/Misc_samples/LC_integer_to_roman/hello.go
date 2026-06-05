package LC_integer_to_Roman

import "fmt"

func main() {
	fmt.Println("hey!")
	for num := 345; num < 1215; num++ {
		fmt.Println(num, "=>", intToRoman(num))
	}
}

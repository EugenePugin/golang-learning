package lc10

import "fmt"

func main() {
	wordsCircleArray := [...]string{"apple", "banana", "computer", "elephant", "guitar", "hospital", "island", "jacket", "kangaroo", "lemon", "mountain", "notebook", "ocean", "pencil", "rabbit", "sunshine", "tiger", "umbrella", "violin", "waterfall", "xylophone", "yacht", "zebra", "adventure"}

	wordsCircleArraySlice := make([]string, len(wordsCircleArray)-1)
	wordsCircleArraySlice = wordsCircleArray[:(len(wordsCircleArray) - 1)]
	startIndex := 1
	target := "elephant"
	fmt.Println(closestTarget(wordsCircleArraySlice, target, startIndex))

}

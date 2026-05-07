// T.Bank practice

// check slice of pointers to int
// double all the slice elements

package t_bank_interview

import (
	"fmt"
	"math/rand/v2"
)

func printInputs(inputs []*int) {
	for i := range inputs {
		fmt.Print(*inputs[i], " ")
	}
	fmt.Println()
}

func doubleInputs(inputs []*int) {
	mapInputs := make(map[*int]bool)

	for i := range inputs {
		fmt.Println(i)
		if _, ok := mapInputs[inputs[i]]; ok != true {
			mapInputs[inputs[i]] = true
			*inputs[i] *= 2
		} else {
			fmt.Println("skipping")
		}
	}

}

func main() {
	const inputsLen = 5
	inputs := make([]*int, inputsLen)
	for i := range inputsLen {
		tmpAnyNum := rand.IntN(10)
		// fmt.Print(tmpAnyNum, " ")
		inputs[i] = &tmpAnyNum
	}
	fmt.Println()
	inputs[inputsLen-1] = inputs[0]

	printInputs(inputs)
	doubleInputs(inputs)
	printInputs(inputs)

}

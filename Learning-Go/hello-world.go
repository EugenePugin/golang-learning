package main

import (
	"fmt"
	"strings"
)

func ToRemoveSymbolsMoreThenOnce() {
	var X string
	fmt.Scan(&X)
	X = strings.TrimSpace(X)
	runeX := []rune(X)
	var runeY []rune
	// fmt.Println("DBG1:", runeX)
	// fmt.Println("DBG1:", runeX[0])

	for i := range runeX {
		// fmt.Println(strings.Count(X, string(runeX[i])))
		if strings.Count(X, string(runeX[i])) == 1 {
			runeY = append(runeY, runeX[i])
		}
	}
	resultingString := string(runeY)
	fmt.Println(resultingString)

}

func main() {
	ToRemoveSymbolsMoreThenOnce()
}

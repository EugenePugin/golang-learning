package hello_ex

import (
	// пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	"strings"
)

type Battery struct {
	Capacity uint
}

type Stringer interface {
	String() string
}

func (a Battery) String() string {
	return ConvertCapUnitsToSymbols(a.Capacity)
}

func sample4stringer() {
	const capacityInputLen = 10
	const allowedSymbols = "01"

	var capacityInput string
	fmt.Scan(&capacityInput)

	//capacityInput = "0011100001"

	if len(capacityInput) != capacityInputLen {
		fmt.Println("The length must be", capacityInputLen)
		return
	}
	if !ContainsOnly(capacityInput, allowedSymbols) {
		fmt.Println("Not allowed symbols discovered")
		return
	}

	numberOfCrosses := uint(strings.Count(capacityInput, "1"))
	// fmt.Println(ConvertCapUnitsToSymbols(numberOfCrosses))

	var a Battery
	a.Capacity = numberOfCrosses
	fmt.Println(a)
}

func ConvertCapUnitsToSymbols(capacityUnits uint) string {
	var numberOfCrossesSymbols string
	switch capacityUnits {
	case 1:
		numberOfCrossesSymbols = "         X"
	case 2:
		numberOfCrossesSymbols = "        XX"
	case 3:
		numberOfCrossesSymbols = "       XXX"
	case 4:
		numberOfCrossesSymbols = "      XXXX"
	case 5:
		numberOfCrossesSymbols = "     XXXXX"
	case 6:
		numberOfCrossesSymbols = "    XXXXXX"
	case 7:
		numberOfCrossesSymbols = "   XXXXXXX"
	case 8:
		numberOfCrossesSymbols = "  XXXXXXXX"
	case 9:
		numberOfCrossesSymbols = " XXXXXXXXX"
	case 10:
		numberOfCrossesSymbols = "XXXXXXXXXX"
	case 0:
		numberOfCrossesSymbols = "          "
	}
	numberOfCrossesSymbols = "[" + numberOfCrossesSymbols + "]"
	return numberOfCrossesSymbols
}

// ContainsOnly checks if the input string 's' contains only characters present in 'allowedSymbols'.
func ContainsOnly(s, allowedSymbols string) bool {
	// Iterate over the input string by rune (character)
	for _, char := range s {
		// strings.IndeXAny returns -1 if the character is NOT found in allowedSymbols
		if !strings.ContainsRune(allowedSymbols, char) {
			return false // Found a character that is not allowed
		}
	}
	return true // All characters were found in the allowed set
}

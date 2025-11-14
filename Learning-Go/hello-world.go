package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	separator := ";"
	decimalSeparatorSymbolRussia := ","
	decimalSeparatorSymbolInternational := "."

	parts := strings.Split(s, separator)

	// fmt.Println("Raw string A", parts[0])
	purified := strings.ReplaceAll(parts[0], " ", "")
	purified = strings.ReplaceAll(purified, decimalSeparatorSymbolRussia, decimalSeparatorSymbolInternational)
	// fmt.Println("Spaces cleanup", purified)
	floatValueA, _ := strconv.ParseFloat(purified, 64)

	// fmt.Println("Raw string B", parts[1])
	purified = strings.ReplaceAll(parts[1], " ", "")
	purified = strings.ReplaceAll(purified, decimalSeparatorSymbolRussia, decimalSeparatorSymbolInternational)
	// fmt.Println("Spaces cleanup", purified)
	floatValueB, _ := strconv.ParseFloat(purified, 64)
	floatDiv := floatValueA / floatValueB
	fmt.Printf("%.4f", floatDiv)

}

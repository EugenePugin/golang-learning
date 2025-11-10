package main

import (
	"fmt"
	"strings"
)

func toFindIndex() {
	var X, S string
	fmt.Scan(&X)
	// reader := bufio.NewReader(os.Stdin)
	// X, _ := reader.ReadString('\n')
	X = strings.TrimSpace(X)
	// fmt.Println("DBG1:", X)

	// reader2 := bufio.NewReader(os.Stdin)
	// S, _ := reader2.ReadString('\n')
	fmt.Scan(&S)
	S = strings.TrimSpace(S)
	// fmt.Println("DBG2:", S)

	index := strings.Index(X, S)

	if index != -1 {
		// fmt.Printf("'%s' starts at byte index %d.\n", S, index)
		fmt.Println(index)
	} else {
		// fmt.Printf("'%s' was not found.\n", S)
		fmt.Println(-1)

	}

}

func main() {

	toFindIndex()
}

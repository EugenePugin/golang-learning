package one11

// "fmt"

// func one11a(dbgMode int) {
// 	var R float64

// 	if dbgMode == 1 { //dbg mode
// 		R = 1000001
// 		// -000012.2123
// 		//  12.12345678
// 	} else { //release mode
// 		fmt.Scan(&R)
// 	}

// 	if R <= 0 {
// 		formattedString := fmt.Sprintf("%4.2f", R)
// 		fmt.Println("число", formattedString, "не подходит")
// 	} else if R > 10000 {
// 		formattedString := fmt.Sprintf("%e", R)
// 		fmt.Println(formattedString)
// 	} else {
// 		convertedR := (float64)((int64)(R*R*10000)) / 10000
// 		formattedString := fmt.Sprintf("%.4f", convertedR)
// 		fmt.Println(formattedString)
// 	}
// }

// func main() {
// 	// one11a(1)
// }

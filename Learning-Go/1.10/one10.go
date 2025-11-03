func 1.10a {



	var a, b int
	// tmp
	fmt.Scan(&a)
	fmt.Scan(&b)

	// a = 1
	// b = 5

	// input data check
	if !(a < b) {
		fmt.Println("A must be < then B")
		return
	}
	if !(a < 100) || !(b < 100) {
		fmt.Println("Both A and B must be below 100")
		return
	}

	sum := a
	// functional part
	for i := a + 1; i <= b; i++ {
		sum += i
	}

	fmt.Println(sum)

}
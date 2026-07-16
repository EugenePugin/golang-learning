package wildberries

import "fmt"

func WbTechInterviewTask() {
	// three goroutines are writing to diff channels
	// we need to merge all three to the one and expose
	fmt.Println("Hey, WB!")

	a := make(chan string)
	b := make(chan string)
	c := make(chan string)
	merged_ch := make(chan string)

	go func() {
		for range 3 {
			a <- "a"
		}
		close(a)
	}()
	go func() {
		for range 2 {
			b <- "b"
		}
		close(b)
	}()
	go func() {
		for range 4 {
			c <- "c"
		}
		close(c)
	}()

	go func() {
		var val string
		ifOpen := true
		for {
			if (a == nil) && (b == nil) && (c == nil) {
				break
			}
			select {
			case val, ifOpen = <-a:
				if !ifOpen {
					a = nil
					continue
				}
			case val, ifOpen = <-b:
				if !ifOpen {
					b = nil
					continue
				}
			case val, ifOpen = <-c:
				if !ifOpen {
					c = nil
					continue
				}
			}
			merged_ch <- val
		}
		close(merged_ch)
	}()

	fmt.Println("Reader from merged_ch")
	for val := range merged_ch {
		fmt.Println(val)
	}

}

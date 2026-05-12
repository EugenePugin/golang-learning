package LC_concurrency

import (
	"fmt"
	"sync"
)

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
}

func (f *Foo) Second(printSecond func()) {
	/// Do not change this line
	printSecond()
}

func (f *Foo) Third(printThird func()) {
	// Do not change this line
	printThird()
}
func printFirst() {
	fmt.Print("first")
}
func printSecond() {
	fmt.Print("second")
}
func printThird() {
	fmt.Print("third")
}

func main() {
	nums := []int{3, 2, 1}
	foo := NewFoo()
	/*
		if len(nums) != 3 {
			fmt.Println("error: nums len must be equal to 3")
		}
		for i := range nums {
			if nums[i] < 1 || nums[i] > 3 {
				fmt.Println("error: nums[i] must be equal to 1 or 2 or 3")
			}
		}*/
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)

	for i := range nums {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// fmt.Println("id#", id, "is running")
			switch id {
			case 1:
				foo.First(printFirst)
				ch1 <- int(1)
				close(ch1)
			case 2:
				<-ch1
				foo.Second(printSecond)
				ch2 <- int(1)
				close(ch2)
			case 3:
				<-ch2
				foo.Third(printThird)
			default:
				panic("wow, that's not expected")
			}
		}(nums[i])
	}
	wg.Wait()
}

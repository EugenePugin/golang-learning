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

func (f *Foo) First(printFirst func() string) string {
	// Do not change this line
	return printFirst()
}

func (f *Foo) Second(printSecond func() string) string {
	/// Do not change this line
	return printSecond()
}

func (f *Foo) Third(printThird func() string) string {
	// Do not change this line
	return printThird()
}
func printFirst() string {
	return "first"
}
func printSecond() string {
	return "second"
}
func printThird() string {
	return "third"
}

func goroutine_runner(nums []int) string {

	if len(nums) != 3 {
		return "error: nums len must be equal to 3"
	}
	for i := range nums {
		if nums[i] < 1 || nums[i] > 3 {
			return "error: nums[i] must be equal to 1 or 2 or 3"
		}
	}
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)
	var result string
	for i := range nums {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// fmt.Println("id#", id, "is running")
			switch id {
			case 1:
				result = foo.First(printFirst)
				ch1 <- int(1)
				close(ch1)
			case 2:
				<-ch1
				result += foo.Second(printSecond)
				ch2 <- int(1)
				close(ch2)
			case 3:
				<-ch2
				result += foo.Third(printThird)
			default:
				panic("wow, that's not expected")
			}
		}(nums[i])
	}
	wg.Wait()
	return result
}

var foo *Foo

func main() {
	nums := []int{3, 2, 1}
	foo = NewFoo()

	fmt.Println(goroutine_runner(nums))
}

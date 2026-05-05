// https://leetcode.com/problems/container-with-most-water/description/

package LCwc

import (
	"math"
)

func maxAreaProxy(height []int) int {
	n := len(height)
	if n < 2 ||
		n > int(math.Pow10(5)) {
		return -1
	}
	for i := range height {
		if height[i] < 0 ||
			height[i] > int(math.Pow10(4)) {
			return -1
		}
	}
	return maxArea(height)
}

func maxAreaProxy_v0(height []int) int {
	n := len(height)
	if n < 2 ||
		n > int(math.Pow10(5)) {
		return -1
	}
	for i := range height {
		if height[i] < 0 ||
			height[i] > int(math.Pow10(4)) {
			return -1
		}
	}
	return maxArea_v0(height)
}

// brute force solution
func maxArea_v0(height []int) int {
	rectangeSquare := func(x, y int) int {
		return x * y
	}
	// fmt.Println("Given array:", height)
	var container, maxContainer int
	for i := range len(height) - 1 {
		// fmt.Println("i:", i)
		for j := i + 1; j < len(height); j++ {
			container = rectangeSquare(min(height[i], height[j]), j-i)
			// fmt.Println(i, j, "...", min(height[i], height[j]), j-i, "...", container)
			if maxContainer < container {
				maxContainer = container
			}
			// fmt.Println(maxContainer)
		}
	}
	return maxContainer
}

// optimised
func maxArea(height []int) int {
	rectangeSquare := func(x, y int) int {
		return x * y
	}
	var tmp int
	n := len(height)
	var container, maxContainer int
	for i := range n {
		// fmt.Println("Given array:", height)
		tmp = 0
		for j := n - 1; j > i; j-- {
			// fmt.Print("Indexes:", i, j, "...")
			if height[j] < tmp {
				// fmt.Println("not worth looking to it ... skipping")
				continue
			}
			container = rectangeSquare(min(height[i], height[j]), j-i)
			// fmt.Println("\th:", min(height[i], height[j]), "w:", j-i, "giving\t", container, "square units")
			if maxContainer < container {
				maxContainer = container
			}
			tmp = height[j]
			// fmt.Println(maxContainer)
		}

	}
	return maxContainer
}

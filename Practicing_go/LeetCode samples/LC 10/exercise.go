// https://leetcode.com/problems/shortest-distance-to-target-string-in-a-circular-array/description/?envType=daily-question&envId=2026-04-15
package lc10

import "sync"

func NextElement(words []string, curItemIdx int) string {
	// fmt.Println(words, curItemIdx)
	nextItemIdx := (curItemIdx + 1 + len(words)) % len(words)
	// fmt.Println(nextItemIdx)
	return words[nextItemIdx]
}
func PrevElement(words []string, curItemIdx int) string {
	// fmt.Println(words, curItemIdx)
	prevItemIdx := (curItemIdx - 1 + len(words)) % len(words)
	// fmt.Println(nextItemIdx)
	return words[prevItemIdx]
}

func ifDiscovered(words []string, target string) bool {
	for i := range len(words) {
		// fmt.Println(words[i], target)
		if target == words[i] {
			return true
		}
	}
	return false
}

func closestTarget(words []string, target string, startIndex int) int {

	// fmt.Println(words, target, startIndex)

	// special case
	if false == ifDiscovered(words, target) {
		return -1
	}
	if words[startIndex] == target {
		return 0
	}

	var ForwardStepCnt, BackwardStepCnt, curIdx int
	var valueToAssess string
	var wg sync.WaitGroup
	wg.Add(1)
	go func() { //forward path check
		defer wg.Done()
		curIdx = startIndex
		valueToAssess = words[curIdx]
		for {
			if valueToAssess == target {
				break
			}
			valueToAssess = NextElement(words, curIdx)
			curIdx++
			ForwardStepCnt++
		}
	}()
	wg.Add(1)
	go func() { //backward path check
		defer wg.Done()
		curIdx = startIndex
		valueToAssess = words[curIdx]
		for {
			if valueToAssess == target {
				break
			}
			valueToAssess = PrevElement(words, curIdx)
			curIdx--
			BackwardStepCnt++
		}
	}()

	wg.Wait()

	return min(ForwardStepCnt, BackwardStepCnt)
}

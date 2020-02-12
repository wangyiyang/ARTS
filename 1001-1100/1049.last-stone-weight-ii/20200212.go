package main

import "fmt"

func lastStoneWeightII(stones []int) int {
	hassum := [3001]bool{}
	sum := 0
	hassum[0] = true

	for _, v := range stones {
		sum += v
		for i := sum; i >= v; i-- {
			hassum[i] = hassum[i] || hassum[i-v]
		}
	}

	less := sum / 2
	for less >= 0 && !hassum[less] {
		less--
	}
	return sum - less - less
}

func main() {
	fmt.Println(lastStoneWeightII([]int{2, 7, 4, 1, 8, 1}))
}

package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	maxCandies := 0
	for i := 0; i < n; i++ {
		maxCandies = max(maxCandies, candies[i])
	}
	ret := make([]bool, n)
	for i := 0; i < n; i++ {
		ret[i] = candies[i]+extraCandies >= maxCandies
	}
	return ret
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}

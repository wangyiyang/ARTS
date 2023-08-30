package main

import "fmt"

func productExceptSelf(nums []int) []int {
	var res []int
	var left, right = 1, 1

	for i := 0; i < len(nums); i++ {
		res = append(res, left)
		left *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(a))
}
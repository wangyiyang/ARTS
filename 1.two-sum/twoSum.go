package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, item1 := range nums {
		for j, item2 := range nums[i+1:] {
			if item1+item2 == target {
				return []int{i, j + i + 1}
			}
		}
	}
	return []int{}
}

// nums = [2, 7, 11, 15], target = 9
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

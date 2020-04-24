package main

import "fmt"
import "sort"

func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	sort.Ints(nums) // 排序
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			tmp := nums[i] + nums[left] + nums[right]
			if tmp > target { // 偏大
				right--
			} else if tmp < target { // 偏小
				left++
			} else { // 一致
				return target
			}
			if distance(tmp, target) < distance(res, target) { // 距离更近则替换
				res = tmp
			}
		}
	}
	return res
}

func distance(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	fmt.Println(threeSumClosest([]int{1, 3, 4, 5, 7}, 2))
}

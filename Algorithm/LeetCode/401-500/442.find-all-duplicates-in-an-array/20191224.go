package main

import (
	"fmt"
	"math"
)

func findDuplicates(nums []int) []int {
	//tmpDic := make(map[int]int)
	//var result []int
	//for _,value := range nums {
	//	tmpDic[value] += 1
	//	if tmpDic[value] == 2 {
	//		result = append(result, value)
	//	}
	//}
	//return result

	result := make([]int, 0)
	for _, v := range nums {
		fmt.Println(nums)
		fmt.Println(result)
		v = int(math.Abs(float64(v)))
		if nums[v-1] > 0 {
			nums[v-1] = nums[v-1] * -1
		} else {
			result = append(result, v)
		}
	}
	return result

}

func main(){
	fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}))
}

package main

import "fmt"

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    var i int
    for j := 1; j < len(nums); j++ {
        if nums[i] != nums[j] {
            nums[i+1] = nums[j]
            i++
        }
    }
    return i + 1
}

func main() {
    fmt.Println(removeDuplicates([]int{1, 1, 2}))
}

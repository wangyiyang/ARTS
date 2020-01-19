package main

import "fmt"

func maxArea(height []int) int {
    ins := 0
    for i:=0; i<len(height)-1; i++ {
        for j:= i+1; j<len(height); j++{
            h:=0
            w:=j-i
            if height[i]<=height[j]{
                h = height[i]
            }else{
                h = height[j]
            }
            res := h*w 
            if ins < res {
                ins = res
            }
        }
    }
    return ins
}

func main() {
	fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

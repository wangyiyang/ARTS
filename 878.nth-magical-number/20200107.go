package main

import (
	"fmt"
	"math"
)

func NthMagicalNumber(N int, A int, B int) int {
	left := 0
	right := math.MaxInt64
	//最小公倍数
	minMul := A * B / MaxDiv(A, B)
	for left < right {
		mid := (left + right) / 2
		temp := mid/A + mid/B - mid/minMul
		if temp < N {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right % 1000000007
}

//求最大公约数
func MaxDiv(a, b int) int {
	if a < b {
		tmp := a
		a = b
		b = tmp
	}
	for a%b != 0 {
		t := b
		b = a % b
		a = t
	}
	return b
}

func main() {
	fmt.Println(NthMagicalNumber(1, 2, 3))
}

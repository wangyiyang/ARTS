package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {

	max := int(math.Pow(2, 31) - 1)

	_x := int(math.Abs(float64(x)))

	s := strconv.Itoa(_x)
	s = res(s)
	res, _ := strconv.Atoi(s)

	if res > max {
		return 0
	}

	if x > 0 {
		return res
	}

	return -res
}

func res(s string) string {

	len := len(s)
	res := ""
	for i := len - 1; i >= 0; i-- {
		res += string(s[i])
	}

	return res
}

func main() {
	fmt.Println(reverse(-312))
}

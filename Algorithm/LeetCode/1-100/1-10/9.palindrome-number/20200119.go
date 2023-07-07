package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	originalX := x
	reverseX := 0
	for x > 0 {
		reverseX = reverseX*10 + x%10
		x /= 10
	}
	return reverseX == originalX
}

func main() {
	fmt.Println(isPalindrome(121))
}

package main

import "fmt"

func longestPalindrome(s string) string {
	size := len(s)
	if size < 2 {
		return s
	}
	maxLen := 1
	res := string(s[0])

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if j-i+1 > maxLen && valid(s, i, j) {
				maxLen = j - i + 1
				res = s[i : j+1]
			}
		}
	}
	return res
}

func valid(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}

package main

import "fmt"

func valid(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		} else {
			left++
			right--
		}
	}

	return true
}

// func longestPalindrome(s string) int {
//  maxLen := 0
// 	chDic := make(map[string]int)
// 	for i := 0; i < len(s); i++ {
// 		chDic[string(s[i])] += 1
// 	}
// 	for _, value := range chDic {
// 		maxLen += value / 2 * 2
// 		if value/2 == 0 && value%2 == 1 {
// 			maxLen++
// 		}
// 	}
//
// 	return maxLen
// }

func longestPalindrome(s string) int {
	count := make([]int, 128)
	for _, c := range s {
		count[c]++
	}
	ans := 0
	for _, v := range count {
		ans += v / 2 * 2
		if ans%2 == 0 && v%2 == 1 {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
}

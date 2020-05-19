package main

import "fmt"

func validPalindrome(s string) bool {
	str := []byte(s)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-i-1] {
			j := l - i - 1
			return isPalindrome(str, i+1, j) || isPalindrome(str, i, j-1)
		}
	}
	return true
}
func isPalindrome(s []byte, l, r int) bool {
	for i := l; i <= l+(r-l)/2; i++ {
		if s[i] != s[r-i+l] {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(validPalindrome("abca"))
}

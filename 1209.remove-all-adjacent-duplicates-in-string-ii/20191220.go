package main

import "fmt"

func removeDuplicates(s string, k int) string {
	var tmpChar byte
	var tmpNum int
	var startIndex int
	for i := 0; i < len(s); i++ {
		if tmpChar == s[i] && tmpNum < k {
			tmpNum++
		} else if tmpChar != s[i] {
			tmpNum = 1
			tmpChar = s[i]
			startIndex = i
		}
		if tmpNum == k {
			return removeDuplicates(s[:startIndex]+s[i+1:], k)
		}
	}
	return s
}

func main() {
	fmt.Println(removeDuplicates("deeedbbcccbdaa", 3))
}

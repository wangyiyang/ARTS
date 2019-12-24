package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	tmp := make(map[string]int)
	if sLen == tLen {
		for _, value := range s {
			val := string(value)
			tmp[val] += 1
		}
		for _, value := range t {
			val := string(value)
			if tmp[val] == 1 {
				delete(tmp, val)
			} else if tmp[val] > 1 {
				tmp[val] -= 1
			} else {
				return false
			}
		}
		if tmp == nil {
			return true
		}
	} else {
		return false
	}
	return true
}

func main() {
	fmt.Println(isAnagram("rac", "car"))
}

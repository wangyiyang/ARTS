package main

import "fmt"

func partition(s string) [][]string {
	maxLen := len(s)
	var result [][]string
	tmp := ""
	for i :=0; i<(len(s)); i++ {
		var tmpSlice []string
		for j:= i;j<=len(s);j++{
			tmp += s[i:j]
			if j-i+1 > maxLen,valid(s,i,j){
				tmpSlice = append(tmpSlice,tmp)
			}
		}
		result = append(result, tmpSlice)
	}
	return result
}

func valid(s string,left,right int) bool{
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
	fmt.Println(partition("abacd"))
}

package main

import "fmt"

var letters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func shiftingLetters(S string, shifts []int) string {
	bs := []byte(S)
	move := 0
	for i := len(shifts) - 1; i >= 0; i-- {
		// 需要移动的位数，因为字母表是循环的，所以对26取模就是移动的次数
		move += (shifts[i] % 26)
		fmt.Println(move)
		bs[i] = letters[(int(bs[i]-'a')+move)%26]
	}
	return string(bs)
}

func main() {
	shifts := []int{3, 5, 9}
	S := shiftingLetters("abc", shifts)
	fmt.Println(S)
}

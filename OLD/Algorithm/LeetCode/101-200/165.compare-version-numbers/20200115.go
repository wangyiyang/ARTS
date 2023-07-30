package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	ret := 1
	i := 0
	for i = 0; i < len(v1) && i < len(v2); i++ {
		a, _ := strconv.Atoi(v1[i])
		b, _ := strconv.Atoi(v2[i])
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
	}
	if len(v1) < len(v2) {
		ret = -1
		v1 = v2
	}
	for ; i < len(v1); i++ {
		a, _ := strconv.Atoi(v1[i])
		if a != 0 {
			return ret
		}
	}
	return 0
}

func main() {
	fmt.Println(compareVersion("1.5.6.7.7.8.9", "1.00000000000000000000000000001"))
}

package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	//拿出第一个，和后面的去比
	ans := strs[0]
	for i, v := range strs {
		//如果slice中有空串，说明没有公共前缀
		if v == "" {
			return ""
		}
		//略过第一个
		if i == 0 {
			continue
		}
		//找到公共最小长度
		var minLen int
		if len(ans) < len(v) {
			minLen = len(ans)
		} else {
			minLen = len(v)
		}
		//从后往前不断对比，直到找到了相同前缀的就跳出
		for j := minLen; j > 0; j-- {
			if strings.HasPrefix(v, ans[0:j]) {
				ans = ans[0:j]
				break
			}
			ans = ans[0:j]
			//如果找到最后一个都没有说明没有公共前缀，可以直接返回
			if j == 1 {
				return ""
			}
		}
	}
	return ans

}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

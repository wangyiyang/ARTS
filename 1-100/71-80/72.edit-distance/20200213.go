package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// dp[i][j] 代表了从 word1[:i] → word2[:j] 所需要的最小步骤
			// 按照题目给出的 3 种操作方法，分别讨论：
			// 1. 先 *删除* word1[:i] 最后的字母，得到 word1[:i-1]。
			//    再 word1[:i-1] → word2[:j] 此方法所需的步骤是
			// 	  1 + dp[i-1][j]
			// 2. 先 word1[:i] → word2[:j-1]，
			//    再 *添加* word2[j-1] 到 word2 的末尾，此方法所需的步骤是
			//    1 + dp[i][j-1]
			// 3. 先 word1[i-1] → word2[j-1]
			// 	   3.1 如果 word1[i-1] = word2[i-1] 的话
			//         无需 *替换*，
			//         总的步骤是 dp[i-1][j-1]
			// 	   3.2 如果 word1[i-1] != word2[i-1] 的话
			//         执行 *替换* 操作，把 word1[i-1] 替换成 word2[j-1]
			//         总的步骤是 1 + dp[i-1][j-1]
			dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])

			replace := 1
			if word1[i-1] == word2[j-1] {
				replace = 0
			}

			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+replace)
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

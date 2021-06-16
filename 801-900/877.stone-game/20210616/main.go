package main

import "fmt"

func stoneGame(piles []int) bool {
	length := len(piles)
    dp := make([][]int, length)
    for i := 0; i < length; i++ {
        dp[i] = make([]int, length)
        dp[i][i] = piles[i]
    }
    for i := length - 2; i >= 0; i-- {
        for j := i + 1; j < length; j++ {
            dp[i][j] = max(piles[i] - dp[i+1][j], piles[j] - dp[i][j-1])
        }
    }
    return dp[0][length-1] > 0


}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func main() {
	fmt.Println(stoneGame([]int{1,2,3,4}))
}

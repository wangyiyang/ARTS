package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	m   int
	n   int
	k   int
	ans int
}{

	{
		11,
		13,
		57,
		24,
	},

	{
		30000,
		30000,
		900000000,
		900000000,
	},

	{
		30000,
		3,
		90000,
		90000,
	},

	{
		300,
		300,
		90000,
		90000,
	},

	{
		3000,
		3000,
		9000000,
		9000000,
	},

	{
		3,
		8,
		19,
		14,
	},

	{
		3,
		6,
		13,
		9,
	},

	{
		3,
		3,
		5,
		3,
	},

	{
		2,
		3,
		6,
		6,
	},

	// 可以有多个 testcase
}

func findKthNumber(m int, n int, k int) int {
	// 后面的 count 函数，m 较小比较有利
	if m > n {
		m, n = n, m
	}

	lo, hi := 1, m*n
	for lo < hi {
		mid := lo + (hi-lo)/2
		c := count(mid, m, n)
		if c >= k {
			// 说明 mid 大了
			// 需要降低 hi
			hi = mid
		} else {
			// 与 hi 同理，提高 lo
			// +1 是为了避免死循环
			lo = mid + 1
		}
	}

	return hi
}

// v >= a * b , 1<=a<=m, 1<=b<=n
// count 返回 a 和 b 的所有可能组合的个数
func count(v, m, n int) int {
	c := 0
	for i := 1; i <= m; i++ {
		c += min(v/i, n)
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test_findKthNumber(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, findKthNumber(tc.m, tc.n, tc.k), "输入:%v", tc)
	}
}

func Benchmark_findKthNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findKthNumber(tc.m, tc.n, tc.k)
		}
	}
}

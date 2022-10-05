//go:build 60
// +build 60

package main

import "strconv"

func main() {
	getPermutation()
}

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 排列序列
 */

// @lc code=start
func getPermutation(n int, k int) string {
	f := make([]int, n)
	f[0] = 1
	for i := 1; i < n; i++ {
		f[i] = f[i-1] * i
	}
	k--
	ans := ""
	valid := make([]int, n+1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}
	for i := 1; i <= n; i++ {
		order := k/f[n-i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= f[n-i]
	}
	return ans
}

// @lc code=end

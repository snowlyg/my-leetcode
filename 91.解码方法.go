//go:build 91
// +build 91

package main

import "fmt"

func main() {
	s := "226"
	fmt.Println(numDecodings(s))
}

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	// 动态规划
	// 网格解法，i 字符长度，j 字符容量
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {

		// 前一位 != 0
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		// 大于 1 位，前位两位 != 0 ，两位数小于等于 26
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}

// @lc code=end

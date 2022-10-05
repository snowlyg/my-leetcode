//go:build 43
// +build 43

package main

import "strconv"

func main() {
	multiply("2", "3")
}

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {

	// 返回 0
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n)

	// 每位计算乘积
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ansArr[i+j+1] += x * y
		}
	}

	// 处理进位
	for i := m + n - 1; i > 0; i-- {
		ansArr[i-1] += ansArr[i] / 10
		ansArr[i] %= 10
	}

	ans := ""
	idx := 0
	if ansArr[0] == 0 {
		idx = 1
	}

	// 处理字符
	for ; idx < m+n; idx++ {
		ans += strconv.Itoa(ansArr[idx])
	}

	return ans
}

// @lc code=end

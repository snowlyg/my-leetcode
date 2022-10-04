//go:build 42
// +build 42

package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	n, res := len(height), 0
	l_max, r_max := make([]int, len(height)), make([]int, len(height))

	l_max[0] = height[0]
	r_max[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		l_max[i] = max(height[i], l_max[i-1])
	}
	for i := n - 2; i >= 0; i-- {
		r_max[i] = max(height[i], r_max[i+1])
	}
	for i := 1; i < n-1; i++ {
		res += min(l_max[i], r_max[i]) - height[i]
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

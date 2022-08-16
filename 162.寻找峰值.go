//go:build 162
// +build 162

package main

import "fmt"

func main() {
	nums := []int{1, 2}
	fmt.Println(findPeakElement(nums))
}

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
func findPeakElement(nums []int) int {
	idx := 0
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return idx
}

// @lc code=end

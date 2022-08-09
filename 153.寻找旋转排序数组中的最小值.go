//go:build 153
// +build 153

package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 1}
	fmt.Println(findMin(nums))
}

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	// 二分查找
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[end] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return nums[end]
}

// @lc code=end

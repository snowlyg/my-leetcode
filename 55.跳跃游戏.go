//go:build 55
// +build 55

package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}

	fmt.Println(canJump(nums))
}

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if i <= maxIndex {
			maxIndex = max(maxIndex, i+nums[i])
		} else {
			return false
		}
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

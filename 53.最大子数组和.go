//go:build 53
// +build 53

package main

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	sum := maxSubArray(nums)
	println(sum)
}

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	// 动态规划
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 相加大于前前值
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		// 替换最大值
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// @lc code=end

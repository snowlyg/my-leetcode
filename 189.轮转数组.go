//go:build 189
// +build 189

package main

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
}

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

// @lc code=end

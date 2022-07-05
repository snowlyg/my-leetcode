//go:build 33
// +build 33

package main

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	println(search(nums, target))
}

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	start, end := 0, len(nums)-1

	// 二分查找
	// 要注意数组被旋转过，有可能出现不是有序的
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[start] <= nums[mid] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}

// @lc code=end

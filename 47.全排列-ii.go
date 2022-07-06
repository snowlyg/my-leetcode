//go:build 47
// +build 47

package main

import "sort"

func main() {
	nums := []int{1, 2, 3}
	permuteUnique(nums)
}

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n) // 判断是否填入
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			// 当前位置已经填入，或者前一位没有填入并且数值和前一位相同
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			// 填入数值
			perm = append(perm, v)
			// 设置位置i已经被填入
			vis[i] = true
			backtrack(idx + 1)

			// 释放位置 i
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

// @lc code=end

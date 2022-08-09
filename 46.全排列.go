//go:build 46
// +build 46

package main

func main() {
	nums := []int{1, 2, 3}
	p := permute(nums)
	if p != nil {
		for i := 0; i < len(nums); i++ {
			if p[i] == nil {
				continue
			}
			for _, v := range p[i] {
				println(v)
			}
			println("================")
		}
	}
}

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	option := make([]int, len(nums))
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, option...))
			return
		}
		for i := idx; i < len(nums); i++ {
			option[idx] = nums[i]
			nums[i], nums[idx] = nums[idx], nums[i]
			backtrack(idx + 1)
			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}
	backtrack(0)
	return ans
}

// @lc code=end

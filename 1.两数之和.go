//go:build 1
// +build 1

package main

func main() {
	twoSum([]int{1, 32, 5, 6}, 5)
}

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	inTable := map[int]int{}
	for i, v := range nums {
		if n, ok := inTable[target-v]; ok {
			return []int{n, i}
		}
		inTable[v] = i
	}
	return nil
}

// @lc code=end

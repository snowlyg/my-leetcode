//go:build 118
// +build 118

package main

import "fmt"

func main() {
	res := generate(4)
	for _, v := range res {
		fmt.Println(v)
	}
}

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		s := make([]int, i+1)
		s[0] = 1
		if i > 0 {
			s[i] = 1
			lastLine := res[i-1]
			for j := 1; j < len(lastLine); j++ {
				s[j] = lastLine[j-1] + lastLine[j]
			}
		}

		res[i] = s
	}
	return res
}

// @lc code=end

//go:build 73
// +build 73

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Printf("%+v\n", matrix)
}

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	// 矩阵长高
	n, m := len(matrix), len(matrix[0])
	// 第一列是否有 0
	col0 := false
	for _, r := range matrix {
		// 行首为 0
		if r[0] == 0 {
			col0 = true
		}

		// 循环第一行，有0 则设置为 0 作为标记
		for j := 1; j < m; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 把标记好的行，列设置为0
	for i := n - 1; i >= 0; i-- {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		// 最后处理第一列
		if col0 {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end

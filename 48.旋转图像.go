//go:build 48
// +build 48

package main

import "fmt"

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	fmt.Printf("%+v\n", matrix)
}

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	// 先把二维矩阵沿对角线反转，然后反转矩阵的每一行
	xLen := len(matrix)
	if xLen == 0 {
		return
	}
	for i := 0; i < xLen; i++ {
		for j := i; j < xLen; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for _, v := range matrix {
		reverse(v)
	}
}

// 反转二维矩阵的每一行
func reverse(m []int) {
	i, j := 0, len(m)-1
	for j > i {
		m[i], m[j] = m[j], m[i]
		i++
		j--
	}
}

// @lc code=end

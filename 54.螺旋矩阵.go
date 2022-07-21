//go:build 54
// +build 54

package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	r := spiralOrder(matrix)

	fmt.Printf("%+v\n", r)

}

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	//排查空数据
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	// 行列长度
	rows, columns := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, columns)
	}

	var (
		total          = rows * columns
		order          = make([]int, total)
		row, column    = 0, 0
		directions     = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 方向集合
		directionIndex = 0                                         // 方向
	)

	for i := 0; i < total; i++ {
		order[i] = matrix[row][column] // 当前位置
		visited[row][column] = true    // 是否访问过
		nextRow, nextColumn := row+directions[directionIndex][0], column+directions[directionIndex][1]
		if nextRow < 0 || nextRow >= rows || nextColumn < 0 || nextColumn >= columns || visited[nextRow][nextColumn] {
			directionIndex = (directionIndex + 1) % 4
		}
		row += directions[directionIndex][0]
		column += directions[directionIndex][1]
	}
	return order

}

// @lc code=end

//go:build 200
// +build 200

package main

import "log"

func main() {

	grids := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	log.Println(numIslands(grids))
}

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	var lands int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				lands++
				dfs(grid, i, j)
			}
		}
	}
	return lands
}

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func dfs(grid [][]byte, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	for _, dir := range dirs {
		dfs(grid, i+dir[0], j+dir[1])
	}

}

// @lc code=end

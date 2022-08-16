//go:build 168
// +build 168

package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("A"))
}

/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel 表列序号
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	number := 0
	for i, multiple := len(columnTitle)-1, 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		number += int(k) * multiple
		multiple *= 26
	}
	return number
}

// @lc code=end

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	version1, version2 := "1.001", "1.01"
	fmt.Println(compareVersion(version1, version2))

}

/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	ver1 := strings.Split(version1, ".")
	ver2 := strings.Split(version2, ".")
	l1 := len(ver1)
	l2 := len(ver2)
	i := 0
	for i < l1 || i < l2 {
		x, y := 0, 0
		if i < l1 {
			x, _ = strconv.Atoi(ver1[i])
		}

		if i < l2 {
			y, _ = strconv.Atoi(ver2[i])
		}

		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
		i++
	}

	return 0
}

// @lc code=end

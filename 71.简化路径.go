//go:build 71
// +build 71

package main

import (
	"fmt"
	"strings"
)

func main() {
	paths := []string{"/home//foo/", "/home/", "/a/./b/../../c/", "/a/../../b/../c//.//"}
	for _, v := range paths {
		fmt.Println(simplifyPath(v))
	}
}

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	newPaths := []string{}
	spe := "/"
	// 分隔字符 /
	paths := strings.Split(path, spe)
	for _, v := range paths {
		if v == "" || v == "." {
			continue
		}
		if v == ".." {
			if len(newPaths) > 0 {
				newPaths = newPaths[:len(newPaths)-1]
			}
		} else {
			newPaths = append(newPaths, v)
		}
	}

	return spe + strings.Join(newPaths, spe)
}

// @lc code=end

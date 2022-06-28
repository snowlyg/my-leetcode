//go:build 5
// +build 5

package main

import (
	"strings"
)

func main() {
	myAtoi("  354354")
}

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	negative := false
	abs := s

	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		abs = s
	case '-':
		negative = true
		abs = s[1:]
	case '+':
		abs = s[1:]
	default:
		return 0
	}

	absNum := 0
	for _, ch := range abs {
		ch -= '0' // 转换数字
		if ch > 9 || ch < 0 {
			break
		}
		absNum = absNum*10 + int(ch)
		if absNum >= 1<<31 {
			absNum = 1 << 31
			break
		}
	}

	if negative {
		absNum = -absNum
	} else {
		if absNum >= 1<<31-1 {
			absNum = 1<<31 - 1
		}
	}

	return absNum
}

// @lc code=end

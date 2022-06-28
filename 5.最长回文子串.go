//go:build 4
// +build 4

package main

func main() {
	longestPalindrome("jfjkkk")
}

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	// 中心扩散
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 区分奇数和偶数
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		// 同时扩散两组数据，比较回文子串长度
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	// 当左边大于等于0，并且右边小于字符长度，同时左边等于右边时，向外扩散一次
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	// 没有符合条件情况，返回扩散前数值
	return left + 1, right - 1
}

// @lc code=end

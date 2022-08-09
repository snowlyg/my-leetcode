//go:build 151
// +build 151

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 颠倒字符串中的单词
 */

// @lc code=start
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	s = reverse(s)
	sub := strings.Split(s, " ")
	res := []string{}
	for _, v := range sub {
		vt := strings.Trim(v, " ")
		if vt != "" {
			res = append(res, reverse(vt))
		}
	}
	return strings.Join(res, " ")
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

// @lc code=end

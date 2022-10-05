//go:build 140
// +build 140

package main

import "fmt"
import "strings"

func main() {
	wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	fmt.Println(wordBreak("pineapplepenapple", wordDict))
}

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

//  回溯算法

// @lc code=start
func wordBreak(s string, wordDict []string) []string {
	res := []string{}
	track := []string{}

	// 回溯算法框架
	var backtrack func(s string, i int, wordDict []string)
	backtrack = func(s string, i int, wordDict []string) {
		// 完整拼接后返回
		if i == len(s) {
			res = append(res, strings.Join(track, " "))
			return
		}

		if i > len(s) {
			return
		}

		// 遍历所有单词
		for _, word := range wordDict {
			// 跳过太长单词
			l := len(word)
			if i+l > len(s) {
				continue
			}

			// 匹配失败
			sub := s[i : i+l]
			if sub != word {
				continue
			}

			// 匹配成功，选择
			track = append(track, word)
			
			backtrack(s, i+l, wordDict)

			// 撤销选择
			if len(track) > 0 {
				track = track[:len(track)-1]
			}
		}
	}
	backtrack(s, 0, wordDict)
	return res
}

// @lc code=end

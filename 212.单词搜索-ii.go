// + build 212
package main

import "log"

func main() {
	board := [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}
	words := []string{"oath", "pea", "eat", "rain"}
	for _, word := range findWords(board, words) {
		log.Println(word)
	}
}

/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start

// 字典树
type Trie struct {
	children [26]*Trie
	word     string
}

// Insert 插入字典
func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.word = word
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func findWords(board [][]byte, words []string) []string {
	t := &Trie{}

	for _, word := range words {
		t.Insert(word)
	}

	m, n := len(board), len(board[0])
	seen := map[string]bool{} //去重

	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y] //取出
		node = node.children[ch-'a']
		if node == nil {
			return
		}

		if node.word != "" {
			seen[node.word] = true
		}

		board[x][y] = '#' //防止回头
		for _, d := range dirs {
			nx, ny := x+d.x, y+d.y
			if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
				dfs(node, nx, ny)
			}
		}
		board[x][y] = ch //回填
	}
	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}

	ans := make([]string, 0, len(seen))
	for s := range seen {
		ans = append(ans, s)
	}
	return ans
}

// @lc code=end

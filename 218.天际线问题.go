//go:build 218
// +build 218

package main

import (
	"container/heap"
	"sort"
)

func main() {
	buildings := [][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}
	getSkyline(buildings)

}

/*
 * @lc app=leetcode.cn id=218 lang=golang
 *
 * [218] 天际线问题
 */

// 具体问题是面积问题，坐标问题

// @lc code=start

type pair struct{ right, height int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].height > h[j].height }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func getSkyline(buildings [][]int) (ans [][]int) {

	n := len(buildings)

	// 二维化数组，左右值
	boundaries := make([]int, 0, n*2)
	for _, building := range buildings {
		boundaries = append(boundaries, building[0], building[1])
	}

	// 排序数组
	sort.Ints(boundaries)

	idx := 0
	h := hp{}

	// 循环左右值
	for _, boundary := range boundaries {

		// 最小左边节点
		for idx < n && buildings[idx][0] <= boundary {
			heap.Push(&h, pair{buildings[idx][1], buildings[idx][2]})
			idx++
		}

		// 右侧节点小于数组移除
		for len(h) > 0 && h[0].right <= boundary {
			heap.Pop(&h)
		}

		// 最高值
		maxn := 0
		if len(h) > 0 {
			maxn = h[0].height
		}

		// 不重复的数据填入
		if len(ans) == 0 || maxn != ans[len(ans)-1][1] {
			ans = append(ans, []int{boundary, maxn})
		}
	}

	return
}

// @lc code=end

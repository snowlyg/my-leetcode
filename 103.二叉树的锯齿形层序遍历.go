//go:build 103
// +build 103

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	fmt.Println(zigzagLevelOrder(root))
}

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	levelRes := map[int][]int{}
	level := 0
	var order func(root *TreeNode, level int)
	order = func(root *TreeNode, level int) {
		levelRes[level] = append(levelRes[level], root.Val)
		if root.Left != nil {
			order(root.Left, level+1)
		}
		if root.Right != nil {
			order(root.Right, level+1)
		}
	}
	order(root, level)
	res := make([][]int, len(levelRes))
	for i := 0; i < len(levelRes); i++ {
		llr := levelRes[i]
		if i%2 != 0 {
			lr := len(llr)
			// 反转数组
			for i, n := 0, lr; i < n/2; i++ {
				llr[i], llr[n-1-i] = llr[n-1-i], llr[i]
			}
		}
		res[i] = llr
	}

	return res
}

// @lc code=end

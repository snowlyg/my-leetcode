//go:build 102
// +build 102

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
	fmt.Println(levelOrder(root))
}

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
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
		res[i] = levelRes[i]
	}

	return res
}

// @lc code=end

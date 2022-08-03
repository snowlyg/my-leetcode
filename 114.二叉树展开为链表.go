//go:build 114
// +build 114

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
			Val:   5,
			Right: &TreeNode{Val: 6},
		},
	}
	flatten(root)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Right
	}
}

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 拉平左右节点
	flatten(root.Left)
	flatten(root.Right)

	// 把左节点切换到右边
	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = left

	// 把右节点接到当前链表后面
	n := root
	for n.Right != nil { //拿到最后的节点
		n = n.Right
	}
	n.Right = right
}

// @lc code=end

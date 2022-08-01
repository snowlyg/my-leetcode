//go:build 101
// +build 101

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
	fmt.Println(isSymmetric(root))
}

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	// 双指针递归
	if root == nil {
		return true
	}
	var checkNode func(left, right *TreeNode) bool
	checkNode = func(left, right *TreeNode) bool {
		if left == nil || right == nil {
			return left == right
		}

		if left.Val != right.Val {
			return false
		}
		return checkNode(left.Right, right.Left) && checkNode(left.Left, right.Right)
	}
	return checkNode(root.Left, root.Right)
}

// @lc code=end

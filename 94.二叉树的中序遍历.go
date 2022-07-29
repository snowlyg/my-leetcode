//go:build 94
// +build 94

package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}

	fmt.Printf("%+v\n", inorderTraversal(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历：根结点 ---> 左子树 ---> 右子树

// 中序遍历：左子树---> 根结点 ---> 右子树

// 后序遍历：左子树 ---> 右子树 ---> 根结点

// 层次遍历：只需按层次遍历即可

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return nil
	}

	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	fmt.Printf("%+v\n", res)
	return res
}

// @lc code=end

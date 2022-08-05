//go:build 124
// +build 124

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(maxPathSum(root))
}

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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

var res = math.MinInt

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 计算单边路径和时顺便计算最大路径和
	oneSideMax(root)
	return res
}

// 定义：计算从根节点 root 为起点的最大单边路径和
func oneSideMax(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftMax := max(0, oneSideMax(node.Left))
	rightMax := max(0, oneSideMax(node.Right))
	// 后序遍历位置，顺便更新最大路径和
	pathMax := node.Val + leftMax + rightMax
	res = max(res, pathMax)

	// 实现函数定义，左右子树的最大单边路径和加上根节点的值
	// 就是从根节点 root 为起点的最大单边路径和
	return max(leftMax, rightMax) + node.Val
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

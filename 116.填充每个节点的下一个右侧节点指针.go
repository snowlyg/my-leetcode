//go:build 116
// +build 116

package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	root := &Node{
		Val: 1,
		Left: &Node{
			Val:   2,
			Left:  &Node{Val: 4},
			Right: &Node{Val: 5},
		},
		Right: &Node{
			Val:   3,
			Left:  &Node{Val: 6},
			Right: &Node{Val: 7},
		},
	}
	connect(root)
	left := root
	right := root
	fmt.Println("right:")
	for right != nil {
		fmt.Println("-", right.Val)
		right = right.Right
	}

	fmt.Println("left:")
	for left != nil {
		fmt.Println("-", left.Val)
		left = left.Left
	}
}

/*
 * @lc app=leetcode.cn id=116 lang=golang
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	traverse(root.Left, root.Right)
	return root
}

func traverse(n1, n2 *Node) {
	if n1 == nil || n2 == nil {
		return
	}
	n1.Next = n2
	traverse(n1.Left, n1.Right) // 左节点
	traverse(n2.Left, n2.Right) // 右节点
	traverse(n1.Right, n2.Left) // 左右节点
}

// @lc code=end

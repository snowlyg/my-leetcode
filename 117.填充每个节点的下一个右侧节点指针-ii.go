//go:build 117
// +build 117

package main

import (
	"fmt"
)

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
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
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
	levelRes := map[int][]*Node{}
	level := 0
	var order func(root *Node, level int)
	order = func(root *Node, level int) {
		levelRes[level] = append(levelRes[level], root)
		if root.Left != nil {
			order(root.Left, level+1)
		}
		if root.Right != nil {
			order(root.Right, level+1)
		}
	}
	order(root, level)
	for i := 0; i < len(levelRes); i++ {
		if len(levelRes[i]) > 1 {
			for j := 0; j < len(levelRes[i])-1; j++ {
				levelRes[i][j].Next = levelRes[i][j+1]
			}
		}
	}
	return root
}

// @lc code=end

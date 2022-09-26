//go:build 173
// +build 173

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=173 lang=golang
 *
 * [173] 二叉搜索树迭代器
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
type BSTIterator struct {
	stack []*TreeNode
}

func (this *BSTIterator) pushLeftBranch(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{
		stack: make([]*TreeNode, 0, 8),
	}
	iter.pushLeftBranch(root)
	return iter
}

func (this *BSTIterator) Pop() *TreeNode {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return top
}

func (this *BSTIterator) Next() int {
	top := this.Pop()
	this.pushLeftBranch(top.Right)
	return top.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

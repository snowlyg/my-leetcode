//go:build 106
// +build 106

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	tn := buildTree(inorder, postorder)
	fmt.Printf("%+v\n", tn)
}

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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

var indexMap = map[int]int{}

func buildTree(inorder []int, postorder []int) *TreeNode {
	inLen := len(inorder)
	postLen := len(postorder)
	if inLen == 0 && postLen == 0 {
		return nil
	}
	sliceToMap(inorder)
	return buildNode(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func buildNode(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}

	rootVal := postorder[postEnd]
	root := &TreeNode{Val: rootVal}

	rootIndex := indexMap[rootVal]
	leftSize := rootIndex - inStart

	root.Left = buildNode(inorder, inStart, rootIndex-1, postorder, postStart, postStart+leftSize-1)
	root.Right = buildNode(inorder, rootIndex+1, inEnd, postorder, postStart+leftSize, postEnd-1)

	return root
}

func sliceToMap(inorder []int) map[int]int {
	for i, v := range inorder {
		indexMap[v] = i
	}
	return indexMap
}

// @lc code=end

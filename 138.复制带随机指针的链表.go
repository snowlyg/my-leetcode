package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func main() {
	root := &Node{
		Val: 7,
		Next: &Node{
			Val: 13,
			Next: &Node{
				Val: 11,
				Next: &Node{
					Val: 10,
					Next: &Node{
						Val: 1,
					},
				},
			},
		},
	}

	fmt.Println(copyRandomList(root))
}

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	originToClone := map[*Node]*Node{}
	var traverse func(head *Node) *Node
	traverse = func(head *Node) *Node {
		if head == nil {
			return nil
		}
		if n, ok := originToClone[head]; ok {
			return n
		}
		newNode := &Node{Val: head.Val}
		originToClone[head] = newNode
		newNode.Next = traverse(head.Next)
		newNode.Random = traverse(head.Random)
		return newNode
	}
	return traverse(head)
}

// @lc code=end

//go:build 24
// +build 24

package main

func main() {

	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	p := swapPairs(l1)
	for p != nil {
		println(p.Val)
		p = p.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

// @lc code=end

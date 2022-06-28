package main

func main() {
	addTwoNumbers(nil, nil)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var node *ListNode
	var next *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if node == nil {
			node = &ListNode{Val: sum}
			next = node
		} else {
			next.Next = &ListNode{Val: sum}
			next = next.Next
		}
	}
	if carry > 0 {
		next.Next = &ListNode{Val: carry}
	}
	return node
}

// @lc code=end

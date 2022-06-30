//go:build 21
// +build 21

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	p := mergeTwoLists(l1, l2)
	for p != nil {
		println(p.Val)
		p = p.Next
	}
}

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	sum := &ListNode{Val: -1}
	p := sum
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {

		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}

	if p.Next == p1 {
		p.Next = p2
	} else {
		p.Next = p1
	}

	return sum.Next

}

// @lc code=end

//go:build 25
// +build 25

package main

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	p := reverseKGroup(l1, 3)
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
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {

	// 空直接返回
	if head == nil {
		return head
	}

	// 分段处理，不够分段直接返回
	p1 := head
	p2 := p1
	for i := 0; i < k; i++ {
		if p2 == nil {
			return head
		}
		p2 = p2.Next
	}

	// 翻转指定长度
	p := reverse(p1, p2)

	// 递归
	p1.Next = reverseKGroup(p2, k)

	return p
}

// 翻转 l1-l2 之间的节点
func reverse(l1, l2 *ListNode) *ListNode {
	// 新初始节点
	ln := new(ListNode)
	// 双指针，当前位置和下一位
	cur := l1
	next := l1
	// 控制翻转区间
	for cur != l2 {
		// 前移下一位
		next = cur.Next

		cur.Next = ln
		ln = cur

		// 前移当前位置
		cur = next
	}
	return ln
}

// @lc code=end

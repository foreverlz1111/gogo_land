package main

import (
	"log"
)

//19. 删除链表的倒数第 N 个结点
//中等
//提示
//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//提示：
//
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100
//1 <= n <= sz

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	slow := head
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
func main() {
	head := &ListNode{}
	cur := head
	for i := 1; i < 5; i++ {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}

	cur = removeNthFromEnd(head, 4).Next

	for cur != nil {
		log.Println(cur)
		cur = cur.Next
	}
}

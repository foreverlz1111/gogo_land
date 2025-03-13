package main

import (
	"leet/util"
	"log"
)

//206. 反转链表
//简单
//相关标签
//相关企业
//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//示例 3：
//
//输入：head = []
//输出：[]
//
//
//提示：
//
//链表中节点的数目范围是 [0, 5000]
//-5000 <= Node.val <= 5000

func reverseList(head *util.ListNode) *util.ListNode {
	cur := head.Next
	var pre *util.ListNode

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
func main() {
	head := &util.ListNode{0, &util.ListNode{}}
	cur := head
	for i := 1; i < 9; i++ {
		cur.Next = &util.ListNode{i, nil}
		cur = cur.Next
	}

	cur = reverseList(head)
	for cur != nil {
		log.Println(cur)
		cur = cur.Next
	}
}

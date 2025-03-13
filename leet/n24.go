package main

import (
	"leet/util"
	"log"
)

// 24. 两两交换链表中的节点
// 中等
// 给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
// 示例 1：
//
// 输入：head = [1,2,3,4]
// 输出：[2,1,4,3]
// 示例 2：
//
// 输入：head = []
// 输出：[]
// 示例 3：
//
// 输入：head = [1]
// 输出：[1]
//
// 提示：
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100

func swapPairs(head *util.ListNode) *util.ListNode {
	cur := head // head is vhead
	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next
		tmp3 := cur.Next.Next.Next
		cur.Next = cur.Next.Next
		cur.Next.Next = tmp1
		cur.Next.Next.Next = tmp3
		cur = cur.Next.Next
	}
	return head
}

func main() {
	head := &util.ListNode{}
	cur := head
	for i := 1; i < 9; i++ {
		cur.Next = &util.ListNode{Val: i}
		cur = cur.Next
	}

	cur = swapPairs(head)
	for cur.Next != nil {
		log.Println(cur.Next)
		cur = cur.Next
	}
}

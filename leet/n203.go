package main

import (
	"fmt"
	"leet/util"
)

// 203. 移除链表元素
// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
//
// 示例 1：
//
// 输入：head = [1,2,6,3,4,5,6], val = 6
// 输出：[1,2,3,4,5]
// 示例 2：
//
// 输入：head = [], val = 1
// 输出：[]
// 示例 3：
//
// 输入：head = [7,7,7,7], val = 7
// 输出：[]
//
// 提示：
//
// 列表中的节点数目在范围 [0, 104] 内
// 1 <= Node.val <= 50
// 0 <= val <= 50

func removeElements(head *util.ListNode, val int) *util.ListNode {
	end := false
	// 用一个虚头
	h := &util.ListNode{}
	h.Next = head
	cur := h
	for !end {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}
		if cur.Next == nil {
			// 当前已经是尾
			end = true
		} else {
			cur = cur.Next
		}

	}

	return h.Next
}

func main() {

	ready_list := []int{6, 0, 1, 6, 2, 6, 3, 4, 5, 6, 7, 6}
	head := &util.ListNode{}
	cur := head
	for i := 0; i < len(ready_list); i++ {
		cur.Next = &util.ListNode{Val: ready_list[i]}
		cur = cur.Next
	}

	new_listnode := removeElements(head.Next, 6)

	end := false
	new_list := []int{}
	for !end {
		new_list = append(new_list, new_listnode.Val)
		if new_listnode.Next == nil {
			end = true
		} else {
			new_listnode = new_listnode.Next
		}
	}
	fmt.Println(new_list)
}

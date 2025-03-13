package main

import (
	"leet/util"
	"log"
)

// 232. 用栈实现队列
// 简单
// 请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：
//
// 实现 MyQueue 类：
//
// void push(int x) 将元素 x 推到队列的末尾
// int pop() 从队列的开头移除并返回元素
// int peek() 返回队列开头的元素
// boolean empty() 如果队列为空，返回 true ；否则，返回 false
// 说明：
//
// 你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
// 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
//
// 示例 1：
//
// 输入：
// ["MyQueue", "push", "push", "peek", "pop", "empty"]
// [[], [1], [2], [], [], []]
// 输出：
// [null, null, null, 1, 1, false]
//
// 解释：
// MyQueue myQueue = new MyQueue();
// myQueue.push(1); // queue is: [1]
// myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
// myQueue.peek(); // return 1
// myQueue.pop(); // return 1, queue is [2]
// myQueue.empty(); // return false
//
// 提示：
//
// 1 <= x <= 9
// 最多调用 100 次 push、pop、peek 和 empty
// 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）
//
// 进阶：
//
// 你能否实现每个操作均摊时间复杂度为 O(1) 的队列？换句话说，执行 n 个操作的总时间复杂度为 O(n) ，即使其中一个操作可能花费较长时间。

type MyQueue struct {
	// 队列的数据结构包含一个入栈和一个出栈
	stackIn  *util.MyIntStack
	stackOut *util.MyIntStack
}

func constructor() MyQueue {
	return MyQueue{
		stackIn:  &util.MyIntStack{},
		stackOut: &util.MyIntStack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	this.fillOutStack()
	return this.stackOut.Pop()
}

func (this *MyQueue) Peek() int {
	// 取出再弹入也可以实现看出栈的数字，减少代码冗余
	// ...
	// ...

	this.fillOutStack()
	return this.stackOut.Top()
}

func (this *MyQueue) fillOutStack() {
	// 才peek和pop之前，全部放入出栈
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			v := this.stackIn.Pop()
			this.stackOut.Push(v)
		}
	}
}

func (this *MyQueue) Empty() bool {
	return len(*this.stackIn) == 0 && len(*this.stackOut) == 0
}

func main() {
	obj := constructor()

	obj.Push(1)
	obj.Push(2)
	log.Println(obj.Peek()) // 需要为 1

	obj.Push(3)
	param_2 := obj.Pop()
	log.Println(param_2) // 需要为 1

	param_3 := obj.Peek()
	log.Println(param_3) // 需要为 2

	obj.Push(4)
	param_4 := obj.Empty()
	log.Println(param_4) // 不为空
}

package main

import "log"

// 225. 用队列实现栈
// 简单
// 请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。
//
// 实现 myStack 类：
//
// void push(int x) 将元素 x 压入栈顶。
// int pop() 移除并返回栈顶元素。
// int top() 返回栈顶元素。
// boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
//
// 注意：
//
// 你只能使用队列的标准操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
// 你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。
//
// 示例：
//
// 输入：
// ["myStack", "push", "push", "top", "pop", "empty"]
// [[], [1], [2], [], [], []]
// 输出：
// [null, null, null, 2, 2, false]
//
// 解释：
// myStack myStack = new myStack();
// myStack.push(1);
// myStack.push(2);
// myStack.top(); // 返回 2
// myStack.pop(); // 返回 2
// myStack.empty(); // 返回 False
//
// 提示：
//
// 1 <= x <= 9
// 最多调用100 次 push、pop、top 和 empty
// 每次调用 pop 和 top 都保证栈不为空

type myStack struct {
	val []int
}

// 使用1个队列实现栈：只需要分别弹出并插入前面n-1个
func constructor() myStack {
	return myStack{
		make([]int, 0),
	}
}

func (this *myStack) Push(x int) {
	//入栈(放入队列)
	this.val = append(this.val, x)
}

func (this *myStack) Pop() int {
	// 出栈
	l := len(this.val) - 1
	for l != 0 {
		v := this.val[0]
		this.val = this.val[1:]
		this.val = append(this.val, v)
		l--
	}
	v := this.val[0]
	this.val = this.val[1:]
	return v
}

func (this *myStack) Top() int {
	//查看第一个的出栈

	v := this.Pop()
	this.val = append(this.val, v)
	return v
}

func (this *myStack) Empty() bool {
	return len(this.val) == 0
}

func main() {
	obj := constructor()
	obj.Push(1)
	obj.Push(2)
	obj.Push(3)
	// ->3 -> 2 -> 1]

	param_2 := obj.Pop()
	log.Println(param_2) // 应该出 3

	param_3 := obj.Top()
	log.Println(param_3) // 应该出 2

	param_3 = obj.Top()
	log.Println(param_3) // 应该继续出 2

	param_4 := obj.Empty()
	log.Println(param_4) // 应该为false
}

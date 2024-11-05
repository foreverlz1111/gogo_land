package main

import "fmt"

// 检查逃逸：go build -gcflags="-m" escape_analysis.go
// Go 编译器怎么知道某个变量需要分配在栈上，还是堆上
// 栈内存的分配和释放非常高效。当一个函数被调用时，Go 会为该函数的局部变量分配栈空间，当函数返回时，栈空间会自动释放。放在栈上可以减轻GC压力，减轻内存开销
// 返回指针、Println、func literal迭代函数、空间不足、使用管道
// 堆是一种相对灵活的内存结构，适用于需要在函数返回后仍然存在的变量，或者大小动态的内存分配
func example1() *int {
	x := 10
	return &x // x 逃逸到堆上，因为它的地址被返回
}

func example1_1() int {
	x := 10
	return x // x 不会逃逸，分配在栈上
}
func exampleprint() {
	name := "string"
	fmt.Println(name)
}

func main() {
	s := make([]int, 1000, 1000)
	for index := range s {
		s[index] = index
	}
}

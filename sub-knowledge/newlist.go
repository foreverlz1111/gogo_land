package main

import "fmt"

// new 是 生成地址
// new 返回指针地址

func main() {
	list := new([]int)
	*list = append(*list, 1)
	// 错误用法：list = append(list, 1)
	fmt.Println(list)
	fmt.Println(*list)
}

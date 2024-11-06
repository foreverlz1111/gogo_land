package main

import "fmt"

// make 是初始化类型为 slice、map、chan 的数据
// make 返回第一个元素,可预设内存空间,避免未来的内存拷贝
// len 和 cap的区别：当前长度与当前容量，append可以进行扩容
// slice 切片是连续内存并且可以动态扩展
func main() {
	list := make([]int, 4)
	list = append(list, 1, 1)
	fmt.Println(list)

}

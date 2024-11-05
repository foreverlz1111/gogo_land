package main

import "fmt"

// make 是 生成
func main() {
	list := make([]int, 4)
	list = append(list, 1, 1)
	fmt.Println(list)
}

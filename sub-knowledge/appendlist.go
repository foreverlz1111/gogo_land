package main

import "fmt"

// 哈希表连接
func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := append(s1, s2...)
	fmt.Println(s3)
	s4 := append(s1, s3[:2]...)
	fmt.Println(s4)
}

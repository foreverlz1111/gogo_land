package main

import (
	"fmt"
)

// 242. 有效的字母异位词
// 简单
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的 字母异位词。
//
// 示例 1:
//
// 输入: s = "anagram", t = "nagaram"
// 输出: true
// 示例 2:
//
// 输入: s = "rat", t = "car"
// 输出: false
//
// 提示:
//
// 1 <= s.length, t.length <= 5 * 104
// s 和 t 仅包含小写字母

func isAnagram(s string, t string) bool {
	hashdit := [26]int{}
	for _, ch := range s {
		hashdit[ch-'a']++
	}
	for _, ch := range t {
		hashdit[ch-'a']--
	}
	if hashdit == [26]int{} {
		return true
	}
	return false
}
func main() {
	s := "aaeeb"
	t := "baaee"
	fmt.Println(isAnagram(s, t))
}

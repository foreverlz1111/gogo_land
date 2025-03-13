package main

import (
	"leet/util"
	"log"
)

// 20. 有效的括号
// 简单
// 提示
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
// 示例 1：
//
// 输入：s = "()"
//
// 输出：true
//
// 示例 2：
//
// 输入：s = "()[]{}"
//
// 输出：true
//
// 示例 3：
//
// 输入：s = "(]"
//
// 输出：false
//
// 示例 4：
//
// 输入：s = "([])"
//
// 输出：true
//
// 提示：
//
// 1 <= s.length <= 104
// s仅由括号 '()[]{}' 组成

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stacks := util.MyStrStack{}
	for _, ch := range s {
		if ch == '(' {
			stacks.Push(")")
		} else if ch == '{' {
			stacks.Push("{")
		} else if ch == '[' {
			stacks.Push("]")
		} else if stacks.Empty() && stacks.Pop() != string(ch) {
			return false
		}
	}

	return true
}
func main() {
	s := "([])"
	log.Println(isValid(s))
}

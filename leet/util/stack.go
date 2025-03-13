package util

// 栈的数据结构
type MyStrStack []string
type MyIntStack []int

func (s *MyStrStack) Push(x string) {
	// 入栈
	*s = append(*s, x)
}
func (s *MyStrStack) Pop() string {
	// 出栈
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}
func (s *MyStrStack) Top() string {
	// peek 栈首
	l := len(*s)
	v := (*s)[l-1]
	return v
}
func (s *MyStrStack) Empty() bool {
	return len(*s) == 0
}

// ---

func (s *MyIntStack) Push(x int) {
	// 入栈
	*s = append(*s, x)
}
func (s *MyIntStack) Pop() int {
	// 出栈
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}
func (s *MyIntStack) Top() int {
	// peek 栈首
	l := len(*s)
	v := (*s)[l-1]
	return v
}
func (s *MyIntStack) Empty() bool {
	return len(*s) == 0
}

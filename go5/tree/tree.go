package tree

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}
func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}
func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		fmt.Println("check left,current", t.Value)
		s += t.Left.String() + "\t-"
	}
	fmt.Println("add+ ", t.Value)
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		fmt.Println("check right,currnt ", t.Value)
		s += "-\t" + t.Right.String()
	}
	return "" + s + ""
}

package ch2_test

import (
	"Go_language/ch2"
	"testing"
)

func Test(t *testing.T) {
	t.Logf("res=%v", ch2.Reverse1([]int64{1, 2, 3, 4, 5, 6}))
	t.Logf("res=%v", ch2.Reverse2([]any{"1", "2", "3", "4", "5", "6"}))
	tree := new(ch2.Tree[int])
	root := new(ch2.Node[int])
	root.Data = 1
	root.Left = nil
	root.Right = nil
	tree.Root = root
	tree.Cmp = ch2.CmpInt
	t.Logf("res=%+v", tree.Find(1))
}

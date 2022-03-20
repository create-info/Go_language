package ch2

import "testing"

func TestDeferFun(t *testing.T) {
	DeferFun()
	DeferFun2(&Ch2Obj1{status: 0})
}

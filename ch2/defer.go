package ch2

import (
	"fmt"
	"strconv"
	"time"
)

type Ch2Obj1 struct {
	status uint32
}

func DeferFun() {
	start := time.Now()
	var code uint32
	defer func() {
		cost := time.Since(start).Milliseconds()
		fmt.Printf("%s, %s, %f , %v\n", "SendPN", strconv.Itoa(int(code)), float64(cost), code == 0)
	}()
	time.Sleep(1 * time.Second)
	code = 1
	return
}

func DeferFun2(obj *Ch2Obj1) {
	defer func() {
		fmt.Printf("%s, %d\n", "SendPN", obj.status)
	}()
	obj.status = 1
	return
}

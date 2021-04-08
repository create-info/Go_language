package go_timer

import (
	"errors"
	"fmt"
	"time"
)

/*
	使用go中的timer实现在当前时间开始的一段时间后再执行某个任务
 */
func GoTimer()  {
	start := time.Now()
	fmt.Printf("time.now:%v \n", start)
	err := task("推送消息")

	//NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	var t = time.NewTimer(5 * time.Second)
	//下一次执行的时间
	//d := []time.Duration{time.Minute*5, time.Minute*10}
	//var count uint32
	//task执行失败，下次再尝试
	if err != nil {
		for i:=1; i<3; i++ {
			<-t.C
			if task("推送消息") != nil {
				//设置当前时间10s后再次执行
				t.Reset(time.Second*10)
			} else {
				//成功
				return
			}
		}
	}
	fmt.Printf("time.now:%vs\n", time.Since(start).Seconds())
	fmt.Println("done")
}

func task(msg string) error {
	fmt.Printf("runing task %s, now:%v: \n", msg, time.Now())
	return errors.New(msg)
}
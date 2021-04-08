package go_timer

import "testing"

func TestGoTimer(t *testing.T) {
	GoTimer()
	//输出
	/*
	time.now:2021-04-09 00:05:44.03932 +0800 CST m=+0.000606864
	runing task 推送消息, now:2021-04-09 00:05:44.039567 +0800 CST m=+0.000853536:
	runing task 推送消息, now:2021-04-09 00:05:49.040743 +0800 CST m=+5.002027844:
	runing task 推送消息, now:2021-04-09 00:05:59.043021 +0800 CST m=+15.004301598:
	time.now:15.003895743s
	 */
}

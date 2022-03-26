package main

//1.必须是main包：package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Hello Word!")

	time.Sleep(1 * time.Second)
	//int64 to float64
	cost := time.Since(start).Milliseconds()
	fmt.Println(float64(cost))
}

//运行结果：
//E:\goWorkPath
//λ go run test01.go
//go run: cannot run non-main package

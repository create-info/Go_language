package main

import(
	"fmt"
	"os"
)

//main函数不支持传入参数
//func main(arg [] string)是不正确的
//在程序中直接通过os.Args获取命令行参数

func main(){
	//打印出命令行参数
	fmt.Println(os.Args)
	
	if len(os.Args) > 1 {
		fmt.Println("参数:", os.Args[1])
	}	
	
	fmt.Println("你好！")
	
	os.Exit(-1)
}

//输出结果：
/*
E:\goWorkPath
λ go run test05.go A,B
[C:\Users\xxm\AppData\Local\Temp\go-build984562086\b001\exe\1.exe A,B]
参数: A,B
你好！
exit status 4294967295
*/


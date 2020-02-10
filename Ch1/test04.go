package main

import "fmt"
import "os"

//退出返回值
//1、Go中main函数不支持任何返回值
//2、通过os.Exit来返回状态

func main(){
  fmt.Println("你好！")
  //return 1
  os.Exit(0)
}


//运行结果
/*
λ go run test04.go
# command-line-arguments
.\test04.go:11:3: too many arguments to return
        have (number)
        want ()
*/        


/*
λ go run test04.go
你好！
*/

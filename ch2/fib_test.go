package first_test

import(
	"fmt"
	"testing"
)

//测试文件必须以_test.go结尾命名，测试方法必须以Test开头
//func TestFibList(t *testing.T)
//单元测试文件 (*_test.go) 里的测试入口必须以 Test 开始，参数为 *testing.T 的函数
func TestFibList(t *testing.T){
	//var a int = 1
	//var b int = 1
	//变量a和b的声明可以简写为
	var (
		a int = 1 //或者a = 1
		b = 1 //或者b = 1
	)
  
  //或者a:=1; b:=1
	
	fmt.Print(a)
	
	/*var tmp int
	var i int
	
	//go语言中只有for，没有while
	for i = 0; i < 5; i++{
		fmt.Print(b)
		b = a
		tmp = b
		b = b + tmp 
	} 
	*/
	
	
	//自动类型推断
	for i := 0; i < 5; i++{
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a 
	} 
	fmt.Println()
}

//输出结果：
/*
E:\goWorkPath
λ go test fib_test.go
ok      command-line-arguments  2.795s
*/






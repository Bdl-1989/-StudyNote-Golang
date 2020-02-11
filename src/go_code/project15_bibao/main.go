package main

import "fmt"

//累加器
func AddUpper() func(int) (int, string) {
	var n int = 10
	var str string = "hello"
	return func(x int) (int, string) {
		n = n + x
		str += "string(x)"
		//fmt.Println(str)
		return n, str
	}
} //闭包
//闭包是类，函数是操作，n是字段。函数和它使用n构成闭包。

func main() {
	//累加器
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
}

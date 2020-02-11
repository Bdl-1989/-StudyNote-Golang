package main

import "fmt"

//指针
func main() {
	var i int = 10
	//address,  we use "&"
	fmt.Println("i's address is ", &i)
	//ptr 指针变量
	//ptr 类型*int
	//ptr 值&i
	var ptr *int = &i
	fmt.Println("ptr is ", ptr)
	fmt.Println("ptr's address is ", &ptr)
}
